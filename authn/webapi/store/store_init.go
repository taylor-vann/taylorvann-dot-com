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

const initFilname = "/root/go/src/webapi/store_db.init.json"

func createOrReadUser(
	environment string,
	email string,
	password string,
) (usersController.SafeUsers, error) {
	// make user
	userRows, errUserRow := usersController.Read(
		&usersController.ReadParams{
			Environment: environment,
			Email:       email,
		},
	)

	if errUserRow != nil {
		return nil, errUserRow
	}

	// if user does not exist, create a one
	if len(userRows) == 0 {
		userRows, errUserRow = usersController.Create(
			&usersController.CreateParams{
				Environment: environment,
				Email:       email,
				Password:    password,
			},
		)
	}

	return userRows, errUserRow
}

// create roles
func createRoles(
	environment string,
	userRow usersController.SafeRow,
	roles []string,
) {
	for _, organization := range roles {
		rolesController.Create(
			&rolesController.CreateParams{
				Environment:  environment,
				UserID:       userRow.ID,
				Organization: organization,
				ReadAccess:   true,
				WriteAccess:  true,
			},
		)
	}
}

func createUserAndRoles(
	environment string,
	details *InitDetails,
) {
	for email, userDetails := range details.Users {
		// read user, they might already exist
		userRows, errUserRows := createOrReadUser(
			environment,
			email,
			userDetails.Password,
		)
		if errUserRows != nil {
			continue
		}

		// create roles for user
		userRow := userRows[0]
		createRoles(environment, userRow, userDetails.Roles)
	}
}

func InitFromJSON() {
	initJSON, errInitFile := ioutil.ReadFile(initFilname)
	if errInitFile != nil {
		return
	}

	var initDetails InitDetails
	errInitDetails := json.Unmarshal(initJSON, &initDetails)
	if errInitDetails != nil {
		return
	}

	createUserAndRoles("DEVELOPMENT", &initDetails)
	createUserAndRoles("PRODUCTION", &initDetails)
}
