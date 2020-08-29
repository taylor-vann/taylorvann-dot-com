package cache

import (
	"encoding/json"
	"errors"

	"webapi/store/cache"
	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/requests"
)

const (
	Role = "ROLE"
	Read = "READ"
)

var (
	errNilParams = errors.New("nil params given")
)

func getReadKey(userID int64, organization string) string {
	return Role + "_" + Read + "_" + string(userID) + "_" + organization
}

func GetReadEntry(p *requests.Read) (*controller.Roles, error) {
	if p == nil {
		return nil, errNilParams
	}
	key := getReadKey(p.UserID, p.Organization)
	entry, errReadEntry := cache.ReadEntry(&cache.ReadEntryParams{
		Environment: p.Environment,
		Key:         key,
	})
	if errReadEntry != nil || entry == nil {
		return nil, errReadEntry
	}
	if entry == nil {
		return nil, nil
	}

	var roles controller.Roles
	bytes, _ := json.Marshal(entry.Payload)
	errRolesUnmarshal := json.Unmarshal(bytes, &roles)

	return &roles, errRolesUnmarshal
}

func UpdateReadEntry(env string, roles *controller.Roles) error {
	role := (*roles)[0]

	key := getReadKey(role.UserID, role.Organization)
	_, errCreateEntry := cache.CreateEntry(&cache.CreateEntryParams{
		Environment: env,
		Key:         key,
		Payload:     *roles,
		Lifetime:    cache.ThreeDaysAsMS,
	})

	return errCreateEntry
}
