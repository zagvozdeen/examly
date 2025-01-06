package api

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/guregu/null/v5"
	"github.com/zagvozdeen/examly/internal/enum"
	"github.com/zagvozdeen/examly/internal/store"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) getQuestions(w http.ResponseWriter, r *http.Request) {
	var filter store.GetQuestionsFilter
	query := r.URL.Query()
	user := getUserFromRequest(r)
	ctx := r.Context()

	if query.Has("created_by") {
		id, err := strconv.Atoi(query.Get("created_by"))
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
		if id != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
			return
		}
		filter.CreatedBy = id
	}
	if query.Get("all") == "true" {
		if user.Role.Level() < enum.ModeratorRole.Level() {
			app.forbiddenErrorResponse(w, r, errors.New("forbidden"))
			return
		}
		filter.All = true
	}

	questions, err := app.store.QuestionsStore.Get(ctx, filter)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": questions,
	})
}

type CreateQuestionPayload struct {
	Title       string          `json:"title" validate:"required"`
	Content     string          `json:"content" validate:""`
	Explanation string          `json:"explanation" validate:""`
	Type        string          `json:"type" validate:"required"`
	CourseID    int             `json:"course_id" validate:"required"`
	ModuleID    int             `json:"module_id" validate:""`
	Options     []OptionPayload `json:"options" validate:"required,dive,required"`
	TagsIDs     []int           `json:"tags_ids" validate:""`
}

type OptionPayload struct {
	ID        int    `json:"id" validate:"required"`
	Content   string `json:"content" validate:"required"`
	IsCorrect bool   `json:"is_correct" validate:""`
}

func (app *Application) createQuestion(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	var payload CreateQuestionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	t, err := enum.NewQuestionType(payload.Type)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromRequest(r)

	correct := false
	var options store.Options
	for i, answer := range payload.Options {
		correct = correct || answer.IsCorrect || t == enum.PlaintextQuestionType
		options = append(options, store.Option{
			ID:        i + 1,
			Content:   answer.Content,
			IsCorrect: answer.IsCorrect,
		})
	}

	if !correct {
		app.badRequestResponse(w, r, errors.New("at least one answer should be correct"))
		return
	}

	question := &store.Question{
		UUID:        uid.String(),
		Title:       payload.Title,
		Content:     null.NewString(payload.Content, payload.Content != ""),
		Explanation: null.NewString(payload.Explanation, payload.Explanation != ""),
		Type:        t,
		Status:      enum.CreatedStatus,
		CourseID:    payload.CourseID,
		ModuleID:    null.NewInt(int64(payload.ModuleID), payload.ModuleID != 0),
		CreatedBy:   user.ID,
		Options:     options,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := app.store.QuestionsStore.Create(r.Context(), question); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.store.TagsStore.CreateQuestionTags(r.Context(), question.ID, payload.TagsIDs)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": question,
	})
}

func (app *Application) getQuestion(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("missing uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	question, err := app.store.QuestionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if question.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to view this question"))
		return
	}

	ids, err := app.store.TagsStore.GetQuestionTags(ctx, question.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	question.TagsIDs = ids

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": question,
	})
}

type UpdateQuestionPayload struct {
	Title       string          `json:"title" validate:"required"`
	Content     string          `json:"content" validate:""`
	Explanation string          `json:"explanation" validate:""`
	Type        string          `json:"type" validate:"required"`
	CourseID    int             `json:"course_id" validate:"required"`
	FileID      int             `json:"file_id" validate:""`
	ModuleID    int             `json:"module_id" validate:""`
	Answers     []OptionPayload `json:"options" validate:"required,dive,required"`
	TagsIDs     []int           `json:"tags_ids" validate:""`
}

func (app *Application) updateQuestion(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	var payload UpdateQuestionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	question, err := app.store.QuestionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if question.CreatedBy != user.ID {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to update this question"))
		return
	}

	newUUID, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	t, err := enum.NewQuestionType(payload.Type)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	correct := false
	var options store.Options
	for i, answer := range payload.Answers {
		correct = correct || answer.IsCorrect || t == enum.PlaintextQuestionType
		options = append(options, store.Option{
			ID:        i + 1,
			Content:   answer.Content,
			IsCorrect: answer.IsCorrect,
		})
	}

	if !correct {
		app.badRequestResponse(w, r, errors.New("at least one answer should be correct"))
		return
	}

	nextQuestion := &store.Question{
		UUID:           newUUID.String(),
		Title:          payload.Title,
		Content:        null.NewString(payload.Content, payload.Content != ""),
		Explanation:    null.NewString(payload.Explanation, payload.Explanation != ""),
		Type:           t,
		Status:         enum.CreatedStatus,
		CourseID:       payload.CourseID,
		ModuleID:       null.NewInt(int64(payload.ModuleID), payload.ModuleID != 0),
		PrevQuestionID: null.IntFrom(int64(question.ID)),
		CreatedBy:      user.ID,
		Options:        options,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = app.store.QuestionsStore.Create(ctx, nextQuestion)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.store.TagsStore.CreateQuestionTags(r.Context(), nextQuestion.ID, payload.TagsIDs)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	question.NextQuestionID = null.IntFrom(int64(nextQuestion.ID))
	question.UpdatedAt = time.Now()

	err = app.store.QuestionsStore.Update(ctx, &question)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": question,
	})
}

func (app *Application) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	question, err := app.store.QuestionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if question.CreatedBy != user.ID && user.Role.Level() < enum.ModeratorRole.Level() {
		app.forbiddenErrorResponse(w, r, errors.New("you are not allowed to delete this question"))
		return
	}

	question.DeletedAt = null.TimeFrom(time.Now())
	question.UpdatedAt = time.Now()

	err = app.store.QuestionsStore.Delete(ctx, &question)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type ModerateQuestionPayload struct {
	ModerationReason string `json:"moderation_reason" validate:"max=1024"`
	Status           string `json:"status" validate:"required"`
}

func (app *Application) moderateQuestion(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.ModeratorRole); !ok {
		return
	}

	var payload ModerateQuestionPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	uid, ok := mux.Vars(r)["uuid"]
	if !ok {
		app.badRequestResponse(w, r, errors.New("empty uuid"))
		return
	}

	ctx := r.Context()
	user := getUserFromRequest(r)

	question, err := app.store.QuestionsStore.GetByUUID(ctx, uid)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	s, err := enum.NewStatus(payload.Status)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	question.ModerationReason = null.NewString(payload.ModerationReason, payload.ModerationReason != "")
	question.Status = s
	question.UpdatedAt = time.Now()
	question.ModeratedBy = null.IntFrom(int64(user.ID))

	err = app.store.QuestionsStore.UpdateStatus(ctx, &question)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": question,
	})
}
