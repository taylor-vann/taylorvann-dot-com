package constants

const (
	Guest            			= "guest"
	Infra									= "infra"
	Client           			= "client"
	Session         		  = "session"
	BrianTaylorVannDotCom = "briantaylorvann.com"

	CreateAccount    = "create_account"
	UpdatePassword   = "update_password"
	UpdateEmail		   = "update_email"
	DeleteAccount		 = "delete_account"
)

var (
	OneDayAsMS = int64(1000 * 60 * 60 * 24)
	ThreeDaysAsMS = 3 * OneDayAsMS
	ThreeSixtyFiveDaysAsMS = 365 * OneDayAsMS
)