/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package error_messages

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"fmt"
)

const(
	basePath = "../error_messages/"
)
type MessageLocalization struct {
	defaultLanguage *map[string]string
	currentLanguage *map[string]string
}

func NewMessageLocalization() (*MessageLocalization, error) {

	m := MessageLocalization{
	}
	defaultLanguage := new(map[string]string)
	defaultLanguage, err :=  m.getLanguageMessages("en", nil)
	if err != nil {
		return nil, err
	}

	m.defaultLanguage = defaultLanguage

	return &m, nil
}

func (m *MessageLocalization) GetMessage(key string) (*string, error) {

	cl :=  *m.currentLanguage

	if v, ok := cl[key]; ok {
		return &v, nil
	}

	cl = *m.defaultLanguage
	if v, ok := cl[key]; ok {
		return &v, nil
	}

	return nil, errors.New("No message found for " + key)
}


//GetComposeMessage : Get a message with one parameter. The parameter will replace the string
// "{1}" found in the message.
// @param key the key to the message
// @param p1 the parameter
// @return the message
func (m *MessageLocalization) GetComposedMessage(key string) string {
	return m.getComposedMessage(key, nil, nil, nil, nil)
}

//GetComposedMessageInt : Get a message with one parameter. The parameter will replace the string
// "{1}" found in the message.
// @param key the key to the message
// @param p1 the parameter
// @return the message
func (m *MessageLocalization) GetComposedMessageInt(key string, p1 int) string {

	str := strconv.Itoa(p1)
	return m.getComposedMessage(key, str, nil, nil, nil)
}

// GetComposedMessageWith2Value : Get a message with one parameter. The parameter will replace the string
// "{1}", "{2}" found in the message.
// @param key the key to the message
// @param p1 the parameter
// @param p2 the parameter
// @return the message

func (m *MessageLocalization) GetComposedMessageWith2Value(key string, p1 interface{}, p2 interface{}) string {

	return m.getComposedMessage(key, p1, p2, nil, nil)
}

// GetComposedMessageWith3Value : Get a message with one parameter. The parameter will replace the string
// "{1}", "{2}", "{3}" found in the message.
// @param key the key to the message
// @param p1 the parameter
// @param p2 the parameter
// @param p3 the parameter
// @return the message
func (m *MessageLocalization) GetComposedMessageWith3Value(key string, p1 interface{}, p2 interface{}, p3 interface{}) string {

	return m.getComposedMessage(key, p1, p2, p3, nil)
}

// getComposedMessage : Get a message with one parameter. The parameter will replace the string
// "{1}", "{2}", "{3}", "{4}" found in the message.
// @param key the key to the message
// @param p1 the parameter
// @param p2 the parameter
// @param p3 the parameter
// @param p4 the parameter
// @return the message
func (m *MessageLocalization) getComposedMessage(key string, p1 interface{}, p2 interface{}, p3 interface{}, p4 interface{}) string {

	template, err := m.GetMessage(key)
	if err != nil {
		return ""
	}
	msg := *template
	if p1 != nil {
		p1string := fmt.Sprint(p1)
		msg = strings.ReplaceAll(msg, "{1}", p1string)
	}
	if p2 != nil {
		p2string := fmt.Sprint(p2)
		msg = strings.ReplaceAll(msg, "{2}", p2string)
	}
	if p3 != nil {
		p3string := fmt.Sprint(p3)
		msg = strings.ReplaceAll(msg, "{3}", p3string)
	}
	if p4 != nil {
		p4string := fmt.Sprint(p4)
		msg = strings.ReplaceAll(msg, "{4}", p4string)
	}
	return msg
}

// Sets the language to be used globally for the error messages. The language
// is a two letter lowercase country designation like "en" or "pt". The country
// is an optional two letter uppercase code like "US" or "PT".
// @param language the language
// @param country the country
// @return true if the language was found, false otherwise
// @throws IOException on error
func (m *MessageLocalization) SetLanguage(language string, country *string) bool {

	lang, err := m.getLanguageMessages(language, country)
	if err != nil {
		return false
	}
	m.currentLanguage = lang

	return true
}

// Sets the error messages directly from a Reader.
// @param r the Reader
// @throws IOException on error

/*
func (m MessageLocalization) setMessages(r io.Reader) {
	m.currentLanguage = m.readLanguageStream(reader)
}
*/
func (m *MessageLocalization) getLanguageMessages(language string, country *string) (*map[string]string, error) {

	var fileName string
	if country != nil {
		fileName = language + "_" + *country + ".lng"
	}else{
		fileName = language + ".lng"
	}

	file, err := os.Open(basePath + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return m.readLanguageStream(file)
}

func (m *MessageLocalization) readLanguageStream(file *os.File) (*map[string]string, error) {

	lang := make(map[string]string)

	//buf := make([]byte, 32*1024) // define your buffer size here.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()
		idx := strings.Index(text, "=")

		if idx < 0 {
			continue
		}

		runes := []rune(text)
		key := string(runes[0:idx])

		if string(key[0]) == "#" {
			continue
		}
		value := string(runes[idx+1:])
		lang[key] = value

	}

	return &lang, nil
}
