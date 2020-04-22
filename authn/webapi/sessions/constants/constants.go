package constants

const (
	Guest            = "guest"
	Public           = "public"
	Document         = "document"
	Session          = "session"
	CreateAccount    = "create_account"
	UpdatePassword   = "update_password"
	UpdateEmail		   = "update_email"
	DeleteAccount		 = "delete_account"
	TaylorVannDotCom = "taylorvann.com"
)

var (
	OneDayAsMS = int64(1000 * 60 * 60 * 24)
	ThreeDaysAsMS = 3 * OneDayAsMS
)