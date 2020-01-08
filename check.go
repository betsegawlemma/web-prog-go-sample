package main

import (
	"fmt"
	"strings"
)

func main() {
	roles := []string{"ADMIN", "USER"}
	if !checkRole("ANY", roles) {
		fmt.Println("Role is not found")
	} else {

		fmt.Println("role is found")
	}
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}
