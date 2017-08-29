// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TextAutonumberBullet struct {
	TypeAttr    ST_TextAutonumberScheme
	StartAtAttr *int32
}

func NewCT_TextAutonumberBullet() *CT_TextAutonumberBullet {
	ret := &CT_TextAutonumberBullet{}
	return ret
}
func (m *CT_TextAutonumberBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.StartAtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "startAt"},
			Value: fmt.Sprintf("%v", *m.StartAtAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextAutonumberBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "startAt" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.StartAtAttr = &pt
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextAutonumberBullet: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TextAutonumberBullet) Validate() error {
	return m.ValidateWithPath("CT_TextAutonumberBullet")
}
func (m *CT_TextAutonumberBullet) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_TextAutonumberSchemeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.StartAtAttr != nil {
		if *m.StartAtAttr < 1 {
			return fmt.Errorf("%s/m.StartAtAttr must be >= 1 (have %v)", path, *m.StartAtAttr)
		}
		if *m.StartAtAttr > 32767 {
			return fmt.Errorf("%s/m.StartAtAttr must be <= 32767 (have %v)", path, *m.StartAtAttr)
		}
	}
	return nil
}