package grep

import (
	"testing"
	"wb_lvl2/develop/dev05/internal/grep"
)

func Test_readFile(t *testing.T) {
	data := []string{"aaaa", "bbb", "cccccc", "dddddd", "test", "aaaaaaa", "bbbbbb", "ccccccc", "dddddd"}

	testData, err := grep.ReadFile("textFiles/test1.txt")

	if err != nil {
		t.Fatal(err)
	}
	if len(data) != len(testData) {
		t.Fatalf("incorrect length of read slice, expected: %d, got %d\n", len(data), len(testData))
	}
	for i := range data {
		if data[i] != testData[i] {
			t.Fatalf("incorrect read data, expected: %s, len=%d, got: %s, len=%d\n",
				data[i], len(data[i]), testData[i], len(testData[i]))
		}
	}
}
