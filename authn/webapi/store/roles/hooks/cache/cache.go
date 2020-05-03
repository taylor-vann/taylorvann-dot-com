package cache

import (
	"encoding/json"

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
	key := getReadKey(p.UserID, p.Organization)
	entry, errReadEntry := cache.ReadEntry(&cache.ReadEntryParams{
		Environment: p.Environment,
		Key: key,
	})
	if errReadEntry != nil {
		return nil, errReadEntry
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