package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var method = r.Method
	var uri = r.RequestURI
	var trace = string(debug.Stack())

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	templateSet, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}
	buffer := new(bytes.Buffer)
	w.WriteHeader(status)
	err := templateSet.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
	w.WriteHeader(status)

	_, _ = buffer.WriteTo(w)
}
