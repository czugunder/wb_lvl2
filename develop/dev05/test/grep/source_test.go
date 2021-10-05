package grep

import (
	"bytes"
	"strings"
	"testing"
	"wb_lvl2/develop/dev05/internal/grep"
)

// Проверка работы обработчика флагов -v, -i и -F
func Test_checkLine(t *testing.T) {
	data := []string{"test", "TEST", "Wrong"}
	c1 := grep.NewConfig()
	c1.Flagv = true
	c1.Pattern = "TEST"
	s1 := grep.Source{
		Config: c1,
	}
	c2 := grep.NewConfig()
	c2.Flagi = true
	c2.Pattern = "TEST"
	s2 := grep.Source{
		Config: c2,
	}
	c3 := grep.NewConfig()
	c3.FlagF = true
	c3.Pattern = "TEST"
	s3 := grep.Source{
		Config: c3,
	}

	r11 := s1.CheckLine(data[0])
	r12 := s1.CheckLine(data[1])
	r13 := s1.CheckLine(data[2])
	r21 := s2.CheckLine(data[0])
	r22 := s2.CheckLine(data[1])
	r23 := s2.CheckLine(data[2])
	r31 := s3.CheckLine(data[0])
	r32 := s3.CheckLine(data[1])
	r33 := s3.CheckLine(data[2])

	if r11 != true || r12 != false || r13 != true {
		t.Fatal("Incorrect work of flag -v\n")
	}
	if r21 != true || r22 != true || r23 != false {
		t.Fatal("Incorrect work of flag -i\n")
	}
	if r31 != false || r32 != true || r33 != false {
		t.Fatal("Incorrect work of flag -F\n")
	}
}

// Проверка работы обработчика флага -n и вывода строчки
func Test_printLine(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.AddLine("TEST")
	exp1 := "stdio:TEST\n"
	exp2 := "stdio:1:TEST\n"
	exp3 := "testfile.txt:TEST\n"
	exp4 := "testfile.txt:1:TEST\n"

	err1 := s.PrintLine(&buff, 0)
	testData1 := buff.String()
	buff.Reset()

	c.Flagn = true
	err2 := s.PrintLine(&buff, 0)
	testData2 := buff.String()
	buff.Reset()

	c.Flagn = false
	s.Path = "testfile.txt"
	err3 := s.PrintLine(&buff, 0)
	testData3 := buff.String()
	buff.Reset()

	c.Flagn = true
	err4 := s.PrintLine(&buff, 0)
	testData4 := buff.String()
	buff.Reset()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		t.Fatal("error while running source.printLine()\n")
	}
	if exp1 != testData1 {
		t.Fatalf("incorrect output, expected: %s, got: %s\n", exp1, testData1)
	}
	if exp2 != testData2 {
		t.Fatalf("incorrect output or flag -n work, expected: %s, got: %s\n", exp2, testData2)
	}
	if exp3 != testData3 {
		t.Fatalf("incorrect output, expected: %s, got: %s\n", exp3, testData3)
	}
	if exp4 != testData4 {
		t.Fatalf("incorrect output or flag -n work, expected: %s, got: %s\n", exp3, testData3)
	}
}

// Проверка работы вывода счетчика
func Test_printCounter(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
		Path:   "testpath.abc",
	}
	expected := "testpath.abc:COUNT:3\n"

	err := s.PrintCounter(&buff, 3)
	testData := buff.String()

	if err != nil {
		t.Fatal("error while running source.printLine()\n")
	}
	if expected != testData {
		t.Fatalf("incorrect output, expected: %s, got: %s\n", expected, testData)
	}
}

