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

	app.render(w, r, http.StatusOK, "home.gohtml", templateData{
		Snippets: snippets,
	})

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

	app.render(w, r, http.StatusOK, "view.gohtml", templateData{
		Snippet: snippet,
	})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("this is snippet create page"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.DB.Insert(title, content, expires)
	if err != nil {
		app.logger.Error("could not create the snippet")
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
