package main

import (
	"fmt"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu/repository"
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

	catRepo := repository.NewCategoryGormRepo(dbconn)
	//ordeRepo := repository.NewOrderGormRepo(dbconn)
	menRepo := repository.NewItemGormRepo(dbconn)

	//ords, errs := ordeRepo.CustomerOrders(usr)

	cat01 := entity.Category{Name: "Lunch", Description: "Lorem Ipsum", Image: "lunch.png"}
	cat02 := entity.Category{Name: "Breakfast", Description: "Lorem Ipsum", Image: "breakfast.png"}

	c02, _ := catRepo.StoreCategory(&cat02)
	c01, _ := catRepo.StoreCategory(&cat01)

	cts1 := []entity.Category{*c01, *c02}
	cts2 := []entity.Category{*c01}

	men01 := entity.Item{Name: "Menu01", Price: 20.5, Description: "Lorem Ipsum", Image: "menu03.png", Categories: cts1}
	men02 := entity.Item{Name: "Menu02", Price: 20.5, Description: "Lorem Ipsum", Image: "menu04.png", Categories: cts2}

	menRepo.StoreItem(&men01)
	menRepo.StoreItem(&men02)
	men001, _ := catRepo.ItemsInCategory(c01)
	men002, _ := catRepo.ItemsInCategory(c02)
	fmt.Println(men001)
	fmt.Println(men002)
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
