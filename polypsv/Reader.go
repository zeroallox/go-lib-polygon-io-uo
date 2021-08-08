package polypsv

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
)

type Reader struct {
	scn    *bufio.Scanner
	header string
}

func (rd *Reader) Header() string {
	return rd.header
}

func NewReader(reader io.Reader, isCompressed bool) (*Reader, error) {
	var n = new(Reader)

	if isCompressed == true {
		gzr, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}

		reader = gzr
	}

	n.scn = bufio.NewScanner(reader)

	headerLine, err := n.NextLine()
	if err != nil {
		return nil, err
	}

	n.header = string(bytes.Join(headerLine, pipeCharBytes))

	return n, nil
}

var ErrScanFailed = errors.New("scan failed")

func (rd *Reader) NextLine() ([][]byte, error) {

	if rd.scn.Scan() == false {

		var err = rd.scn.Err()
		if err == nil {
			return nil, io.EOF
		}

		return nil, err
	}

	return bytes.Split(rd.scn.Bytes(), pipeCharBytes), nil
}

func (rd *Reader) NextObject(obj PSVer) error {

	cLine, err := rd.NextLine()
	if err != nil {
		return err
	}

	return obj.FromPSV(cLine)
}
