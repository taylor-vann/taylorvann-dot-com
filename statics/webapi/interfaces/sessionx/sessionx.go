package sessionx


func getLifetimeByAudience(audience string) int64 {
	switch audience{
	case constants.Guest:
		return constants.OneDayAsMS
	case constants.Public:
		return constants.ThreeDaysAsMS
	default:
		return constants.OneDayAsMS
	}
}

// verify jwt

// if available
// fetch updated document / guest session through http request

// else return new with document / session through http request
