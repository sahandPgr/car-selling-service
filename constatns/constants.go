package constatns

const (
	//Auth header
	AuthorizationHeaderKey string = "Authorization"
	//User
	DefaultUserRole    string = "default"
	AdminRole          string = "admin"
	AdminName          string = "admin"
	RedisOtpDefaultKey string = "otp"

	// Claims
	UserIdKey     string = "UserId"
	UsernameKey   string = "Username"
	FirstName     string = "FirstName"
	LastName      string = "LastName"
	Email         string = "Email"
	PhoneNumber   string = "PhoneNumber"
	Roles         string = "Roles"
	ExpireTimeKey string = "Exp"
	//DB
	GetRecordNotNullQuery string = "id = ? AND deleted_by is null"
	NotNullQuery          string = "deleted_by is null"
	//Dir
	UploadDirName string = "uploads"
)
