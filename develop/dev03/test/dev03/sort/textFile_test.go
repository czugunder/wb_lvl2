package sort_test

import (
	"os"
	"testing"
	"wb_lvl2/develop/dev03/internal/dev03/sort"
)

func TestTextFile_Read(t *testing.T) {
	fl := sort.NewFlags()
	tf := sort.NewTextFile("testFiles/testRead.txt", fl)
	exp := []string{"test", "Read", "test Write"}

	err := tf.Read()

	if err != nil {
		t.Fatalf("error while reading textfile: %s", err)
	}
	testData := tf.GetStrings()
	if len(exp) != len(testData) {
		t.Fatalf("incorrect quantitiy of strings in textfile, expected: %d, got: %d", len(exp), len(testData))
	}
	for i := range exp {
		if exp[i] != testData[i] {
			t.Fatalf("incorrect string was read, expected: %s, got: %s", exp[i], testData[i])
		}
	}
}

func TestTextFile_Write(t *testing.T) {
	fl := sort.NewFlags()
	ss := sort.NewSortString(fl)
	tempDir := os.TempDir()
	tfW := sort.NewTextFile(tempDir+"/testWrite.txt", fl)
	tfW.SetStrings(ss)
	exp := []string{"test", "Read", "test Write"}
	ss.SetStrings(exp)
	tfR := sort.NewTextFile(tempDir+"/testWrite.txt", fl)

	err := tfW.Write()
	_ = tfR.Read()

	if err != nil {
		t.Fatalf("error while reading textfile: %s", err)
	}
	testData := tfR.GetStrings()
	if len(exp) != len(testData) {
		t.Fatalf("incorrect quantitiy of strings in textfile, expected: %d, got: %d", len(exp), len(testData))
	}
	for i := range exp {
		if exp[i] != testData[i] {
			t.Fatalf("incorrect string was read, expected: %s, got: %s", exp[i], testData[i])
		}
	}

}
