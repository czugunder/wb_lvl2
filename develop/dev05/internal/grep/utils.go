package grep

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func stringToLower(s string) string {
	return strings.ToLower(s)
}

func stringContains(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

func setFurtherMap(m map[int]int, n, i int) {
	for j := i + 1; j < n+i+1; j++ {
		if _, ok := m[j]; ok {
			m[j]++
		} else {
			m[j] = 1
		}
	}
}

func addToSlice(s *[]string, line string) {
	*s = append(*s, line)
}

func ReadFile(path string) ([]string, error) {
	var line string
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			} else {
				if len(line) > 0 {
					addToSlice(&lines, line)
				}
				break
			}
		} else {
			if len(line) > 0 {
				addToSlice(&lines, line[:len(line)-1])
			}
		}
	}
	return lines, nil
}
