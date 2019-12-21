package main

import (
	"github.com/betsegawlemma/restaurant/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Item{}, &entity.Order{}, &entity.Category{}, &entity.User{}, &entity.Role{}, &entity.Ingredient{}, &entity.Comment{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	// createTables(dbconn)

	/*
		tmpl := template.Must(template.ParseGlob("ui/templates/*"))

		categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
		categoryServ := service.NewCategoryServiceImpl(categoryRepo)

		roleRepo := repository.NewRoleRepositoryImpl(dbconn)
		roleSrv := service.NewRoleServiceImpl(roleRepo)

		adminCatgHandler := handler.NewAdminCategoryHandler(tmpl, categoryServ)
		menuHandler := handler.NewMenuHandler(tmpl, categoryServ)
		roleHandler := handler.NewAdminRoleHandler(tmpl, roleSrv)

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

		http.HandleFunc("/admin/roles/new", roleHandler.AdminRolesNew)

		http.ListenAndServe(":8181", nil)
	*/
}
