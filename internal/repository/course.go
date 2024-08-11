package repository

import (
	"github.com/Den4ik117/examly/internal/model"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetCourses() ([]model.Course, error) {
	courses := make([]model.Course, 0)

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE status = $1 AND deleted_at IS NULL",
		model.ActiveCourseStatus,
	)

	return courses, err
}

func (r *CourseRepository) GetCoursesByUserID(id int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE user_id = $1 AND deleted_at IS NULL",
		id,
	)

	return courses, err
}

func (r *CourseRepository) GetAllCourses(id int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	err := r.db.Select(
		&courses,
		"SELECT * FROM courses WHERE (user_id = $1 OR status = $2) AND deleted_at IS NULL",
		id,
		model.ActiveCourseStatus,
	)

	return courses, err
}

func (r *CourseRepository) GetCoursesByIDs(ids []int) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query, args, err := sqlx.In("SELECT * FROM courses WHERE id IN (?) AND deleted_at IS NULL", ids)
	query = r.db.Rebind(query)
	err = r.db.Select(&courses, query, args...)

	return courses, err
}

func (r *CourseRepository) CreateCourse(course *model.Course) (id int, err error) {
	err = r.db.QueryRow(
		"INSERT INTO courses (uuid, name, description, color, icon, status, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		course.UUID,
		course.Name,
		course.Description,
		course.Color,
		course.Icon,
		course.Status,
		course.UserID,
		course.CreatedAt,
		course.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *CourseRepository) GetCourseByUUID(uuid string) (course model.Course, err error) {
	err = r.db.Get(
		&course,
		"SELECT * FROM courses WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1",
		uuid,
	)

	return course, err
}

func (r *CourseRepository) GetModulesByCourseID(id int) ([]model.Module, error) {
	modules := make([]model.Module, 0)

	err := r.db.Select(
		&modules,
		"SELECT * FROM modules WHERE course_id = $1 AND status = $2 AND deleted_at IS NULL",
		id,
		model.ActiveCourseStatus,
	)

	return modules, err
}

func (r *CourseRepository) GetQuestionsByCourseID(id int) ([]model.Question, error) {
	questions := make([]model.Question, 0)

	err := r.db.Select(
		&questions,
		"SELECT * FROM questions WHERE course_id = $1 AND status = $2 AND deleted_at IS NULL",
		id,
		model.ActiveCourseStatus,
	)

	return questions, err
}

func (r *CourseRepository) GetAnswersByIDs(ids []int) ([]model.Answer, error) {
	answers := make([]model.Answer, 0)

	query, args, err := sqlx.In("SELECT * FROM answers WHERE question_id IN (?) AND deleted_at IS NULL", ids)
	query = r.db.Rebind(query)
	err = r.db.Select(&answers, query, args...)

	return answers, err
}

func (r *CourseRepository) CreateUserQuestions(questions []model.UserQuestion) (map[int]int, error) {
	rows, err := r.db.NamedQuery(
		"INSERT INTO user_questions (uuid, content, explanation, type, is_true, sort, course_id, question_id, module_id, file_id, created_at, updated_at) VALUES (:uuid, :content, :explanation, :type, :is_true, :sort, :course_id, :question_id, :module_id, :file_id, :created_at, :updated_at) RETURNING id, question_id",
		questions,
	)
	if err != nil {
		return nil, err
	}

	ids := map[int]int{}

	for rows.Next() {
		var id, qid int
		err = rows.Scan(&id, &qid)
		if err != nil {
			return nil, err
		}
		ids[qid] = id
	}

	return ids, nil
}

func (r *CourseRepository) CreateUserCourse(course *model.UserCourse) error {
	return r.db.QueryRow(
		"INSERT INTO user_courses (uuid, name, type, user_id, course_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		course.UUID,
		course.Name,
		course.Type,
		course.UserID,
		course.CourseID,
		course.CreatedAt,
		course.UpdatedAt,
	).Scan(&course.ID)
}

func (r *CourseRepository) CreateUserModules(modules []model.UserModule) error {
	_, err := r.db.NamedQuery(
		"INSERT INTO user_modules (name, course_id, created_at, updated_at) VALUES (:name, :course_id, :created_at, :updated_at)",
		modules,
	)

	return err
}

func (r *CourseRepository) CreateUserAnswers(answers []model.UserAnswer) error {
	_, err := r.db.NamedQuery(
		"INSERT INTO user_answers (content, question_id, is_true, is_chosen, sort, created_at, updated_at) VALUES (:content, :question_id, :is_true, :is_chosen, :sort, :created_at, :updated_at)",
		answers,
	)

	return err
}

type CourseStats struct {
	Name      string `json:"name"`
	Type      string `json:"-" db:"type"`
	Total     int    `json:"total" db:"total"`
	Completed int    `json:"completed" db:"completed"`
	Sort      int    `json:"-"`
}

func (r *CourseRepository) GetUserStatsByCourse(id int) (d []CourseStats, err error) {
	err = r.db.Select(
		&d,
		`
SELECT c.type as type, COUNT(q.*) as total, COUNT(q.is_true) as completed
FROM (SELECT id,
             type,
             RANK() OVER (PARTITION BY type ORDER BY created_at DESC) AS rank
      FROM user_courses WHERE user_id = $1 AND deleted_at IS NULL) c
         JOIN user_questions q on q.course_id = c.id
WHERE c.rank = 1
GROUP BY c.type
`,
		id,
	)

	return d, err
}

func (r *CourseRepository) GetCourseStatsByUUID(params *model.CourseStatsParams) (s []model.FullCourseStats, err error) {
	if params.CourseID == 0 {
		err = r.db.Select(
			&s,
			`
			SELECT c.id                                           AS id,
				   c.uuid                                         AS uuid,
				   c.type                                         AS type,
				   c.created_at                                   AS created_at,
				   COUNT(q.*) FILTER ( WHERE q.is_true is true )  AS correct,
				   COUNT(q.*) FILTER ( WHERE q.is_true is false ) AS incorrect,
				   COUNT(q.*)                                     AS total
			FROM user_courses c
					 JOIN public.user_questions q on c.id = q.course_id
			WHERE c.user_id = $1 AND c.deleted_at IS NULL
			GROUP BY c.id, c.uuid, c.type, c.created_at
			ORDER BY c.id DESC;
			`,
			params.UserID,
		)

		return s, err
	}

	err = r.db.Select(
		&s,
		`
		SELECT c.id                                           AS id,
			   c.uuid                                         AS uuid,
			   c.type                                         AS type,
			   c.created_at                                   AS created_at,
			   COUNT(q.*) FILTER ( WHERE q.is_true is true )  AS correct,
			   COUNT(q.*) FILTER ( WHERE q.is_true is false ) AS incorrect,
			   COUNT(q.*)                                     AS total
		FROM user_courses c
				 JOIN public.user_questions q on c.id = q.course_id
		WHERE c.course_id = $1 AND c.user_id = $2 AND c.deleted_at IS NULL
		GROUP BY c.id, c.uuid, c.type, c.created_at
		ORDER BY c.id DESC;
		`,
		params.CourseID,
		params.UserID,
	)

	return s, err
}

func (r *CourseRepository) GetAllEntities() (
	courses []model.Course,
	modules []model.Module,
	questions []model.Question,
	answers []model.Answer,
	err error,
) {
	if err = r.db.Select(&courses, "SELECT * FROM courses WHERE deleted_at IS NULL"); err != nil {
		return nil, nil, nil, nil, err
	}

	if err = r.db.Select(&modules, "SELECT * FROM modules WHERE deleted_at IS NULL"); err != nil {
		return nil, nil, nil, nil, err
	}

	if err = r.db.Select(&questions, "SELECT * FROM questions WHERE deleted_at IS NULL"); err != nil {
		return nil, nil, nil, nil, err
	}

	if err = r.db.Select(&answers, "SELECT * FROM answers WHERE deleted_at IS NULL"); err != nil {
		return nil, nil, nil, nil, err
	}

	return courses, modules, questions, answers, err
}
