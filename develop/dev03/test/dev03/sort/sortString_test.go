package sort_test

import (
	"testing"
	"wb_lvl2/develop/dev03/internal/dev03/sort"
)

func TestSortString_RemoveDuplicates(t *testing.T) {
	fl := sort.NewFlags()
	ss := sort.NewSortString(fl)
	strs := []string{"aa", "a", "bb", "aa", "b", "bb"}
	ss.SetStrings(strs)
	exp := []string{"aa", "a", "bb", "b"}

	ss.RemoveDuplicates()

	testData := ss.GetStrings()
	if len(exp) != len(testData) {
		t.Fatalf("incorrect quantitiy of strings after removing duplicates, expected: %d, got: %d",
			len(exp), len(testData))
	}
	for i := range exp {
		if exp[i] != testData[i] {
			t.Fatalf("incorrect string was found after removing duplicates, expected: %s, got: %s",
				exp[i], testData[i])
		}
	}

}

func Test_CheckSuffix(t *testing.T) {
	sufs := map[string]int{"n": -9, "mi": -6, "m": -3, "K": 3, "M": 6, "G": 9}
	checkData := []string{"2", "2n", "2mi", "2G", "2m", "2M", "2K"}
	exp := []string{"", "n", "mi", "G", "m", "M", "K"}
	testData := make([]string, len(checkData), len(checkData))

	for i := range checkData {
		testData[i] = sort.CheckSuffix(sufs, checkData[i])
	}

	for i := range exp {
		if exp[i] != testData[i] {
			t.Fatalf("incorrect suffix was found, expected: %s, got: %s", exp[i], testData[i])
		}
	}
}

func Test_ProcessKFlag(t *testing.T) {
	checkData := []struct {
		input string
		f     *sort.Flags
		exp   string
	}{
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 3},
			},
			exp: "cccccddddeee",
		},
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 2, F2: 4},
			},
			exp: "bbbbcccccdddd",
		},
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 1, F2: 3, C2: 3},
			},
			exp: "aaabbbbccc",
		},
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 1, C1: 3, F2: 2},
			},
			exp: "abbbb",
		},
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 3, C1: 2, F2: 5, C2: 1},
			},
			exp: "ccccdddde",
		},
		{
			input: "aaa bbbb ccccc dddd eee",
			f: &sort.Flags{
				KFlag: &sort.KFlag{F1: 3, C1: 2, F2: 3, C2: 4},
			},
			exp: "ccc",
		},
	}

	for _, v := range checkData {
		testData := sort.ProcessKFlag(v.input, v.f)
		if v.exp != testData {
			t.Fatalf("incorrect output of ProccessKFlag, expected: %s, got: %s", v.exp, testData)
		}
	}
}
