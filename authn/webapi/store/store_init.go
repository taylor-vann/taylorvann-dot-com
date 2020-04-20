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

// func getConfigFromEnv() (*redisx.RedisConfig, error) {
// 	config := redisx.RedisConfig{
// 		Host:        constants.Env.Host,
// 		Port:        constants.Env.Port,
// 		Protocol:    constants.Env.Protocol,
// 		MaxIdle:     constants.Env.MaxIdle,
// 		IdleTimeout: constants.Env.IdleTimeout,
// 		MaxActive:   constants.Env.MaxActive,
// 	}

// 	return &config, nil
// }

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
		userRows, errUserRow := users.Create(&users.CreateParams{
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
			_, errRoleRow := roles.Create(&roles.CreateParams{
				UserID: userRow.ID,
				Organization: organization,
			})
			if errRoleRow != nil {
				// TODO: log failure
			}
		}
	}
}
