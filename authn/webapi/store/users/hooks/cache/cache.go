package cache

import (
	"encoding/json"

	"webapi/store/cache"
	"webapi/store/users/controller"
	"webapi/store/users/hooks/requests"
)

const (
	User = "USER"
	Read = "READ"
)


func getReadKey(email string) string {
	return User + "_" + Read + "_" + email
}

func GetReadEntry(p *requests.Read) (*controller.SafeUsers, error) {
	key := getReadKey(p.Email)
	entry, errReadEntry := cache.ReadEntry(&cache.ReadEntryParams{
		Environment: p.Environment,
		Key: key,
	})
	if errReadEntry != nil {
		return nil, errReadEntry
	}
	if entry == nil {
		return nil, nil
	}

	bytes, _ := json.Marshal(entry.Payload)
	var users controller.SafeUsers
	errUsersUnmarshal := json.Unmarshal(bytes, &users)
	
	return &users, errUsersUnmarshal
}

func UpdateReadEntry(env string, users *controller.SafeUsers) (error) {
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