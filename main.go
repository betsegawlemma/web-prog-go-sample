package main

import (
	"database/sql"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu/repository"
	"github.com/betsegawlemma/restaurant/menu/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var categoryServ *service.CategoryServiceImpl

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	categories, err := categoryServ.Categories()
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "index.layout", categories)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.layout", nil)
}

func menu(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "menu.layout", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.layout", nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "admin.index.layout", nil)
}

func adminCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryServ.Categories()
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "admin.categ.layout", categories)
}

func adminCategoriesNew(w http.ResponseWriter, r *http.Request) {

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

		err = categoryServ.StoreCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {

		tmpl.ExecuteTemplate(w, "admin.categ.new.layout", nil)

	}
}

func adminCategoriesUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := categoryServ.Category(id)

		if err != nil {
			panic(err)
		}

		tmpl.ExecuteTemplate(w, "admin.categ.update.layout", cat)

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

		err = categoryServ.UpdateCategory(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}

}

func adminCategoriesDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = categoryServ.DeleteCategory(id)

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

	path := filepath.Join(wd, "delivery", "web", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

func main() {

	dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
	categoryServ = service.NewCategoryServiceImpl(categoryRepo)

	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/menu", menu)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/categories", adminCategories)
	http.HandleFunc("/admin/categories/new", adminCategoriesNew)
	http.HandleFunc("/admin/categories/update", adminCategoriesUpdate)
	http.HandleFunc("/admin/categories/delete", adminCategoriesDelete)
	http.ListenAndServe(":8181", nil)
}