// Проверка работы флага -A в работе с файлом, если строк снизу достаточно
func Test_fileRunA1(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedA1 := []string{"textFiles/test2.txt:nice", "textFiles/test2.txt:pink", "textFiles/test2.txt:pick",
		"textFiles/test2.txt:feed", ""}

	c.FlagA = 3
	c.Pattern = "nice"
	errA1 := s.FileRun(&buff)
	testDataA1 := strings.Split(buff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.fileRun() with -A flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -A, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -A, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -A в работе с файлом, если строк снизу НЕ достаточно
func Test_fileRunA2(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedA2 := []string{"textFiles/test2.txt:LaMP", "textFiles/test2.txt:aletsa", ""}
	c.FlagA = 3
	c.Pattern = "LaMP"

	errA2 := s.FileRun(&buff)
	testDataA2 := strings.Split(buff.String(), "\n")

	if errA2 != nil {
		t.Fatal("incorrect work of source.fileRun() with -A flag\n")
	}
	if len(testDataA2) != len(expectedA2) {
		t.Fatalf("incorrect length of output with flag -A, expected: %d, got: %d\n",
			len(expectedA2), len(testDataA2))
	} else {
		for i := range expectedA2 {
			if expectedA2[i] != testDataA2[i] {
				t.Fatalf("incorrect output with flag -A, expected: %s, got: %s\n", expectedA2[i], testDataA2[i])
			}
		}
	}
}

// Проверка работы флага -B в работе с файлом, если строк сверху достаточно
func Test_fileRunB1(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedB1 := []string{"textFiles/test2.txt:nest", "textFiles/test2.txt:next", "textFiles/test2.txt:little",
		"textFiles/test2.txt:nice", ""}

	c.FlagB = 3
	c.Pattern = "nice"
	errB1 := s.FileRun(&buff)
	testDataB1 := strings.Split(buff.String(), "\n")

	if errB1 != nil {
		t.Fatal("incorrect work of source.fileRun() with -B flag\n")
	}
	if len(testDataB1) != len(expectedB1) {
		t.Fatalf("incorrect length of output with flag -B, expected: %d, got: %d\n",
			len(expectedB1), len(testDataB1))
	} else {
		for i := range expectedB1 {
			if expectedB1[i] != testDataB1[i] {
				t.Fatalf("incorrect output with flag -B, expected: %s, got: %s\n", expectedB1[i], testDataB1[i])
			}
		}
	}
}

// Проверка работы флага -B в работе с файлом, если строк сверху НЕ достаточно
func Test_fileRunB2(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedB2 := []string{"textFiles/test2.txt:red", "textFiles/test2.txt:word", ""}

	c.FlagB = 3
	c.Pattern = "word"
	errB2 := s.FileRun(&buff)
	testDataB2 := strings.Split(buff.String(), "\n")

	if errB2 != nil {
		t.Fatal("incorrect work of source.fileRun() with -B flag\n")
	}
	if len(testDataB2) != 3 {
		t.Fatalf("incorrect length of output with flag -B, expected: %d, got: %d\n",
			len(expectedB2), len(testDataB2))
	} else {
		for i := range expectedB2 {
			if expectedB2[i] != testDataB2[i] {
				t.Fatalf("incorrect output with flag -B, expected: %s, got: %s\n", expectedB2[i], testDataB2[i])
			}
		}
	}
}

// Проверка работы флага -С в работе с файлом
func Test_fileRunC1(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedC1 := []string{"textFiles/test2.txt:nest", "textFiles/test2.txt:next", "textFiles/test2.txt:little",
		"textFiles/test2.txt:nice", "textFiles/test2.txt:pink", "textFiles/test2.txt:pick",
		"textFiles/test2.txt:feed", ""}
	c.Pattern = "nice"
	c.FlagC = 3

	errC1 := s.FileRun(&buff)
	testDataC1 := strings.Split(buff.String(), "\n")

	if errC1 != nil {
		t.Fatal("incorrect work of source.fileRun() with -C flag\n")
	}
	if len(testDataC1) != len(expectedC1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedC1), len(testDataC1))
	} else {
		for i := range expectedC1 {
			if expectedC1[i] != testDataC1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedC1[i], testDataC1[i])
			}
		}
	}
}

