package cache

import (
	"encoding/json"
	"errors"

	"log"

	"webapi/store/cache"
	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/requests"
)

const Role = "ROLE"
const Read = "READ"

func getReadKey(userID int64, organization string) string {
	return Role + "_" + Read + "_" + string(userID) + "_" + organization
}

func GetReadEntry(p *requests.Read) (*controller.Roles, error) {
	log.Println("ROLES CACHE -  get read entry")
	log.Println(p)

	if p == nil {
		log.Println("ROLES CACHE -  nil params")

		return nil, errors.New("nil params given")
	}
	key := getReadKey(p.UserID, p.Organization)
	entry, errReadEntry := cache.ReadEntry(&cache.ReadEntryParams{
		Environment: p.Environment,
		Key: key,
	})
	if errReadEntry != nil || entry == nil {
		return nil, errReadEntry
	}

	log.Println("ROLES CACHE -  get read entry")
	log.Println(entry)
	if entry == nil {
		return nil, nil
	}

	bytes, _ := json.Marshal(entry.Payload)
	var roles controller.Roles
	errRolesUnmarshal := json.Unmarshal(bytes, &roles)

	return &roles, errRolesUnmarshal
}

func UpdateReadEntry(env string, roles *controller.Roles) (error) {
	role := (*roles)[0]

	key := getReadKey(role.UserID, role.Organization)
	_, errCreateEntry := cache.CreateEntry(&cache.CreateEntryParams{
		Environment: env,
		Key: key,
		Payload: *roles,
		Lifetime: cache.ThreeDaysAsMS,
	})
	
	return errCreateEntry	
}