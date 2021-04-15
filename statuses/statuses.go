package statuses

type StatusCode string

const (
	OK StatusCode = "Operation completed successfully"

	UserRegFail StatusCode = "Couldn't register user"

	NoSuchUser StatusCode = "User with such login doesn't exists"

	IncorrectPassword StatusCode = "Incorrect password"

	NotLoggedIn StatusCode = "User isn't logged in"
)
