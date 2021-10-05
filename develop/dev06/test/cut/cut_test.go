package cut_test

import (
	"bytes"
	"strings"
	"testing"
	"wb_lvl2/develop/dev06/internal/config"
	"wb_lvl2/develop/dev06/internal/cut"
)

func TestCut_MakePattern(t *testing.T) {
	cfg := config.NewConfig()
	r1 := config.NewSingleRange(5)
	r2 := config.NewLongRange(7, 9)
	r3 := config.NewFromStartRange(3)
	r4 := config.NewToEndRange(11)
	cfg.Ranges = append(cfg.Ranges, r1, r2, r3, r4)
	c := cut.NewCut()
	c.SetConfig(cfg)
	expected := []bool{true, true, true, false, true, false, true, true, true, false, true, true}

	testData := c.MakePattern(len(expected))

	if len(expected) != len(testData) {
		t.Fatalf("incorrect length of output slice, expected %d, got: %d", len(expected), len(testData))
	}
	for i := range expected {
		if expected[i] != testData[i] {
			t.Fatalf("incorrect output, expected %t, got: %t", expected[i], testData[i])
		}
	}

}

func TestCut_FormatString(t *testing.T) {
	cfg := config.NewConfig()
	r1 := config.NewSingleRange(5)
	r2 := config.NewLongRange(7, 9)
	r3 := config.NewFromStartRange(3)
	r4 := config.NewToEndRange(11)
	cfg.Ranges = append(cfg.Ranges, r1, r2, r3, r4)
	c := cut.NewCut()
	c.SetConfig(cfg)
	str1 := "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\t11\t12\t13"
	str2 := "1WWW2WWW3WWW4WWW5WWW6WWW7WWW8WWW9WWW10WWW11WWW12WWW13"
	str3 := "1WWW2WWW3WWW4WWW5WWW6WWW7"
	str4 := "1MMM2MMM3"
	expected1 := "1\t2\t3\t5\t7\t8\t9\t11\t12\t13"
	expected2 := "1WWW2WWW3WWW5WWW7WWW8WWW9WWW11WWW12WWW13"
	expected3 := "1WWW2WWW3WWW5WWW7"
	expected4 := "1MMM2MMM3"
	expectedBool1 := true
	expectedBool2 := true
	expectedBool3 := true
	expectedBool4 := false

	cfg.D = "\t"
	testData1, testBool1 := c.FormatString(str1)
	cfg.D = "WWW"
	testData2, testBool2 := c.FormatString(str2)
	testData3, testBool3 := c.FormatString(str3)
	testData4, testBool4 := c.FormatString(str4)

	if expectedBool1 != testBool1 {
		t.Fatalf("incorrect delimeter detection, expected: %t, got: %t", expectedBool1, testBool1)
	}
	if expected1 != testData1 {
		t.Fatalf("incorrect output string, expected: %s, got: %s", expected1, testData1)
	}
	if expectedBool2 != testBool2 {
		t.Fatalf("incorrect delimeter detection, expected: %t, got: %t", expectedBool2, testBool2)
	}
	if expected2 != testData2 {
		t.Fatalf("incorrect output string, expected: %s, got: %s", expected2, testData2)
	}
	if expectedBool3 != testBool3 {
		t.Fatalf("incorrect delimeter detection, expected: %t, got: %t", expectedBool3, testBool3)
	}
	if expected3 != testData3 {
		t.Fatalf("incorrect output string, expected: %s, got: %s", expected3, testData3)
	}
	if expectedBool4 != testBool4 {
		t.Fatalf("incorrect delimeter detection, expected: %t, got: %t", expectedBool4, testBool4)
	}
	if expected4 != testData4 {
		t.Fatalf("incorrect output string, expected: %s, got: %s", expected4, testData4)
	}
}

func TestCut_Run(t *testing.T) {
	cfg := config.NewConfig()
	r1 := config.NewSingleRange(5)
	r2 := config.NewLongRange(7, 9)
	r3 := config.NewFromStartRange(3)
	r4 := config.NewToEndRange(11)
	cfg.Ranges = append(cfg.Ranges, r1, r2, r3, r4)
	cfg.D = "WWW"
	cfg.S = true // флаг -s поднят
	c := cut.NewCut()
	c.SetConfig(cfg)
	wBuff, rBuff := bytes.Buffer{}, bytes.Buffer{}
	c.SetWriter(&wBuff)
	c.SetReader(&rBuff)
	str := []string{"1WWW2WWW3WWW4WWW5WWW6WWW7WWW8WWW9WWW10WWW11WWW12WWW13", "1MMM2MMM3"}
	expected := []string{"1WWW2WWW3WWW5WWW7WWW8WWW9WWW11WWW12WWW13", ""}
	for _, v := range str {
		rBuff.WriteString(v + "\n")
	}

	err := c.Run()
	testData := strings.Split(wBuff.String(), "\n")

	if err != nil {
		t.Fatalf("incorrect work: %s", err)
	}
	if len(expected) != len(testData) {
		t.Fatalf("incorrect length of output, expected: %d, got: %d", len(expected), len(testData))
	}
	for i := range expected {
		if expected[i] != testData[i] {
			t.Fatalf("incorrect output string, expected: %s, got: %s", expected[i], testData[i])
		}
	}
}
