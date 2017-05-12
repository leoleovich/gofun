package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"fmt"
)

type Page struct {
	Title string
	Body  template.HTML
}

var templates = template.Must(template.ParseFiles("frontend/view.html", "frontend/edit.html", "frontend/delete.html", "frontend/root.html"))

func loadPage(title string) (*Page, error) {
	filename := "frontend/data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: template.HTML(body)}, nil
}

func (p *Page) save() error {
	filename := "frontend/data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func deletePage(title string) (error) {
	filename := "frontend/data/" + title + ".txt"
	return os.Remove(filename)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	fmt.Sprint(p)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{title, template.HTML(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := deletePage(title)
	if err != nil {
		renderTemplate(w, "delete", &Page{Title:title})
		return
	}
	http.Redirect(w, r, "/delete/"+title , http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request, title string) {
	files, err := ioutil.ReadDir("frontend/data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lof := ""
	for _, f := range files {
		filename := strings.Replace(f.Name(), ".txt", "", -1)
		lof += "<a href=\"/view/" + filename + "\">" + filename + "</a><br>"
	}

	renderTemplate(w, "root", &Page{Body: template.HTML(lof)})
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		split := strings.Split(r.URL.Path, "/")
		if len(split) <= 2 || split[2] == "" {
			if r.URL.Path == "/" {
				fn(w, r, "/")
			} else {
				http.NotFound(w, r)
			}
			return
		}
		fn(w, r, split[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.HandleFunc("/", makeHandler(rootHandler))

	http.ListenAndServe(":8080", nil)
}