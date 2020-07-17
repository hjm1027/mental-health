package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// Auth errors
	ErrAuthFailed   = &Errno{Code: 20101, Message: "The sid or password was incorrect."}
	ErrTokenInvalid = &Errno{Code: 20102, Message: "The token was invalid."}

	// user errors
	ErrCreateUser   = &Errno{Code: 20201, Message: "Error occurred in creating user."}
	ErrUpdateUser   = &Errno{Code: 20202, Message: "Error occurred in updating user"}
	ErrUserNotFound = &Errno{Code: 20203, Message: "The user was not found."}
	ErrGetUserInfo  = &Errno{Code: 20204, Message: "Error in getting user info"}
	ErrUserInfo     = &Errno{Code: 20205, Message: "The user information json cannot be null"}
)