// Проверка работы связки флагов -С и -A в работе с файлом
func Test_fileRunC2(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedC2 := []string{"textFiles/test2.txt:nest", "textFiles/test2.txt:next", "textFiles/test2.txt:little",
		"textFiles/test2.txt:nice", "textFiles/test2.txt:pink", ""}
	c.Pattern = "nice"
	c.FlagC = 3
	c.FlagA = 1 // флаг -С активен

	errC2 := s.FileRun(&buff)
	testDataC2 := strings.Split(buff.String(), "\n")

	if errC2 != nil {
		t.Fatal("incorrect work of source.fileRun() with -A and -С flags\n")
	}
	if len(testDataC2) != len(testDataC2) {
		t.Fatalf("incorrect length of output with flags -A and -C, expected: %d, got: %d\n",
			len(expectedC2), len(testDataC2))
	} else {
		for i := range expectedC2 {
			if expectedC2[i] != testDataC2[i] {
				t.Fatalf("incorrect output with flags -A and -C, expected: %s, got: %s\n",
					expectedC2[i], testDataC2[i])
			}
		}
	}
}

// Проверка работы связки флагов -С и -B в работе с файлом
func Test_fileRunC3(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedC3 := []string{"textFiles/test2.txt:little", "textFiles/test2.txt:nice", "textFiles/test2.txt:pink",
		"textFiles/test2.txt:pick", "textFiles/test2.txt:feed", ""}
	c.Pattern = "nice"
	c.FlagC = 3
	c.FlagB = 1 // флаг -С активен

	errC3 := s.FileRun(&buff)
	testDataC3 := strings.Split(buff.String(), "\n")

	if errC3 != nil {
		t.Fatal("incorrect work of source.fileRun() with -B and -C flags\n")
	}
	if len(testDataC3) != len(expectedC3) {
		t.Fatalf("incorrect length of output with flags -B and -C, expected: %d, got: %d\n",
			len(expectedC3), len(testDataC3))
	} else {
		for i := range expectedC3 {
			if expectedC3[i] != testDataC3[i] {
				t.Fatalf("incorrect output with flags -B and -C, expected: %s, got: %s\n",
					expectedC3[i], testDataC3[i])
			}
		}
	}

}

// Проверка работы флага -c в работе с файлом
func Test_fileRunc1(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedc := []string{"textFiles/test2.txt:COUNT:2", ""}

	c.Flagc = true
	c.Pattern = "ace"
	errc := s.FileRun(&buff)
	testDatac := strings.Split(buff.String(), "\n")

	if errc != nil {
		t.Fatal("incorrect work of source.fileRun() with -c flag\n")
	}
	if len(testDatac) != len(expectedc) {
		t.Fatalf("incorrect length of output with flag -c, expected: %d, got: %d\n",
			len(expectedc), len(testDatac))
	} else {
		for i := range expectedc {
			if expectedc[i] != testDatac[i] {
				t.Fatalf("incorrect output with flag -c, expected: %s, got: %s\n",
					expectedc[i], testDatac[i])
			}
		}
	}
}

// Проверка работы связки флагов -c и -v в работе с файлом
func Test_fileRunc2(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedc := []string{"textFiles/test2.txt:COUNT:23", ""}

	c.Flagc = true
	c.Flagv = true // инвертированный счетчик
	c.Pattern = "ace"
	errc := s.FileRun(&buff)
	testDatac := strings.Split(buff.String(), "\n")

	if errc != nil {
		t.Fatal("incorrect work of source.fileRun() with -c flag\n")
	}
	if len(testDatac) != len(expectedc) {
		t.Fatalf("incorrect length of output with flag -c, expected: %d, got: %d\n",
			len(expectedc), len(testDatac))
	} else {
		for i := range expectedc {
			if expectedc[i] != testDatac[i] {
				t.Fatalf("incorrect output with flag -c, expected: %s, got: %s\n",
					expectedc[i], testDatac[i])
			}
		}
	}
}

