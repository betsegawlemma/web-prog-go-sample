package main

import (
	"html/template"
	"net/http"

	"github.com/betsegawlemma/webproggob/entity"
	"github.com/betsegawlemma/webproggob/menu/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var categoryService *service.CategoryService

func index(w http.ResponseWriter, r *http.Request) {

	categories, err := categoryService.Categories()
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "index.layout", categories)
}

func init() {
	categoryService = service.NewCategoryService("category.gob")
	categories := []entity.Category{
		entity.Category{ID: 1, Name: "Breakfast", Description: "Lorem ipsum", Image: "bkt.png"},
		entity.Category{ID: 2, Name: "Lunch", Description: "Lorem ipsum", Image: "lnc.png"},
		entity.Category{ID: 3, Name: "Dinner", Description: "Lorem ipsum", Image: "dnr.png"},
		entity.Category{ID: 4, Name: "Snack", Description: "Lorem ipsum", Image: "snk.png"},
	}
	categoryService.StoreCategories(categories)
}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8181", nil)
}
