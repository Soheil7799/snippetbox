package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Soheil7799/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.DB.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.gohtml", data)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.DB.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.gohtml", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "create.gohtml", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	id, err := app.DB.Insert(title, content, expires)
	if err != nil {
		app.logger.Error("could not create the snippet")
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
