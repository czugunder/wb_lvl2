package sort

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"wb_lvl2/develop/dev03/internal/dev03/sort"
)

func TestSorting_SortFiles(t *testing.T) {
	testFilesDir := "testFiles/test"
	expFilesDir := "testFiles/exp"

	sorting := sort.NewSorting()
	checkData := []struct {
		testName      string
		flags         *sort.Flags
		filesSort     []string
		filesExpected []string
	}{
		{
			testName: "lowercase letters",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
			},
			filesSort:     []string{"test1.txt"},
			filesExpected: []string{"test1.txt"},
		},
		{
			testName: "reverse (flag -r)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				R:     true,
			},
			filesSort:     []string{"test1.txt"},
			filesExpected: []string{"test1r.txt"},
		},
		{
			testName: "numbers (flag -n)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				N:     true,
			},
			filesSort:     []string{"test2.txt"},
			filesExpected: []string{"test2.txt"},
		},
		{
			testName: "unique (flag -u)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				U:     true,
			},
			filesSort:     []string{"test3.txt"},
			filesExpected: []string{"test3.txt"},
		},
		{
			testName: "month name (flag -M)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				M:     true,
			},
			filesSort:     []string{"test4.txt"},
			filesExpected: []string{"test4.txt"},
		},
		{
			testName: "ignore tailing blanks (-b)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				B:     true,
			},
			filesSort:     []string{"test5.txt"},
			filesExpected: []string{"test5.txt"},
		},
		{
			testName: "suffix sort (-h)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				H:     true,
			},
			filesSort:     []string{"test6.txt"},
			filesExpected: []string{"test6.txt"},
		},
		{
			testName: "month column sort (-k, -M)",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				K:     "1,1",
				M:     true,
			},
			filesSort:     []string{"testK.txt"},
			filesExpected: []string{"testK1.txt"},
		},
		{
			testName: "part column sort",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				K:     "3,3.2",
			},
			filesSort:     []string{"testK.txt"},
			filesExpected: []string{"testK2.txt"},
		},
		{
			testName: "part multi-column sort",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				K:     "2.1,3.2",
			},
			filesSort:     []string{"testK.txt"},
			filesExpected: []string{"testK3.txt"},
		},
		{
			testName: "number column sort",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				K:     "3.3,3.4",
				N:     true,
			},
			filesSort:     []string{"testK.txt"},
			filesExpected: []string{"testK4.txt"},
		},
		{
			testName: "suffix column sort",
			flags: &sort.Flags{
				KFlag: &sort.KFlag{},
				K:     "4",
				H:     true,
			},
			filesSort:     []string{"testK.txt"},
			filesExpected: []string{"testK5.txt"},
		},
	}

	for _, v := range checkData {
		testDir := os.TempDir()
		files := make([]string, len(v.filesSort), len(v.filesSort))
		for i, fileName := range v.filesSort {
			if err := copyFile(t, testDir, testFilesDir, fileName); err != nil {
				t.Fatalf("test %s: cant copy test file: %s from: %s to test directory: %s",
					v.testName, fileName, testFilesDir, testDir)
			}
			files[i] = testDir + "/" + fileName
		}
		if err := sorting.SetFlags(v.flags); err != nil {
			t.Fatalf("test %s: error while setting flags: %s", v.testName, err)
		}
		sorting.SetFiles(files)
		if err := sorting.SortFiles(); err != nil {
			t.Fatalf("test %s: error while sorting: %s", v.testName, err)
		}
		for i := range v.filesExpected {
			expFileName := expFilesDir + "/" + v.filesExpected[i]
			if equal, err := compareFiles(t, expFileName, files[i]); err != nil {
				t.Fatalf("test %s: error while comparing file: %s with file: %s",
					v.testName, expFileName, files[i])
			} else if !equal {
				t.Fatalf("test %s: file: %s is not equal to file: %s", v.testName, expFileName, files[i])
			}
		}
	}
}

func copyFile(t *testing.T, testDir, curPath, fileName string) error {
	t.Helper()
	src := curPath + "/" + fileName
	dst := testDir + "/" + fileName
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func compareFiles(t *testing.T, expFilePath, testFilePath string) (bool, error) {
	t.Helper()
	f := &sort.Flags{}
	tf1 := sort.NewTextFile(expFilePath, f)
	tf2 := sort.NewTextFile(testFilePath, f)
	if err := tf1.Read(); err != nil {
		return false, err
	}
	if err := tf2.Read(); err != nil {
		return false, err
	}
	if len(tf1.GetStrings()) != len(tf2.GetStrings()) {
		return false, nil
	}
	for i := range tf1.GetStrings() {
		if tf1.GetStrings()[i] != tf2.GetStrings()[i] {
			return false, nil
		}
	}
	return true, nil
}

func TestSorting_FlagC(t *testing.T) {
	testFilesDir := "testFiles/test"
	buff := bytes.Buffer{}
	sorting := sort.NewSorting()
	checkData := []struct {
		testName  string
		flags     *sort.Flags
		filesSort []string
		output    []string
	}{
		{
			testName: "flag -c",
			flags: &sort.Flags{
				KFlag:  &sort.KFlag{},
				C:      true,
				Writer: &buff,
			},
			filesSort: []string{"testCsorted.txt", "testCunsorted.txt"},
			output:    []string{"Input is sorted", "Input is NOT sorted", ""},
		},
	}

	for _, v := range checkData {
		buff.Reset()
		testDir := os.TempDir()
		files := make([]string, len(v.filesSort), len(v.filesSort))
		for i, fileName := range v.filesSort {
			if err := copyFile(t, testDir, testFilesDir, fileName); err != nil {
				t.Fatalf("test %s: cant copy test file: %s from: %s to test directory: %s",
					v.testName, fileName, testFilesDir, testDir)
			}
			files[i] = testDir + "/" + fileName
		}
		if err := sorting.SetFlags(v.flags); err != nil {
			t.Fatalf("test %s: error while setting flags: %s", v.testName, err)
		}
		sorting.SetFiles(files)
		if err := sorting.SortFiles(); err != nil {
			t.Fatalf("test %s: error while sorting: %s", v.testName, err)
		}
		outputRaw := buff.String()
		output := strings.Split(outputRaw, "\n")
		if len(v.output) != len(output) {
			t.Fatalf("test %s: incorrect length of output, expected: %d, got: %d",
				v.testName, len(v.output), len(output))
		}
		for i := range v.output {
			if v.output[i] != output[i] {
				t.Fatalf("test %s: incorrect output, expected: %s, got: %s", v.testName, v.output[i], output[i])
			}
		}
	}
}
