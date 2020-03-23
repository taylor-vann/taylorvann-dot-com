package constants

type DomainDetails struct {
	Name		string
	RequestedAddress	string
	TargetAddress string
}

type DomainDetailsMap = map[string]DomainDetails


const (
	taylorvann = "https://www.taylorvann.com/"
	devTaylorvann = "https://127.0.0.1/"

	statics = "https://www.statics.taylorvann.com/"
	devStatics = "https://127.0.0.1:3005/statics/"
	
	authn = "https://www.authn.taylorvann.com/"
	devAuthn = "https://127.0.0.1:3005/authn/"
)

var RouteMap = createDomainDetailsMap()

func createDomainDetailsMap() *DomainDetailsMap {
	domains := make(DomainDetailsMap)

	domains[taylorvann] = DomainDetails{
		Name:							"home",
		RequestedAddress:	taylorvann,
		TargetAddress:		"https://127.0.0.1:4000/",
	}
	
	domains[devTaylorvann] = DomainDetails{
		Name:							"dev.home",
		RequestedAddress:	devTaylorvann,
		TargetAddress:		"https://127.0.0.1:4000/",
	}
	
	domains[authn] = DomainDetails{
		Name:							"authn",
		RequestedAddress:	authn,
		TargetAddress:		"https://127.0.0.1:5000/",
	}
	
	domains[devAuthn] = DomainDetails{
		Name:							"dev.authn",
		RequestedAddress:	devAuthn,
		TargetAddress:		"https://127.0.0.1:5000/",
	}
	
	domains[statics] = DomainDetails{
		Name:							"statics",
		RequestedAddress:	statics,
		TargetAddress:		"https://127.0.0.1:4000/",
	}
	
	domains[devStatics] = DomainDetails{
		Name:							"dev.statics",
		RequestedAddress:	devStatics,
		TargetAddress:		"https://127.0.0.1:4000/",
	}

	return &domains
}