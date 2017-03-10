package exportinfo

import (
	"testing"
)

type Tst struct {
	Input    string
	Expected ExportInfo
	IsError  bool
}

func TestGetExportInfo(t *testing.T) {
	tests := []Tst{
		Tst{
			Input: `<ns2:control99ProtocolMismatch schemeVersion="7.0">`,
			Expected: ExportInfo{
				Title:   "control99ProtocolMismatch",
				Version: "7.0",
			},
		},
		Tst{
			Input: `<ns2:sketchPlan schemeVersion="1.0" xmlns="http://zakupki.gov.ru/oos/types/1" xmlns:ns2="http://zakupki.gov.ru/oos/printform/1">`,
			Expected: ExportInfo{
				Title:   "sketchPlan",
				Version: "1.0",
			},
		},
		Tst{
			Input: `<ns2:fcsPurchaseDocument schemeVersion="4.2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "fcsPurchaseDocument",
				Version: "4.2",
			},
		},
		Tst{
			Input: `<fcsPurchaseDocument schemeVersion="4.2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "fcsPurchaseDocument",
				Version: "4.2",
			},
		},
	}

	for _, test := range tests {
		got, err := getExportInfo(test.Input)
		if err != nil && test.IsError {
			return
		}

		if *got != test.Expected {
			t.Fail()
		}
	}
}
