package store

// This is a temporary package designed to get store off the ground
// once initial users are created, it is uneccessary.

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
	initJSON, errInitFile := ioutil.ReadFile(initFilname)
	if errInitFile != nil {
		// TODO: log a failure, but first need logging system
		return
	}

	var initDetails InitDetails
	err := json.Unmarshal(initJSON, &initDetails)

	if err != nil {
		// TODO: log a failure, but first need logging system
		return
	}

	for email, details := range initDetails.Users {
		userRows, errUserRow := usersController.Read(&usersController.ReadParams{
			Environment: Environment,
			Email:    email,
		})

		if len(userRows) == 0 || errUserRow != nil {
			userRows, errUserRow = usersController.Create(&usersController.CreateParams{
				Environment: Environment,
				Email:    email,
				Password: details.Password,
			})
		}

		if errUserRow != nil {
			continue
		}

		if len(userRows) == 0 {
			continue
		}

		userRow := userRows[0]
		for _, organization := range details.Roles {
			_, errRoleRow := rolesController.Create(&rolesController.CreateParams{
				Environment: Environment,
				UserID: userRow.ID,
				Organization: organization,
				ReadAccess: true,
				WriteAccess: true,
			})
			if errRoleRow != nil {
				// TODO: log failure
				continue
			}
		}
	}
}
