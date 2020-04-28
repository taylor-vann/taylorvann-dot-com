package constants

const (
	Guest            = "guest"
	Public           = "public"
	Document         = "document"
	Session          = "session"
	TaylorVannDotCom = "taylorvann.com"

	CreateAccount    = "create_account"
	UpdatePassword   = "update_password"
	UpdateEmail		   = "update_email"
	DeleteAccount		 = "delete_account"
)

var (
	OneDayAsMS = int64(1000 * 60 * 60 * 24)
	ThreeDaysAsMS = 3 * OneDayAsMS
)