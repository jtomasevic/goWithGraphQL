package errors

type WrongPasswordError struct{}

func (m *WrongPasswordError) Error() string {
	return "wrong username"
}