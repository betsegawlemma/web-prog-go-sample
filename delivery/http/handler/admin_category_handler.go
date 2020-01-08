package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/form"
	"github.com/betsegawlemma/restaurant/menu"
	"github.com/betsegawlemma/restaurant/rtoken"
)

// AdminCategoryHandler handles category handler admin requests
type AdminCategoryHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
	csrfSignKey []byte
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminCategoryHandler(t *template.Template, cs menu.CategoryService, csKey []byte) *AdminCategoryHandler {
	return &AdminCategoryHandler{tmpl: t, categorySrv: cs, csrfSignKey: csKey}
}

// AdminCategories handle requests on route /admin/categories
func (ach *AdminCategoryHandler) AdminCategories(w http.ResponseWriter, r *http.Request) {
	categories, errs := ach.categorySrv.Categories()
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
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
	ach.tmpl.ExecuteTemplate(w, "admin.categ.layout", tmplData)
}

// AdminCategoriesNew hanlde requests on route /admin/categories/new
func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		newCatForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		newCatForm.Required("catname", "catdesc")
		newCatForm.MinLength("catdesc", 10)
		newCatForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newCatForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
			return
		}
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			newCatForm.VErrors.Add("catimg", "File error")
			ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
			return
		}
		defer mf.Close()
		ctg := &entity.Category{
			Name:        r.FormValue("catname"),
			Description: r.FormValue("catdesc"),
			Image:       fh.Filename,
		}
		writeFile(&mf, fh.Filename)
		_, errs := ach.categorySrv.StoreCategory(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}
}

// AdminCategoriesUpdate handle requests on /admin/categories/update
func (ach *AdminCategoryHandler) AdminCategoriesUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		cat, errs := ach.categorySrv.Category(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		values := url.Values{}
		values.Add("catid", idRaw)
		values.Add("catname", cat.Name)
		values.Add("catdesc", cat.Description)
		values.Add("catimg", cat.Image)
		upCatForm := struct {
			Values   url.Values
			VErrors  form.ValidationErrors
			Category *entity.Category
			CSRF     string
		}{
			Values:   values,
			VErrors:  form.ValidationErrors{},
			Category: cat,
			CSRF:     token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categ.update.layout", upCatForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		updateCatForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		updateCatForm.Required("catname", "catdesc")
		updateCatForm.MinLength("catdesc", 10)
		updateCatForm.CSRF = token

		catID, err := strconv.Atoi(r.FormValue("catid"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		ctg := &entity.Category{
			ID:          uint(catID),
			Name:        r.FormValue("catname"),
			Description: r.FormValue("catdesc"),
			Image:       r.FormValue("imgname"),
		}
		mf, fh, err := r.FormFile("catimg")
		if err == nil {
			ctg.Image = fh.Filename
			err = writeFile(&mf, ctg.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		_, errs := ach.categorySrv.UpdateCategory(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
		return
	}
}

// AdminCategoriesDelete handle requests on route /admin/categories/delete
func (ach *AdminCategoryHandler) AdminCategoriesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := ach.categorySrv.DeleteCategory(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
