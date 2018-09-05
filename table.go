package tablego

import (
	"bufio"
	"io"
	"strings"
)

// Scanner - scan function center
type Scanner struct {
	Bufio    bufio.Scanner
	Spliter  string
	MaxLine  int
	Colnames []string
}

// default value of Scanner
const (
	spliter = "\t"
	maxLine = 0
)

// Input - input table file object(io.Reader)
func Input(r io.Reader) *Scanner {
	return &Scanner{
		Bufio:   *bufio.NewScanner(r),
		Spliter: spliter,
		MaxLine: maxLine,
	}
}

// Iter - return []string by channel
func (s *Scanner) Iter() <-chan []string {
	chnl := make(chan []string)
	countLine := 0
	go func() {
		for s.Bufio.Scan() {
			tmpLine := s.Bufio.Text()
			if (s.MaxLine != 0) && (countLine > s.MaxLine) {
				break
			} else if len(tmpLine) == 0 {
				continue
			} else if tmpLine[:1] == "#" {
				s.Colnames = strings.Split(tmpLine[1:], s.Spliter)
			} else {
				chnl <- strings.Split(s.Bufio.Text(), s.Spliter)
			}
			countLine++
		}
		close(chnl)
	}()
	return chnl
}
