package store

import (
	"encoding/json"
	"io/ioutil"

	"webapi/controllers/roles"
	"webapi/controllers/users"
)

type InitUserDetails struct {
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

// define
type InitDetails struct {
	Users map[string]InitUserDetails `json:"users"`
}

const initFilname = "./store_db.init.json"

// InitFromJSON -
func InitFromJSON() {
	initJSON, _ := ioutil.ReadFile(initFilname)
	var initDetails InitDetails
	err := json.Unmarshal(initJSON, &initDetails)

	if err != nil {
		// TODO: log a failure, but first need logging system
		return
	}

	for email, details := range initDetails.Users {
		userRow, errUserRow := users.Create(&users.CreateParams{
			Email:    email,
			Password: details.Password,
		})

		if errUserRow != nil {
			// TODO: log a failure
			continue
		}

		if userRow != nil {
			for _, role := range details.Roles {
				_, errRoleRow := roles.Create(&roles.CreateParams{
					UserID: userRow.ID,
					Role:   role,
				})
				if errRoleRow != nil {
					// TODO: log failure
				}
			}
		}
	}
}
