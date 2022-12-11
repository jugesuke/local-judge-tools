package testcases

import (
	"embed"
	"path/filepath"
)

type Testcase struct {
	name string
	in   string
	out  string
}

//go:embed data/*/*.in
var testcasesIn embed.FS

//go:embed data/*/*.out
var testcasesOut embed.FS

func NewTestCase(name string, in string, out string) *Testcase {
	return &Testcase{name: name, in: in, out: out}
}

func GetTestcase(questionName string, testcaseName string) (*Testcase, error) {
	in, err := testcasesIn.ReadFile("data/" + questionName + "/" + testcaseName + ".in")
	if err != nil {
		return nil, err
	}

	out, err := testcasesOut.ReadFile("data/" + questionName + "/" + testcaseName + ".out")
	if err != nil {
		return nil, err
	}
	return NewTestCase(testcaseName, string(in), string(out)), nil
}

func GetTestcases(questionName string) (*[]*Testcase, error) {
	dirs, err := testcasesIn.ReadDir("data/" + questionName)
	if err != nil {
		return nil, err
	}
	var testcases []*Testcase

	for _, file := range dirs {
		if file.IsDir() {
			continue
		}

		testcaseName := getFilepathWithoutExt(file.Name())

		t, err := GetTestcase(questionName, testcaseName)
		if err != nil {
			return nil, err
		}

		testcases = append(testcases, t)
	}
	return &testcases, nil
}

func (t *Testcase) GetName() string {
	return t.name
}

func (t *Testcase) GetStdin() string {
	return t.in
}

func (t *Testcase) Is(stdout string) bool {
	return t.out == stdout
}

func getFilepathWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
