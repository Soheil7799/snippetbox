package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Soheil7799/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/pages/home.gohtml",
		"ui/html/base.gohtml",
		"ui/html/partials/nav.gohtml",
	}
	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	snippets, err := app.DB.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}
	// for _, snippet := range snippets {
	// 	fmt.Fprint(w, "%+v\n", snippet)

	// }
	templateData := templateData{
		Snippets: snippets,
	}
	err = templateSet.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/partials/nav.gohtml",
		"./ui/html/pages/view.gohtml",
	}
	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
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
	templateData := templateData{
		Snippet: snippet,
	}

	err = templateSet.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is snippet create page"))
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