// Проверка работы флага -i в работе с файлом
func Test_fileRuni(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedi := []string{"textFiles/test2.txt:lamp", "textFiles/test2.txt:LaMP", ""}

	c.Flagi = true
	c.Pattern = "lAmP"
	erri := s.FileRun(&buff)
	testDatai := strings.Split(buff.String(), "\n")

	if erri != nil {
		t.Fatal("incorrect work of source.fileRun() with -i flag\n")
	}
	if len(testDatai) != len(expectedi) {
		t.Fatalf("incorrect length of output with flag -i, expected: %d, got: %d\n",
			len(expectedi), len(testDatai))
	} else {
		for i := range expectedi {
			if expectedi[i] != testDatai[i] {
				t.Fatalf("incorrect output with flag -i, expected: %s, got: %s\n",
					expectedi[i], testDatai[i])
			}
		}
	}
}

// Проверка работы флага -i в работе с файлом
func Test_fileRunv(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedv := []string{"textFiles/test2.txt:word", "textFiles/test2.txt:pink", "textFiles/test2.txt:pick",
		"textFiles/test2.txt:quick", "textFiles/test2.txt:suit", "textFiles/test2.txt:lamp",
		"textFiles/test2.txt:stunt", "textFiles/test2.txt:toast", "textFiles/test2.txt:and",
		"textFiles/test2.txt:cook", "textFiles/test2.txt:LaMP", ""}

	c.Flagv = true
	c.Pattern = "e"
	errv := s.FileRun(&buff)
	testDatav := strings.Split(buff.String(), "\n")

	if errv != nil {
		t.Fatal("incorrect work of source.fileRun() with -v flag\n")
	}
	if len(testDatav) != len(expectedv) {
		t.Fatalf("incorrect length of output with flag -v, expected: %d, got: %d\n",
			len(expectedv), len(testDatav))
	} else {
		for i := range expectedv {
			if expectedv[i] != testDatav[i] {
				t.Fatalf("incorrect output with flag -v, expected: %s, got: %s\n",
					expectedv[i], testDatav[i])
			}
		}
	}
}

// Проверка работы флага -F в работе с файлом
func Test_fileRunF(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedF := []string{"textFiles/test2.txt:lets", "textFiles/test2.txt:aletsa", ""}

	c.FlagF = true
	c.Pattern = "lets"
	errF := s.FileRun(&buff)
	testDataF := strings.Split(buff.String(), "\n")

	if errF != nil {
		t.Fatal("incorrect work of source.fileRun() with -F flag\n")
	}
	if len(testDataF) != len(expectedF) {
		t.Fatalf("incorrect length of output with flag -F, expected: %d, got: %d\n",
			len(expectedF), len(testDataF))
	} else {
		for i := range expectedF {
			if expectedF[i] != testDataF[i] {
				t.Fatalf("incorrect output with flag -F, expected: %s, got: %s\n",
					expectedF[i], testDataF[i])
			}
		}
	}
}

// Проверка работы флага -n в работе с файлом
func Test_fileRunn(t *testing.T) {
	var buff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	s.Path = "textFiles/test2.txt"
	expectedn := []string{"textFiles/test2.txt:19:toast", ""}

	c.Flagn = true
	c.Pattern = "toast"
	errn := s.FileRun(&buff)
	testDatan := strings.Split(buff.String(), "\n")

	if errn != nil {
		t.Fatal("incorrect work of source.fileRun() with -n flag\n")
	}
	if len(testDatan) != len(expectedn) {
		t.Fatalf("incorrect length of output with flag -n, expected: %d, got: %d\n",
			len(expectedn), len(testDatan))
	} else {
		for i := range expectedn {
			if expectedn[i] != testDatan[i] {
				t.Fatalf("incorrect output with flag -n, expected: %s, got: %s\n",
					expectedn[i], testDatan[i])
			}
		}
	}
}

