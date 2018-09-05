package tablego

import (
	"bufio"
	"io"
	"strings"
)

// Scanner - scan function center
type Scanner struct {
	bufio   bufio.Scanner
	spliter string
}

// default value of Scanner
const (
	spliter = "\t"
)

// Input - input table file object(io.Reader)
func Input(r io.Reader) *Scanner {
	return &Scanner{
		bufio:   *bufio.NewScanner(r),
		spliter: spliter,
	}
}

// Iter - return []string by channel
func (s *Scanner) Iter() <-chan []string {
	chnl := make(chan []string)
	go func() {
		for s.bufio.Scan() {
			chnl <- strings.Split(s.bufio.Text(), s.spliter)
		}
		close(chnl)
	}()
	return chnl
}
