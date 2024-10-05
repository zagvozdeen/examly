package main

import (
	"errors"
	"fmt"
	"github.com/den4ik117/examly/internal/enum"
	"github.com/den4ik117/examly/internal/store"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const FileDir = "files"

func (app *application) uploadFile(w http.ResponseWriter, r *http.Request) {
	if ok := app.checkRole(w, r, enum.MemberRole); !ok {
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			app.log.Err(err).Msg("failed to close file")
		}
	}()

	user := getUserFromRequest(r)

	err = os.MkdirAll(FileDir, os.ModePerm)
	if err != nil {
		app.log.Err(err).Msg("failed to create directory")
		err = errors.Join(err, errors.New("failed to create directory"))
		app.internalServerError(w, r, err)
		return
	}

	uid, err := uuid.NewV7()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	filePath := fmt.Sprintf(
		"/%s/%s%s",
		FileDir,
		uid.String(),
		strings.ToLower(filepath.Ext(handler.Filename)),
	)
	dst, err := os.Create("./public" + filePath)
	defer func() {
		if err := dst.Close(); err != nil {
			app.log.Err(err).Msg("failed to close dst")
		}
	}()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	model := &store.File{
		UUID:       uid.String(),
		Content:    filePath,
		Size:       int(handler.Size),
		MimeType:   handler.Header.Get("Content-Type"),
		OriginName: handler.Filename,
		CreatedBy:  user.ID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = app.store.FilesStore.Create(r.Context(), model)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusCreated, map[string]any{
		"data": model,
	})
}
