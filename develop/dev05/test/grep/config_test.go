package grep

import (
	"testing"
	"wb_lvl2/develop/dev05/internal/grep"
)

func Test_getRest(t *testing.T) {
	var a1 []string
	var a2 = []string{"request"}
	var a3 = []string{"request", "test.txt"}
	var a4 = []string{"request", "test.txt", "test2.txt"}
	config1 := grep.NewConfig()
	config2 := grep.NewConfig()
	config3 := grep.NewConfig()
	config4 := grep.NewConfig()

	err1 := config1.GetRest(a1)
	err2 := config2.GetRest(a2)
	err3 := config3.GetRest(a3)
	err4 := config4.GetRest(a4)

	if err1 == nil {
		t.Fatalf("Error is expected when no pattern is given, but it hasn't happened: %s\n", err1)
	}
	if err2 != nil {
		t.Fatalf("Error isn't expected when pattern is given, but it has happened: %s\n", err2)
	} else {
		if config2.Pattern != a2[0] {
			t.Fatalf("Wrong pattern was written to config, expected: %s, got: %s\n", a2[0], config2.Pattern)
		}
	}
	if err3 != nil {
		t.Fatalf("Error isn't expected when pattern is given, but it has happened: %s\n", err3)
	} else {
		if config3.Pattern != a3[0] {
			t.Fatalf("Wrong pattern was written to config, expected: %s, got: %s\n", a3[0], config3.Pattern)
		}
		if len(config3.Files) != len(a3)-1 {
			t.Fatalf("Quanttity of input files written to config is incorrect, expected: %d, got: %d\n",
				len(a3)-1, len(config3.Files))
		} else {
			if config3.Files[0] != a3[1] {
				t.Fatalf("Path of input file written to config is incorrect, expected: %s, got: %s\n",
					a3[1], config3.Files[0])
			}
		}
	}
	if err4 != nil {
		t.Fatalf("Error isn't expected when pattern is given, but it has happened: %s\n", err4)
	} else {
		if config4.Pattern != a4[0] {
			t.Fatalf("Wrong pattern was written to config, expected: %s, got: %s\n", a4[0], config4.Pattern)
		}
		if len(config4.Files) != len(a4)-1 {
			t.Fatalf("Quanttity of input files written to config is incorrect, expected: %d, got: %d\n",
				len(a4)-1, len(config4.Files))
		} else {
			if config4.Files[0] != a4[1] {
				t.Fatalf("Path of input file written to config is incorrect, expected: %s, got: %s\n",
					a4[1], config4.Files[0])
			}
			if config4.Files[1] != a4[2] {
				t.Fatalf("Path of input file written to config is incorrect, expected: %s, got: %s\n",
					a4[2], config4.Files[1])
			}
		}
	}
}
