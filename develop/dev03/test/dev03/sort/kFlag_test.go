package sort_test

import (
	"testing"
	"wb_lvl2/develop/dev03/internal/dev03/sort"
)

func TestKFlag_DecodeKFlag(t *testing.T) {
	kf1 := sort.NewKFlag()
	kf2 := sort.NewKFlag()
	kf3 := sort.NewKFlag()
	kf4 := sort.NewKFlag()
	kf5 := sort.NewKFlag()
	exp1 := sort.KFlag{F1: 1}
	exp2 := sort.KFlag{F1: 1, F2: 2}
	exp3 := sort.KFlag{F1: 1, F2: 2, C2: 3}
	exp4 := sort.KFlag{F1: 1, C1: 3, F2: 2}
	exp5 := sort.KFlag{F1: 1, C1: 3, F2: 2, C2: 4}

	err1 := kf1.DecodeKFlag("1")
	err2 := kf2.DecodeKFlag("1,2")
	err3 := kf3.DecodeKFlag("1,2.3")
	err4 := kf4.DecodeKFlag("1.3,2")
	err5 := kf5.DecodeKFlag("1.3,2.4")

	if err1 != nil {
		t.Fatalf("error while decoding KFlag: %s", err1)
	}
	if err2 != nil {
		t.Fatalf("error while decoding KFlag: %s", err2)
	}
	if err3 != nil {
		t.Fatalf("error while decoding KFlag: %s", err3)
	}
	if err4 != nil {
		t.Fatalf("error while decoding KFlag: %s", err4)
	}
	if err5 != nil {
		t.Fatalf("error while decoding KFlag: %s", err5)
	}
	if exp1.C1 != kf1.C1 || exp1.F1 != kf1.F1 || exp1.C2 != kf1.C2 || exp1.F2 != kf1.F2 {
		t.Fatalf("incorrect KFlag decoding, expected: %v, got: %v", exp1, kf1)
	}
	if exp2.C1 != kf2.C1 || exp2.F1 != kf2.F1 || exp2.C2 != kf2.C2 || exp2.F2 != kf2.F2 {
		t.Fatalf("incorrect KFlag decoding, expected: %v, got: %v", exp2, kf2)
	}
	if exp3.C1 != kf3.C1 || exp3.F1 != kf3.F1 || exp3.C2 != kf3.C2 || exp3.F2 != kf3.F2 {
		t.Fatalf("incorrect KFlag decoding, expected: %v, got: %v", exp3, kf3)
	}
	if exp4.C1 != kf4.C1 || exp4.F1 != kf4.F1 || exp4.C2 != kf4.C2 || exp4.F2 != kf4.F2 {
		t.Fatalf("incorrect KFlag decoding, expected: %v, got: %v", exp4, kf4)
	}
	if exp5.C1 != kf5.C1 || exp5.F1 != kf5.F1 || exp5.C2 != kf5.C2 || exp5.F2 != kf5.F2 {
		t.Fatalf("incorrect KFlag decoding, expected: %v, got: %v", exp5, kf5)
	}

}
