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

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	//files := []string{
	//	"./ui/html/pages/home.tmpl",
	//	"./ui/html/base.tmpl",
	//	"./ui/html/partials/nav.tmpl",
	//}
	//template_set, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	http.Error(w, "Internal server error", http.StatusInternalServerError)
	//	return
	//}
	//err = template_set.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.serverError(w, err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		w.Write([]byte("No id was found\nDefault snippet view"))
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// fmt.Fprintf(w, "%+v", snippet)
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("before method check")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	//fmt.Println("handler called with post (after method check)")
	title := "0 snail"
	content := "0 snail\nClimb mount fuji,\nBut slowly slowly!\n\n- Kobayashi Issa"
	expires := 7
	//fmt.Println("before calling insert")
	id, err := app.snippets.Insert(title, content, expires)
	//fmt.Println("after calling insert")
	if err != nil {
		app.serverError(w, err)
		return
	}
	//fmt.Sprintf("after error checking insert with id: %d", id)
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
