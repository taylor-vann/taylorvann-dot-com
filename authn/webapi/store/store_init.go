package store

import (
	"encoding/json"
	"io/ioutil"

	rolesController "webapi/store/roles/controller"
	usersController "webapi/store/users/controller"
)

type InitUserDetails struct {
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type InitDetails struct {
	Users map[string]InitUserDetails `json:"users"`
}

const initFilname = "./store_db.init.json"

func InitFromJSON() {
	initJSON, _ := ioutil.ReadFile(initFilname)
	var initDetails InitDetails
	err := json.Unmarshal(initJSON, &initDetails)

	if err != nil {
		// TODO: log a failure, but first need logging system
		return
	}

	for email, details := range initDetails.Users {
		userRows, errUserRow := usersController.Create(&usersController.CreateParams{
			Email:    email,
			Password: details.Password,
		})

		if errUserRow != nil {
			// TODO: log a failure
			continue
		}

		if len(userRows) == 0 {
			continue
		}

		userRow := userRows[0]
		for _, organization := range details.Roles {
			_, errRoleRow := rolesController.Create(&rolesController.CreateParams{
				UserID: userRow.ID,
				Organization: organization,
			})
			if errRoleRow != nil {
				// TODO: log failure
			}
		}
	}
}
