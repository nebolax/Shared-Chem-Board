package status

type StatusCode string

const (
	OK StatusCode = "Operation completed successfully"

	UserRegFail StatusCode = "Couldn't register user"

	IncorrectLogPass StatusCode = "Incorrect password"

	UserAlreadyExists StatusCode = "User with such nickname or password is already registered"

	NotLoggedIn StatusCode = "User isn't logged in"

	NoSuchBoard StatusCode = "Board with this is doesn't exist"
)
