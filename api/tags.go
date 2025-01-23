package api

import (
	"net/http"
)

func (app *Application) getTags(w http.ResponseWriter, r *http.Request) {
	tags, err := app.store.TagsStore.Get(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.jsonResponse(w, r, http.StatusOK, map[string]any{
		"data": tags,
	})
}
