// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestCT_TblPPrConstructor(t *testing.T) {
	v := wml.NewCT_TblPPr()
	if v == nil {
		t.Errorf("wml.NewCT_TblPPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPPr should validate: %s", err)
	}
}

func TestCT_TblPPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPPr()
	xml.Unmarshal(buf, v2)
}
