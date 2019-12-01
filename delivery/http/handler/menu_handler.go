package handler

import (
	"html/template"
	"net/http"

	"github.com/betsegawlemma/restaurant/menu"
)

// MenuHandler handles menu related requests
type MenuHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
}

// NewMenuHandler initializes and returns new MenuHandler
func NewMenuHandler(T *template.Template, CS menu.CategoryService) *MenuHandler {
	return &MenuHandler{tmpl: T, categorySrv: CS}
}

// Index handles request on route /
func (mh *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	categories, err := mh.categorySrv.Categories()
	if err != nil {
		panic(err)
	}

	mh.tmpl.ExecuteTemplate(w, "index.layout", categories)
}

// About handles requests on route /about
func (mh *MenuHandler) About(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "about.layout", nil)
}

// Menu handle request on route /menu
func (mh *MenuHandler) Menu(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "menu.layout", nil)
}

// Contact handle request on route /Contact
func (mh *MenuHandler) Contact(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "contact.layout", nil)
}

// Admin handle request on route /admin
func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}
