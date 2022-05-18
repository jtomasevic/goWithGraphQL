package errors

type WrongUsernameError struct{}

func (m *WrongUsernameError) Error() string {
	return "wrong username"
}