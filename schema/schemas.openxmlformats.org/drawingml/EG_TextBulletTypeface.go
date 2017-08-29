// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_TextBulletTypeface struct {
	BuFontTx *CT_TextBulletTypefaceFollowText
	BuFont   *CT_TextFont
}

func NewEG_TextBulletTypeface() *EG_TextBulletTypeface {
	ret := &EG_TextBulletTypeface{}
	return ret
}
func (m *EG_TextBulletTypeface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BuFontTx != nil {
		sebuFontTx := xml.StartElement{Name: xml.Name{Local: "a:buFontTx"}}
		e.EncodeElement(m.BuFontTx, sebuFontTx)
	}
	if m.BuFont != nil {
		sebuFont := xml.StartElement{Name: xml.Name{Local: "a:buFont"}}
		e.EncodeElement(m.BuFont, sebuFont)
	}
	return nil
}
func (m *EG_TextBulletTypeface) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextBulletTypeface:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "buFontTx":
				m.BuFontTx = NewCT_TextBulletTypefaceFollowText()
				if err := d.DecodeElement(m.BuFontTx, &el); err != nil {
					return err
				}
			case "buFont":
				m.BuFont = NewCT_TextFont()
				if err := d.DecodeElement(m.BuFont, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextBulletTypeface
		case xml.CharData:
		}
	}
	return nil
}
func (m *EG_TextBulletTypeface) Validate() error {
	return m.ValidateWithPath("EG_TextBulletTypeface")
}
func (m *EG_TextBulletTypeface) ValidateWithPath(path string) error {
	if m.BuFontTx != nil {
		if err := m.BuFontTx.ValidateWithPath(path + "/BuFontTx"); err != nil {
			return err
		}
	}
	if m.BuFont != nil {
		if err := m.BuFont.ValidateWithPath(path + "/BuFont"); err != nil {
			return err
		}
	}
	return nil
}