package resp

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //system error
	ArgumentError ErrorCode = 1002 //argument error
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "system error",
		ArgumentError: "argument error",
	}
)
