package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/betsegawlemma/restaurant/delivery/http/handler"
	"github.com/betsegawlemma/restaurant/menu/repository"
	"github.com/betsegawlemma/restaurant/menu/service"
)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
	categoryServ := service.NewCategoryServiceImpl(categoryRepo)

	adminCatgHandler := handler.NewAdminCategoryHandler(tmpl, categoryServ)
	menuHandler := handler.NewMenuHandler(tmpl, categoryServ)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", menuHandler.Index)
	http.HandleFunc("/about", menuHandler.About)
	http.HandleFunc("/contact", menuHandler.Contact)
	http.HandleFunc("/menu", menuHandler.Menu)
	http.HandleFunc("/admin", menuHandler.Admin)

	http.HandleFunc("/admin/categories", adminCatgHandler.AdminCategories)
	http.HandleFunc("/admin/categories/new", adminCatgHandler.AdminCategoriesNew)
	http.HandleFunc("/admin/categories/update", adminCatgHandler.AdminCategoriesUpdate)
	http.HandleFunc("/admin/categories/delete", adminCatgHandler.AdminCategoriesDelete)

	http.ListenAndServe(":8181", nil)
}
