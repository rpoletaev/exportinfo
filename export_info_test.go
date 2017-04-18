package exportinfo

import (
	"fmt"
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
		Tst{
			Input: `<ns2:contractProcedure xsi:type="ns3:zfcs_contractProcedureType" schemeVersion="4.5" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "contractProcedure",
				Version: "4.5",
			},
		},
	}

	for i, test := range tests {
		got, err := GetExportInfo(test.Input)
		fmt.Printf("%d\n", i)
		if err != nil {
			if test.IsError {
				return
			}

			t.Error(err)
		}

		if *got != test.Expected {
			t.Errorf("got %v != Expected %v", *got, test.Expected)
			t.Fail()
		}
	}
}
