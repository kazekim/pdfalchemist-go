/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package xml

import "github.com/kazekim/pdfalchemist-go"

type XMLPeer struct {
	tagName string
	customTagName string
	attributeAliases map[string]string
	attributeValues map[string]string
	defaultContent *string
}

func NewXMLPeer(name string, alias string) *XMLPeer {
	return &XMLPeer{
		tagName: name,
		customTagName: alias,
		attributeAliases: make(map[string]string),
		attributeValues: make(map[string]string),
	}
}

func (p *XMLPeer) Tag() string {
	return p.tagName
}

func (p *XMLPeer) Alias() string {
	return p.customTagName
}

func (p *XMLPeer) Attributes(attrs *Attributes) map[string]string {
	var attributes map[string]string

	if len(p.attributeValues) > 0 {
		for k, v := range p.attributeValues {
			attributes[k] = v
		}
	}

	if p.defaultContent != nil {
		attributes[pdfalchemist.PDFAlchemist] = *p.defaultContent
	}

	if attrs != nil {
		as := *attrs
		for i := 0; i < as.Length() ; i++ {
			attribute := p.Name(as.QName(i))
			attributes[attribute] = as.Value(i)
		}
	}

	return attributes
}

func (p *XMLPeer) AddAlias(name string, alias string) {
	p.attributeAliases[alias] = name
}

func (p *XMLPeer) AddValue(name string, value string) {
	p.attributeValues[name] = value
}

func (p *XMLPeer) SetContent(content string) {
	p.defaultContent = &content
}

func (p *XMLPeer) Name(name string) string {
	value := name
	if v, ok := p.attributeAliases[name] ; ok {
		value = v
	}
	return value
}

func (p *XMLPeer) DefaultValues() map[string]string {
	return p.attributeValues
}