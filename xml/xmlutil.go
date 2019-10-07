/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package xml


func escapeXML(s string, onlyASCII bool) string {

	var sb string
	for _, r := range s {
		c := string(r)

		switch c {
		case "<":
			sb = sb + "&lt;"
		case ">":
			sb = sb + "&gt;"
		case "&":
			sb = sb + "&amp;"
		case "\"":
			sb = sb + "&quot;"
		case "'":
			sb = sb + "&apos;"
		default:
			if r == 0x9 || r == 0xA || r == 0xD ||
				( r >= 0x20 && r <= 0xD7FF ) ||
				( r >= 0xE000 && r <= 0xFFFD) ||
				( r >= 0x10000 && r <= 0x10FFFF) {
				if onlyASCII && r > 127 {
					sb = sb + "&#" + r + ";"
				}else{
					sb = sb + c
				}
			}
		}
		cs = append(cs, c)
	}

	len := len(cs)

	for
}