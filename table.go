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
	Colnames []string
}

// default value of Scanner
const (
	spliter = "\t"
)

// Input - input table file object(io.Reader)
func Input(r io.Reader) *Scanner {
	return &Scanner{
		Bufio:   *bufio.NewScanner(r),
		Spliter: spliter,
	}
}

// Iter - return []string by channel
func (s *Scanner) Iter() <-chan []string {
	chnl := make(chan []string)
	countLine := 0
	go func() {
		for s.Bufio.Scan() {
			tmpLine := s.Bufio.Text()
			if len(tmpLine) == 0 {
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

// GetColnames - only get the colnames
func (s *Scanner) GetColnames() []string {
	<-s.Iter()
	if s.Colnames == nil {
		return []string{}
	}
	return s.Colnames
}
