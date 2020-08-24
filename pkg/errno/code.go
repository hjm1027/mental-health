package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrDatabase         = &Errno{Code: 10002, Message: "Database error."}
	ErrBind             = &Errno{Code: 10003, Message: "Error occurred while binding the request body to the struct."}
	ErrToken            = &Errno{Code: 10004, Message: "Error occurred while signing the JSON web token."}
	ErrGetQuery         = &Errno{Code: 10005, Message: "Error occurred while getting query. "}
	ErrGetParam         = &Errno{Code: 10006, Message: "Error occurred while getting path params. "}

	// Auth errors
	ErrAuthFailed   = &Errno{Code: 20101, Message: "The sid or password was incorrect."}
	ErrTokenInvalid = &Errno{Code: 20102, Message: "The token was invalid."}

	// user errors
	ErrCreateUser   = &Errno{Code: 20201, Message: "Error occurred in creating user."}
	ErrUpdateUser   = &Errno{Code: 20202, Message: "Error occurred in updating user"}
	ErrUserNotFound = &Errno{Code: 20203, Message: "The user was not found."}
	ErrGetUserInfo  = &Errno{Code: 20204, Message: "Error in getting user info"}
	ErrUserInfo     = &Errno{Code: 20205, Message: "The user information json cannot be null"}

	//mood errors
	ErrGetScoreInfo = &Errno{Code: 20301, Message: "Error in getting mood score."}
	ErrGetNoteInfo  = &Errno{Code: 20302, Message: "Error in getting mood note."}

	//hole errors
	ErrWordLimitation = &Errno{Code: 20401, Message: "Word limit exceeded"}
	ErrGetHoleInfo    = &Errno{Code: 20402, Message: "Error occurred while getting hole."}
	ErrNotLiked       = &Errno{Code: 20403, Message: "User has not liked yet. "}
	ErrHasLiked       = &Errno{Code: 20404, Message: "User has already liked. "}
	ErrNotFavorited   = &Errno{Code: 20405, Message: "User has not favorited yet. "}
	ErrHasFavorited   = &Errno{Code: 20406, Message: "User has already favorited. "}

	ErrGetSubCommentInfo    = &Errno{Code: 20501, Message: "Error occurred while getting subComment info"}
	ErrGetParentCommentInfo = &Errno{Code: 20502, Message: "Error occurred while getting parent comment info"}
	ErrCommentList          = &Errno{Code: 20503, Message: "Error occurred while getting comment list. "}

	//search errors
	ErrSearchCourse = &Errno{Code: 20601, Message: "Error occured in searching courses."}
	ErrSearchHole   = &Errno{Code: 20602, Message: "Error occured in searching holes."}

	// message errors
	ErrGetMessage = &Errno{Code: 20701, Message: "Error occurred in getting message list"}

	// upload errors
	ErrGetFile    = &Errno{Code: 20801, Message: "Error occurred in getting file from FormFile()"}
	ErrUploadFile = &Errno{Code: 20802, Message: "Error occurred in uploading file to oss"}

	// poster errors
	ErrGetPoster = &Errno{Code: 20901, Message: "Error occurred while getting poster."}

	// reserve errors
	ErrQueryReserve         = &Errno{Code: 21001, Message: "Error occurred while query reserve."}
	ErrGetTeacherList       = &Errno{Code: 21002, Message: "Error occurred while getting teacher list."}
	ErrGetTeacherBySchedule = &Errno{Code: 21003, Message: "Error occurred while getting teacher by schedule."}
	ErrCreateReserve        = &Errno{Code: 21004, Message: "Error occurred in creating reserve."}
	ErrSendMail             = &Errno{Code: 21005, Message: "Error occurred while sending mail."}
	ErrGetStatus            = &Errno{Code: 21006, Message: "Error occurred while geting reserve status."}
	ErrCheckReserve         = &Errno{Code: 21007, Message: "Error occurred while checking reserve."}
)
