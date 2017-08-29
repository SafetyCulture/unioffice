// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CustomXmlPr struct {
	// Custom XML Element Placeholder Text
	Placeholder *CT_String
	// Custom XML Attribute
	Attr []*CT_Attr
}

func NewCT_CustomXmlPr() *CT_CustomXmlPr {
	ret := &CT_CustomXmlPr{}
	return ret
}
func (m *CT_CustomXmlPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Placeholder != nil {
		seplaceholder := xml.StartElement{Name: xml.Name{Local: "w:placeholder"}}
		e.EncodeElement(m.Placeholder, seplaceholder)
	}
	if m.Attr != nil {
		seattr := xml.StartElement{Name: xml.Name{Local: "w:attr"}}
		e.EncodeElement(m.Attr, seattr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CustomXmlPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomXmlPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "placeholder":
				m.Placeholder = NewCT_String()
				if err := d.DecodeElement(m.Placeholder, &el); err != nil {
					return err
				}
			case "attr":
				tmp := NewCT_Attr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Attr = append(m.Attr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomXmlPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CustomXmlPr) Validate() error {
	return m.ValidateWithPath("CT_CustomXmlPr")
}
func (m *CT_CustomXmlPr) ValidateWithPath(path string) error {
	if m.Placeholder != nil {
		if err := m.Placeholder.ValidateWithPath(path + "/Placeholder"); err != nil {
			return err
		}
	}
	for i, v := range m.Attr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Attr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}