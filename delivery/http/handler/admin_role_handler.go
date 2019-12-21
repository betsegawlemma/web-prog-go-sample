package handler

import (
	"html/template"
	"net/http"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu"
)

// AdminRoleHandler handles roles related requests
type AdminRoleHandler struct {
	tmpl    *template.Template
	roleSrv menu.RoleService
}

// NewAdminRoleHandler returns a new AdminRoleHandler object
func NewAdminRoleHandler(t *template.Template, rs menu.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{tmpl: t, roleSrv: rs}
}

// AdminRolesNew handles routes at /admin/roles/new
func (arh AdminRoleHandler) AdminRolesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		role := &entity.Role{Name: name}

		err := arh.roleSrv.StoreRole(role)

		if err != nil {
			panic(err)
		}
	} else {
		arh.tmpl.ExecuteTemplate(w, "admin.roles.new.layout", nil)
	}

	//http.Redirect(w, r, "/admin/roles", http.StatusSeeOther)
}
