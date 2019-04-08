package configs

import (
	"strings"
	"testing"
)

func TestSingleAction(t *testing.T) {
	type testcase struct {
		Input     string
		Source    string
		Dest      string
		Filetypes string
	}

	tst := testcase{
		Input: `<config>
					<action source="C:\workspace\V81\AscensionSC8" destination="C:\SC82\Website_Data\Website\bin" filetypes="dll,pdb" />
				</config>`,
		Source:    `C:\workspace\V81\AscensionSC8`,
		Dest:      `C:\SC82\Website_Data\Website\bin`,
		Filetypes: `dll,pdb`,
	}

	cnfg, err := GetConfigs(strings.NewReader(tst.Input))
	if err != nil {
		t.Errorf("GetConfigs() is unable to parse input: %v", err)
	}

	act := cnfg.Actions[0]
	if act.Source != tst.Source || act.Destination != tst.Dest || act.FileTypes != tst.Filetypes {
		t.Error("Unexpected result from GetConfigs() XML parse")
	}
}

func TestMultipleAction(t *testing.T) {
	type testcase struct {
		Input string
		Count int
	}

	tst := testcase{
		Input: `<config>
					<action source="C:\workspace\V81\AscensionSC8" destination="C:\SC82\Website_Data\Website\bin" filetypes="dll,pdb" />
					<action source="C:\workspace\V81\AscensionSC8" destination="C:\SC82\Website_Data\Website\bin" filetypes="dll,pdb" />
					<action source="C:\workspace\V81\AscensionSC8" destination="C:\SC82\Website_Data\Website\bin" filetypes="dll,pdb" />
				</config>`,
		Count: 3,
	}

	cnfg, err := GetConfigs(strings.NewReader(tst.Input))
	if err != nil {
		t.Errorf("GetConfigs() is unable to parse input: %v", err)
	}

	if len(cnfg.Actions) != tst.Count {
		t.Error("GetConfigs() came back with an unexpected count")
	}
}

func TestIsValidFileType(t *testing.T) {
	type testcase struct {
		TestAction Action
		TestSrc    string
		result     bool
	}

	cases := []testcase{
		testcase{TestAction: Action{FileTypes: "dll"}, TestSrc: `C:\test\test.dll`, result: true},
		testcase{TestAction: Action{FileTypes: "dll"}, TestSrc: `C:\test\test.fake`, result: false},
		testcase{TestAction: Action{FileTypes: "dll,fake"}, TestSrc: `C:\test\test.dll`, result: true},
		testcase{TestAction: Action{FileTypes: "dll,fake"}, TestSrc: `C:\test\test.bat`, result: false},
	}

	for _, tcase := range cases {
		if tcase.TestAction.IsValidFileType(tcase.TestSrc) != tcase.result {
			t.Fatalf("Unexpected result from IsValidFileType() during this test case: %+v", tcase)
		}
	}
}

func TestGetActionBySource(t *testing.T) {
	type testcase struct {
		Input string
	}

	tst := testcase{
		Input: `<config>
					<action source="C:\workspace\V81\AscensionSC8" destination="C:\SC82\Website_Data\Website\bin" filetypes="dll,pdb" />
					<action source="C:\workspace\V81\success" destination="success" filetypes="dll,pdb" />
					<action source="C:\workspace\V81\AscensionSC8222" destination="C:\SC82\Website_Data\Website\bin222" filetypes="dll,pdb" />
				</config>`,
	}

	cnfg, err := GetConfigs(strings.NewReader(tst.Input))
	if err != nil {
		t.Fatalf("GetConfigs() is unable to parse input: %v", err)
	}

	act := cnfg.GetActionBySource(`C:\workspace\V81\success\file.bat`)
	if act.Destination != "success" {
		t.Fatal("GetActionBySource() returned an unexpected Action")
	}
}
