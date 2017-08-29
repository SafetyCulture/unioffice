// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_DateAx struct {
	AxId           *CT_UnsignedInt
	Scaling        *CT_Scaling
	Delete         *CT_Boolean
	AxPos          *CT_AxPos
	MajorGridlines *CT_ChartLines
	MinorGridlines *CT_ChartLines
	Title          *CT_Title
	NumFmt         *CT_NumFmt
	MajorTickMark  *CT_TickMark
	MinorTickMark  *CT_TickMark
	TickLblPos     *CT_TickLblPos
	SpPr           *drawingml.CT_ShapeProperties
	TxPr           *drawingml.CT_TextBody
	CrossAx        *CT_UnsignedInt
	Choice         *EG_AxSharedChoice
	Auto           *CT_Boolean
	LblOffset      *CT_LblOffset
	BaseTimeUnit   *CT_TimeUnit
	MajorUnit      *CT_AxisUnit
	MajorTimeUnit  *CT_TimeUnit
	MinorUnit      *CT_AxisUnit
	MinorTimeUnit  *CT_TimeUnit
	ExtLst         *CT_ExtensionList
}

func NewCT_DateAx() *CT_DateAx {
	ret := &CT_DateAx{}
	ret.AxId = NewCT_UnsignedInt()
	ret.Scaling = NewCT_Scaling()
	ret.AxPos = NewCT_AxPos()
	ret.CrossAx = NewCT_UnsignedInt()
	return ret
}
func (m *CT_DateAx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seaxId := xml.StartElement{Name: xml.Name{Local: "axId"}}
	e.EncodeElement(m.AxId, seaxId)
	sescaling := xml.StartElement{Name: xml.Name{Local: "scaling"}}
	e.EncodeElement(m.Scaling, sescaling)
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	seaxPos := xml.StartElement{Name: xml.Name{Local: "axPos"}}
	e.EncodeElement(m.AxPos, seaxPos)
	if m.MajorGridlines != nil {
		semajorGridlines := xml.StartElement{Name: xml.Name{Local: "majorGridlines"}}
		e.EncodeElement(m.MajorGridlines, semajorGridlines)
	}
	if m.MinorGridlines != nil {
		seminorGridlines := xml.StartElement{Name: xml.Name{Local: "minorGridlines"}}
		e.EncodeElement(m.MinorGridlines, seminorGridlines)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.MajorTickMark != nil {
		semajorTickMark := xml.StartElement{Name: xml.Name{Local: "majorTickMark"}}
		e.EncodeElement(m.MajorTickMark, semajorTickMark)
	}
	if m.MinorTickMark != nil {
		seminorTickMark := xml.StartElement{Name: xml.Name{Local: "minorTickMark"}}
		e.EncodeElement(m.MinorTickMark, seminorTickMark)
	}
	if m.TickLblPos != nil {
		setickLblPos := xml.StartElement{Name: xml.Name{Local: "tickLblPos"}}
		e.EncodeElement(m.TickLblPos, setickLblPos)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	secrossAx := xml.StartElement{Name: xml.Name{Local: "crossAx"}}
	e.EncodeElement(m.CrossAx, secrossAx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.Auto != nil {
		seauto := xml.StartElement{Name: xml.Name{Local: "auto"}}
		e.EncodeElement(m.Auto, seauto)
	}
	if m.LblOffset != nil {
		selblOffset := xml.StartElement{Name: xml.Name{Local: "lblOffset"}}
		e.EncodeElement(m.LblOffset, selblOffset)
	}
	if m.BaseTimeUnit != nil {
		sebaseTimeUnit := xml.StartElement{Name: xml.Name{Local: "baseTimeUnit"}}
		e.EncodeElement(m.BaseTimeUnit, sebaseTimeUnit)
	}
	if m.MajorUnit != nil {
		semajorUnit := xml.StartElement{Name: xml.Name{Local: "majorUnit"}}
		e.EncodeElement(m.MajorUnit, semajorUnit)
	}
	if m.MajorTimeUnit != nil {
		semajorTimeUnit := xml.StartElement{Name: xml.Name{Local: "majorTimeUnit"}}
		e.EncodeElement(m.MajorTimeUnit, semajorTimeUnit)
	}
	if m.MinorUnit != nil {
		seminorUnit := xml.StartElement{Name: xml.Name{Local: "minorUnit"}}
		e.EncodeElement(m.MinorUnit, seminorUnit)
	}
	if m.MinorTimeUnit != nil {
		seminorTimeUnit := xml.StartElement{Name: xml.Name{Local: "minorTimeUnit"}}
		e.EncodeElement(m.MinorTimeUnit, seminorTimeUnit)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DateAx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AxId = NewCT_UnsignedInt()
	m.Scaling = NewCT_Scaling()
	m.AxPos = NewCT_AxPos()
	m.CrossAx = NewCT_UnsignedInt()
lCT_DateAx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "axId":
				if err := d.DecodeElement(m.AxId, &el); err != nil {
					return err
				}
			case "scaling":
				if err := d.DecodeElement(m.Scaling, &el); err != nil {
					return err
				}
			case "delete":
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case "axPos":
				if err := d.DecodeElement(m.AxPos, &el); err != nil {
					return err
				}
			case "majorGridlines":
				m.MajorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MajorGridlines, &el); err != nil {
					return err
				}
			case "minorGridlines":
				m.MinorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MinorGridlines, &el); err != nil {
					return err
				}
			case "title":
				m.Title = NewCT_Title()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "majorTickMark":
				m.MajorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MajorTickMark, &el); err != nil {
					return err
				}
			case "minorTickMark":
				m.MinorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MinorTickMark, &el); err != nil {
					return err
				}
			case "tickLblPos":
				m.TickLblPos = NewCT_TickLblPos()
				if err := d.DecodeElement(m.TickLblPos, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case "crossAx":
				if err := d.DecodeElement(m.CrossAx, &el); err != nil {
					return err
				}
			case "crosses":
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.Crosses, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "crossesAt":
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.CrossesAt, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "auto":
				m.Auto = NewCT_Boolean()
				if err := d.DecodeElement(m.Auto, &el); err != nil {
					return err
				}
			case "lblOffset":
				m.LblOffset = NewCT_LblOffset()
				if err := d.DecodeElement(m.LblOffset, &el); err != nil {
					return err
				}
			case "baseTimeUnit":
				m.BaseTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.BaseTimeUnit, &el); err != nil {
					return err
				}
			case "majorUnit":
				m.MajorUnit = NewCT_AxisUnit()
				if err := d.DecodeElement(m.MajorUnit, &el); err != nil {
					return err
				}
			case "majorTimeUnit":
				m.MajorTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.MajorTimeUnit, &el); err != nil {
					return err
				}
			case "minorUnit":
				m.MinorUnit = NewCT_AxisUnit()
				if err := d.DecodeElement(m.MinorUnit, &el); err != nil {
					return err
				}
			case "minorTimeUnit":
				m.MinorTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.MinorTimeUnit, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DateAx
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DateAx) Validate() error {
	return m.ValidateWithPath("CT_DateAx")
}
func (m *CT_DateAx) ValidateWithPath(path string) error {
	if err := m.AxId.ValidateWithPath(path + "/AxId"); err != nil {
		return err
	}
	if err := m.Scaling.ValidateWithPath(path + "/Scaling"); err != nil {
		return err
	}
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if err := m.AxPos.ValidateWithPath(path + "/AxPos"); err != nil {
		return err
	}
	if m.MajorGridlines != nil {
		if err := m.MajorGridlines.ValidateWithPath(path + "/MajorGridlines"); err != nil {
			return err
		}
	}
	if m.MinorGridlines != nil {
		if err := m.MinorGridlines.ValidateWithPath(path + "/MinorGridlines"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.MajorTickMark != nil {
		if err := m.MajorTickMark.ValidateWithPath(path + "/MajorTickMark"); err != nil {
			return err
		}
	}
	if m.MinorTickMark != nil {
		if err := m.MinorTickMark.ValidateWithPath(path + "/MinorTickMark"); err != nil {
			return err
		}
	}
	if m.TickLblPos != nil {
		if err := m.TickLblPos.ValidateWithPath(path + "/TickLblPos"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if err := m.CrossAx.ValidateWithPath(path + "/CrossAx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Auto != nil {
		if err := m.Auto.ValidateWithPath(path + "/Auto"); err != nil {
			return err
		}
	}
	if m.LblOffset != nil {
		if err := m.LblOffset.ValidateWithPath(path + "/LblOffset"); err != nil {
			return err
		}
	}
	if m.BaseTimeUnit != nil {
		if err := m.BaseTimeUnit.ValidateWithPath(path + "/BaseTimeUnit"); err != nil {
			return err
		}
	}
	if m.MajorUnit != nil {
		if err := m.MajorUnit.ValidateWithPath(path + "/MajorUnit"); err != nil {
			return err
		}
	}
	if m.MajorTimeUnit != nil {
		if err := m.MajorTimeUnit.ValidateWithPath(path + "/MajorTimeUnit"); err != nil {
			return err
		}
	}
	if m.MinorUnit != nil {
		if err := m.MinorUnit.ValidateWithPath(path + "/MinorUnit"); err != nil {
			return err
		}
	}
	if m.MinorTimeUnit != nil {
		if err := m.MinorTimeUnit.ValidateWithPath(path + "/MinorTimeUnit"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}