package handler

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/form"
	"github.com/betsegawlemma/restaurant/menu"
	"github.com/betsegawlemma/restaurant/rtoken"
)

// MenuHandler handles menu related requests
type MenuHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
	csrfSignKey []byte
}

// NewMenuHandler initializes and returns new MenuHandler
func NewMenuHandler(T *template.Template, CS menu.CategoryService, csKey []byte) *MenuHandler {
	return &MenuHandler{tmpl: T, categorySrv: CS, csrfSignKey: csKey}
}

// Index handles request on route /
func (mh *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	categories, errs := mh.categorySrv.Categories()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Categories []entity.Category
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Categories: categories,
		CSRF:       token,
	}

	mh.tmpl.ExecuteTemplate(w, "index.layout", tmplData)
}

// About handles requests on route /about
func (mh *MenuHandler) About(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "about.layout", tmplData)
}

// Menu handle request on route /menu
func (mh *MenuHandler) Menu(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "menu.layout", tmplData)
}

// Contact handle request on route /Contact
func (mh *MenuHandler) Contact(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "contact.layout", tmplData)
}

// Admin handle request on route /admin
func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", tmplData)
}
