package cache

import (
	"encoding/json"

	"webapi/store/cache"
	"webapi/store/users/controller"
	"webapi/store/users/hooks/requests"
)

const User = "USER"
const Read = "READ"

func getReadKey(email string) string {
	return User + "_" + Read + "_" + email
}

func GetReadEntry(p *requests.Read) (controller.Users, error) {
	key := getReadKey(p.Email)
	entry, errReadEntry := cache.ReadEntry(&cache.ReadEntryParams{
		Environment: p.Environment,
		Key: key,
	})
	if errReadEntry != nil {
		return nil, errReadEntry
	}
	if entry != nil {
		bytes, _ := json.Marshal(entry.Payload)
		var users controller.Users
		errUsersUnmarshal := json.Unmarshal(bytes, &users)
		if errUsersUnmarshal != nil {
			return nil, errUsersUnmarshal
		}
		return users, nil
	}
	
	return controller.Read(p)
}

func UpdateReadEntry(env string, users *controller.Users) (error) {
	userID := (*users)[0].Email

	key := getReadKey(userID)
	_, errCreateEntry := cache.CreateEntry(&cache.CreateEntryParams{
		Environment: env,
		Key: key,
		Payload: *users,
		Lifetime: cache.ThreeDaysAsMS,
	})
	
	return errCreateEntry	
}