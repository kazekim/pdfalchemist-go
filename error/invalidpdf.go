/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package error

const invalidPdfSerialVersionUID = "-2319614911517026938L"

type InvalidPdfError struct {
	message string
}

func NewInvalidPdfError(message string) error {
	return &InvalidPdfError{
		message: message,
	}
}

func (e InvalidPdfError) Error() string {
	return e.message
}