// Проверка работы флага -A в работе с STDIO, если вводимых строк после достаточно
func Test_stdioRunA1(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"little", "nice", "pink", "pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:nice", "stdio:pink", "stdio:pick", "stdio:feed", ""}

	c.FlagA = 3
	c.Pattern = "nice"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -A flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -A, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -A, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -A в работе с STDIO, если вводимых строк после НЕ достаточно
func Test_stdioRunA2(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"little", "nice", "pink"}
	expectedA1 := []string{"stdio:nice", "stdio:pink", ""}

	c.FlagA = 3
	c.Pattern = "nice"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -A flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -A, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -A, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -B в работе с STDIO, если вводимых строк до достаточно
func Test_stdioRunB1(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"little", "nice", "pink", "pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:nice", "stdio:pink", "stdio:pick", "stdio:feed", ""}

	c.FlagB = 3
	c.Pattern = "feed"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -B flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -B, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -B, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -B в работе с STDIO, если вводимых строк до НЕ достаточно
func Test_stdioRunB2(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:pick", "stdio:feed", ""}

	c.FlagB = 3
	c.Pattern = "feed"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -B flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -B, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -B, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -С в работе с STDIO
func Test_stdioRunC1(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "next", "little", "nice", "pink", "pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:word", "stdio:nest", "stdio:next", "stdio:little", "stdio:nice", "stdio:pink",
		"stdio:pick", ""}

	c.FlagC = 3
	c.Pattern = "little"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флагов -С и -A в работе с STDIO
func Test_stdioRunC2(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "next", "little", "nice", "pink", "pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:word", "stdio:nest", "stdio:next", "stdio:little", "stdio:nice", "stdio:pink", ""}

	c.FlagA = 2
	c.FlagC = 3
	c.Pattern = "little"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флагов -С и -B в работе с STDIO
func Test_stdioRunC3(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "next", "little", "nice", "pink", "pick", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:nest", "stdio:next", "stdio:little", "stdio:nice", "stdio:pink", "stdio:pick", ""}

	c.FlagB = 2
	c.FlagC = 3
	c.Pattern = "little"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -i в работе с STDIO
func Test_stdioRuni(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "Pie", "little", "nice", "pInk", "PIck", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:Pie", "stdio:pInk", "stdio:PIck", ""}

	c.Flagi = true
	c.Pattern = "pi"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -v в работе с STDIO
func Test_stdioRunv(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"nest", "pie", "little", "nice", "pink", "pick", "feed"}
	expectedA1 := []string{"stdio:nest", "stdio:little", "stdio:nice", "stdio:feed", ""}

	c.Flagv = true
	c.Pattern = "pi"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -F в работе с STDIO
func Test_stdioRunF(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "Pi", "little", "nice", "pInk", "PIck", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:Pi", ""}

	c.FlagF = true
	c.Pattern = "Pi"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}

// Проверка работы флага -n в работе с STDIO
func Test_stdioRunn(t *testing.T) {
	var wBuff, rBuff bytes.Buffer
	c := grep.NewConfig()
	s := grep.Source{
		Config: c,
	}
	inputA1 := []string{"red", "word", "nest", "Pi", "little", "nice", "pInk", "PIck", "feed", "pace", "race"}
	expectedA1 := []string{"stdio:3:nest", ""}

	c.Flagn = true
	c.Pattern = "nest"
	for _, v := range inputA1 {
		rBuff.WriteString(v + "\n")
	}
	errA1 := s.StdioRun(&wBuff, &rBuff)
	testDataA1 := strings.Split(wBuff.String(), "\n")

	if errA1 != nil {
		t.Fatal("incorrect work of source.stdioRun() with -C flag\n")
	}
	if len(testDataA1) != len(expectedA1) {
		t.Fatalf("incorrect length of output with flag -C, expected: %d, got: %d\n",
			len(expectedA1), len(testDataA1))
	} else {
		for i := range expectedA1 {
			if expectedA1[i] != testDataA1[i] {
				t.Fatalf("incorrect output with flag -C, expected: %s, got: %s\n", expectedA1[i], testDataA1[i])
			}
		}
	}
}
