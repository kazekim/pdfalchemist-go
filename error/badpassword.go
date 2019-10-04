/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package error

const badPasswordSerialVersionUID = "-4333706268155063964"

type BadPasswordError struct {
	message string
}

func NewBadPasswordError(message string) error {
	return &BadPasswordError{
		message: message,
	}
}

func (e BadPasswordError) Error() string {
	return e.message
}
