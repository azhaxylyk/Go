package function

import (
	"net/http"
	"text/template"
)

func errorResponse(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := template.ParseFiles("../function/templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error (500)", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if err := tmpl.Execute(w, message); err != nil {
		http.Error(w, "Internal Server Error (500)", http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorResponse(w, http.StatusNotFound, "Not Found (404)")
		return
	}
	renderTemplate(w, nil)
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed (405)")
		return
	}
	text := r.FormValue("Stroka")
	banner := r.FormValue("banner")

	answer, err := generateAsciiArt(text, banner)
	if err != nil {
		switch err.Error() {
		case "Not Found":
			errorResponse(w, http.StatusNotFound, "Not Found (404)")
		case "Bad Request":
			errorResponse(w, http.StatusBadRequest, "Bad Request (400)")
		default:
			errorResponse(w, http.StatusInternalServerError, "Internal Server Error (500)")
		}
		return
	}

	data := struct {
		PassingString string
	}{
		PassingString: answer,
	}

	renderTemplate(w, data)
}

func renderTemplate(w http.ResponseWriter, data interface{}) {
	tmpl, err := template.ParseFiles("../function/templates/index.html")
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "Template Parsing Error: "+err.Error())
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		errorResponse(w, http.StatusInternalServerError, "Template Execution Error: "+err.Error())
		return
	}
}


