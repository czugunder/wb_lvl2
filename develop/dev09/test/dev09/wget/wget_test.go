package wget_test

import (
	"io/ioutil"
	"testing"
	"wb_lvl2/develop/dev09/internal/dev09/wget"
	"wb_lvl2/develop/dev09/test/dev09/wget/testSite"
)

func TestWget_GetSavePathAndName(t *testing.T) {
	w := wget.NewWget()
	w.SetDomain("https://test.one")
	url1 := "https://test.one"
	url2 := "https://test.one/two"
	url3 := "https://test.one/zone/three"
	exp1P, exp1N := "", "index"
	exp2P, exp2N := "", "two"
	exp3P, exp3N := "/zone", "three"

	r1P, r1N := w.GetSavePathAndName(url1)
	r2P, r2N := w.GetSavePathAndName(url2)
	r3P, r3N := w.GetSavePathAndName(url3)

	if exp1P != r1P {
		t.Fatalf("incorrect path, expected: %s, got: %s", exp1P, r1P)
	}
	if exp1N != r1N {
		t.Fatalf("incorrect name, expected: %s, got: %s", exp1N, r1N)
	}
	if exp2P != r2P {
		t.Fatalf("incorrect path, expected: %s, got: %s", exp2P, r2P)
	}
	if exp2N != r2N {
		t.Fatalf("incorrect name, expected: %s, got: %s", exp2N, r2N)
	}
	if exp3P != r3P {
		t.Fatalf("incorrect path, expected: %s, got: %s", exp3P, r3P)
	}
	if exp3N != r3N {
		t.Fatalf("incorrect name, expected: %s, got: %s", exp3N, r3N)
	}
}

func TestWget_GetSavePathAndName2(t *testing.T) {
	w := wget.NewWget()
	w.SetDomain("https://demo2.chethemes.com/shoesmarket")
	url1 := "https://demo2.chethemes.com/shoesmarket/comments/feed/"
	exp1P, exp1N := "/comments", "feed"

	r1P, r1N := w.GetSavePathAndName(url1)

	if exp1P != r1P {
		t.Fatalf("incorrect path, expected: %s, got: %s", exp1P, r1P)
	}
	if exp1N != r1N {
		t.Fatalf("incorrect name, expected: %s, got: %s", exp1N, r1N)
	}

}

func TestWget_ParseLinks(t *testing.T) {
	w := wget.NewWget()
	data := `one href="one", two three href="two"four`
	exp := []string{
		"href=\"one\"",
		"href=\"two\"",
	}

	testData := w.ParseLinks(data)

	if len(exp) != len(testData) {
		t.Fatalf("length of output isn't correct, expected: %d, got: %d", len(exp), len(testData))
	}
	for i := range exp {
		if exp[i] != testData[i] {
			t.Fatalf("incorrect output, expected: %s, got: %s", exp[i], testData[i])
		}
	}
}

func TestWget_GetSite(t *testing.T) {
	saveDir := t.TempDir()
	w := wget.NewWget()
	exp := map[string]bool{"index.html": false, "one": false, "one.html": false, "one.php": false}
	expInternal := "two.html"
	go testSite.Server()

	w.SetSaveDirectory(saveDir)
	w.SetDomain("http://localhost:8888")
	w.SaveSite()

	files, _ := ioutil.ReadDir(saveDir)
	if len(exp) != len(files) {
		t.Fatalf("incorrect quantity of downloaded pages, expected: %d, got: %d", len(exp), len(files))
	}
	for _, f := range files {
		if _, found := exp[f.Name()]; !found {
			t.Fatalf("unexpected page name: %s", f.Name())
		}
		if f.IsDir() {
			if f.Name() != "one" {
				t.Fatalf("incorrect name of internal directory, expected: one, got: %s", f.Name())
			}
			internalFiles, _ := ioutil.ReadDir(saveDir + "/one")
			if 1 != len(internalFiles) {
				t.Fatalf("incorrect quantity of downloaded internal pages, expected: 1, got: %d", len(internalFiles))
			}
			if expInternal != internalFiles[0].Name() {
				t.Fatalf("unexpected internal page name: %s", internalFiles[0].Name())
			}
		}
	}
}
