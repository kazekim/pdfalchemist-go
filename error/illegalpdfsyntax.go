/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package error

const illegalPdfSyntaxSerialVersionUID = "-643024246596031671L"

type IllegalPdfSyntaxError struct {
	message string
}

func NewIllegalPdfSyntaxError(message string) error {
	return &IllegalPdfSyntaxError{
		message: message,
	}
}

func (e IllegalPdfSyntaxError) Error() string {
	return e.message
}

