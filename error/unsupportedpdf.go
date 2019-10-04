/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package error

const unsupportedSerialVersionUID = "-2180764250839096628L"

type UnsupportedPdfError struct {
	message string
}

func NewUnsupportedPdfError(message string) error {
	return &UnsupportedPdfError{
		message: message,
	}
}

func (e UnsupportedPdfError) Error() string {
	return e.message
}


