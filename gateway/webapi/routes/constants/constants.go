package constants

type DomainDetails struct {
	Name					string
	SubDomain			string
	TargetAddress string
}

type DomainDetailsMap = map[string]DomainDetails

const (
	authn		= "authn"
	home 		= "taylorvann"
	statics = "statics"
)

var RouteMap = createDomainDetailsMap()

func createDomainDetailsMap() *DomainDetailsMap {
	domains := make(DomainDetailsMap)
	
	domains[home] = DomainDetails{
		Name:					 "home",
		SubDomain:		 home,
		TargetAddress: "https://127.0.0.1:3000/",
	}

	domains[statics] = DomainDetails{
		Name:					 "statics",
		SubDomain:		 statics,
		TargetAddress: "https://127.0.0.1:4000/",
	}

	domains[authn] = DomainDetails{
		Name:					 "authentication",
		SubDomain:		 authn,
		TargetAddress: "https://127.0.0.1:5000/",
	}


	return &domains
}