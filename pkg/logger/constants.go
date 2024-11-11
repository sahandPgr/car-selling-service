package logger

type Category string
type SubCategory string
type ExtraKey string

// Catergory constants
const (
	Potgres         Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Internal        Category = "Internal"
	General         Category = "General"
)

// SubCategory constants
const (
	//Postgres
	Select SubCategory = "Select"
	Insert SubCategory = "Insert"
	Update SubCategory = "Update"
	Delete SubCategory = "Delete"
	//Redis
	Set SubCategory = "Set"
	Get SubCategory = "Get"
	// Validation
	MobileNumberValidation SubCategory = "MobileNumberValidation"
	PasswordValidation     SubCategory = "PasswordValidation"
	// Internal
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"
)

// ExtraKey constants
const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIP     ExtraKey = "ClientIP"
	HostIP       ExtraKey = "HostIP"
	Method       ExtraKey = "Method"
	Path         ExtraKey = "Path"
	StatusCode   ExtraKey = "StatusCode"
	Latency      ExtraKey = "Latency"
	BodySize     ExtraKey = "BodySize"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
)
