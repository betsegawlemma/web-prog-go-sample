package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu"
)

// AdminCategoryHandler handles category handler admin requests
type AdminCategoryHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminCategoryHandler(T *template.Template, CS menu.CategoryService) *AdminCategoryHandler {
	return &AdminCategoryHandler{tmpl: T, categorySrv: CS}
}

// AdminCategories handle requests on route /admin/categories
func (ach *AdminCategoryHandler) AdminCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := ach.categorySrv.Categories()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.categ.layout", categories)
}

// AdminCategoriesNew hanlde requests on route /admin/categories/new
func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.Category{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.categorySrv.StoreCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", nil)

	}
}

// AdminCategoriesUpdate handle requests on /admin/categories/update
func (ach *AdminCategoryHandler) AdminCategoriesUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.categorySrv.Category(id)

		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.categ.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Category{}
		ctg.ID, _ = strconv.Atoi(r.FormValue("id"))
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		ctg.Image = r.FormValue("image")

		mf, _, err := r.FormFile("catimg")

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		writeFile(&mf, ctg.Image)

		err = ach.categorySrv.UpdateCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}

}

// AdminCategoriesDelete handle requests on route /admin/categories/delete
func (ach *AdminCategoryHandler) AdminCategoriesDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.categorySrv.DeleteCategory(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
