package config_test

import (
	"testing"
	"wb_lvl2/develop/dev06/internal/config"
)

func TestDecodeRange(t *testing.T) {
	rr1 := "1-10"               // валидно
	rr2 := "5"                  // валидно
	rr3 := "6-"                 // валидно
	rr4 := "-5"                 // валидно
	rr5 := "-"                  // не валидно
	rr6 := "4a5"                // не валидно
	rr7 := "1-2-5"              // не валидно
	rr8 := "4a5-1b7"            // не валидно
	rr9 := "4a5-"               // не валидно
	rr10 := "-4a5"              // не валидно
	rr11 := "34-4a5"            // не валидно
	rr12 := "4a5-34"            // не валидно
	rr13 := "    1   -   10   " // валидно
	rr14 := "   5    "          // валидно
	rr15 := "    6    -   "     // валидно
	rr16 := "    -   5    "     // валидно

	r1, err1 := config.DecodeRange(rr1)
	r2, err2 := config.DecodeRange(rr2)
	r3, err3 := config.DecodeRange(rr3)
	r4, err4 := config.DecodeRange(rr4)
	r5, err5 := config.DecodeRange(rr5)
	r6, err6 := config.DecodeRange(rr6)
	r7, err7 := config.DecodeRange(rr7)
	r8, err8 := config.DecodeRange(rr8)
	r9, err9 := config.DecodeRange(rr9)
	r10, err10 := config.DecodeRange(rr10)
	r11, err11 := config.DecodeRange(rr11)
	r12, err12 := config.DecodeRange(rr12)
	r13, err13 := config.DecodeRange(rr13)
	r14, err14 := config.DecodeRange(rr14)
	r15, err15 := config.DecodeRange(rr15)
	r16, err16 := config.DecodeRange(rr16)

	if err1 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r1.GetStart() != 1 || r1.GetEnd() != 10 || r1.GetToEnd() != false || r1.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r1")
		}
	}

	if err2 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r2.GetStart() != 5 || r2.GetEnd() != 5 || r2.GetToEnd() != false || r2.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r2")
		}
	}

	if err3 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r3.GetStart() != 6 || r3.GetEnd() != 0 || r3.GetToEnd() != true || r3.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r3")
		}
	}

	if err4 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r4.GetStart() != 0 || r4.GetEnd() != 5 || r4.GetToEnd() != false || r4.GetFromStart() != true {
			t.Fatalf("incorrect output of config.decodeRange(string) on r4")
		}
	}

	if err5 == nil || r5 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err6 == nil || r6 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err7 == nil || r7 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err8 == nil || r8 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err9 == nil || r9 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err10 == nil || r10 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err11 == nil || r11 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err12 == nil || r12 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	}

	if err13 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r13.GetStart() != 1 || r13.GetEnd() != 10 || r13.GetToEnd() != false || r13.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r13")
		}
	}

	if err14 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r14.GetStart() != 5 || r14.GetEnd() != 5 || r14.GetToEnd() != false || r14.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r14")
		}
	}

	if err15 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r15.GetStart() != 6 || r15.GetEnd() != 0 || r15.GetToEnd() != true || r15.GetFromStart() != false {
			t.Fatalf("incorrect output of config.decodeRange(string) on r15")
		}
	}

	if err16 != nil {
		t.Fatalf("incorrect work of config.decodeRange(string)")
	} else {
		if r16.GetStart() != 0 || r16.GetEnd() != 5 || r16.GetToEnd() != false || r16.GetFromStart() != true {
			t.Fatalf("incorrect output of config.decodeRange(string) on r16")
		}
	}
}

func TestCheckValidRange(t *testing.T) {
	fr1 := config.NewSingleRange(5)     // валидно
	fr2 := config.NewLongRange(1, 10)   // валидно
	fr3 := config.NewToEndRange(5)      // валидно
	fr4 := config.NewFromStartRange(5)  // валидно
	fr5 := config.NewSingleRange(0)     // не валидно
	fr6 := config.NewLongRange(0, 0)    // не валидно
	fr7 := config.NewLongRange(0, 1)    // не валидно
	fr8 := config.NewLongRange(1, 0)    // не валидно
	fr9 := config.NewLongRange(10, 1)   // не валидно
	fr10 := config.NewToEndRange(0)     // не валидно
	fr11 := config.NewFromStartRange(0) // не валидно

	r1 := config.CheckValidRange(fr1)
	r2 := config.CheckValidRange(fr2)
	r3 := config.CheckValidRange(fr3)
	r4 := config.CheckValidRange(fr4)
	r5 := config.CheckValidRange(fr5)
	r6 := config.CheckValidRange(fr6)
	r7 := config.CheckValidRange(fr7)
	r8 := config.CheckValidRange(fr8)
	r9 := config.CheckValidRange(fr9)
	r10 := config.CheckValidRange(fr10)
	r11 := config.CheckValidRange(fr11)

	if r1 != true || r2 != true || r3 != true || r4 != true {
		t.Fatal("incorrect output of CheckValidRange, expected: true, got: false")
	}
	if r5 != false || r6 != false || r7 != false || r8 != false || r9 != false || r10 != false || r11 != false {
		t.Fatal("incorrect output of CheckValidRange, expected: false, got: true")
	}
}

func TestConfig_DecodeFlagF(t *testing.T) {
	c := config.NewConfig()
	c.F = "1, 1-10, 1-, -10,  1  ,   1  -   10   , 1     -   ,  -   10"
	var expected = []*config.FRange{
		config.NewSingleRange(1),
		config.NewLongRange(1, 10),
		config.NewToEndRange(1),
		config.NewFromStartRange(10),
		config.NewSingleRange(1),
		config.NewLongRange(1, 10),
		config.NewToEndRange(1),
		config.NewFromStartRange(10),
	}

	err := c.DecodeFlagF()
	if err != nil {
		t.Fatal("incorrect execution of DecodeFlagF")
	}
	if len(expected) != len(c.Ranges) {
		t.Fatalf("incorrect length of output, expected: %d, got: %d", len(expected), len(c.Ranges))
	}
	for i := range c.Ranges {
		if expected[i].GetStart() != c.Ranges[i].GetStart() ||
			expected[i].GetEnd() != c.Ranges[i].GetEnd() ||
			expected[i].GetFromStart() != c.Ranges[i].GetFromStart() ||
			expected[i].GetToEnd() != c.Ranges[i].GetToEnd() {
			t.Fatalf("incorrect output, expected: %v, got: %v", expected[i], c.Ranges[i])
		}
	}
}
