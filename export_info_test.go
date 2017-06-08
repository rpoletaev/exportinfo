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
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><ns2:control99ProtocolMismatch schemeVersion="7.0">`,
			Expected: ExportInfo{
				Title:   "control99ProtocolMismatch",
				Version: "7.0",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><ns2:sketchPlan schemeVersion="1.0" xmlns="http://zakupki.gov.ru/oos/types/1" xmlns:ns2="http://zakupki.gov.ru/oos/printform/1">`,
			Expected: ExportInfo{
				Title:   "sketchPlan",
				Version: "1.0",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><ns2:fcsPurchaseDocument schemeVersion="4.2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "fcsPurchaseDocument",
				Version: "4.2",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><fcsPurchaseDocument schemeVersion="4.2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "fcsPurchaseDocument",
				Version: "4.2",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><ns2:contractProcedure xsi:type="ns3:zfcs_contractProcedureType" schemeVersion="4.5" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">`,
			Expected: ExportInfo{
				Title:   "contractProcedure",
				Version: "4.5",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8"?><export xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://zakupki.gov.ru/oos/export/1" xmlns:oos="http://zakupki.gov.ru/oos/types/1"><tenderPlan schemeVersion="1.0">`,
			Expected: ExportInfo{
				Title:   "tenderPlan",
				Version: "1.0",
			},
		},
		Tst{
			Input: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns2:sketchPlan schemeVersion="1.0" xmlns="http://zakupki.gov.ru/oos/types/1" xmlns:ns2="http://zakupki.gov.ru/oos/printform/1"> <commonInfo><id>225</id><planNumber>201403082000041</planNumber>`,
			Expected: ExportInfo{
				Title:   "sketchPlan",
				Version: "1.0",
			},
		},
		Tst{
<<<<<<< HEAD
			Input: `<contract schemeVersion="4.5">`,
			Expected: ExportInfo{
				Title:   "contract",
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

func TestGetExportInfo2(t *testing.T) {
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
		Tst{
			Input: `<tenderPlan schemeVersion="1.0">`,
			Expected: ExportInfo{
				Title:   "tenderPlan",
				Version: "1.0",
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
			Input: `<contract schemeVersion="4.5">`,
			Expected: ExportInfo{
				Title:   "contract",
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
