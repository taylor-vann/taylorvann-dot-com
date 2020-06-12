package constants

const (
	Guest            			= "guest"
	Infra									= "infra"
	Client           			= "client"
	Session         		  = "session"
	CreateAccount					= "create_account"
	UpdateEmail           = "update_email"
	UpdatePassword        = "update_password"
	DeleteAccount        	= "delete_account"

	BrianTaylorVannDotCom = "briantaylorvann.com"
)

var (
	OneDayAsMS = int64(1000 * 60 * 60 * 24)
	ThreeDaysAsMS = 3 * OneDayAsMS
	SevenDaysAsMS = 7 * OneDayAsMS
	ThreeSixtyFiveDaysAsMS = 365 * OneDayAsMS
)