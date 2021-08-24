package polypsv

import (
	"fmt"
	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"io"
	"strconv"
	"time"
)

type Writer struct {
	gzw            *gzip.Writer
	writer         io.Writer
	header         []byte
	buff           []byte
	sep            byte
	didWriteHeader bool
}

func NewPSVWriter(file *PSVFile, writer io.WriteCloser) (*Writer, error) {

	var n = new(Writer)
	n.writer = writer
	n.sep = pipeCharByte

	if file.Compressed() == true {

		gzr, err := gzip.NewWriterLevel(n.writer, flate.BestSpeed)
		if err != nil {
			return nil, err
		}

		gzr.Name = file.MakeFileName(false)
		gzr.Comment = fmt.Sprintf("CreatedAt: %d", time.Now().UnixNano())

		n.writer = gzr
		n.gzw = gzr
	}

	return n, nil
}

func (this *Writer) SetHeader(header []byte) {
	this.header = header
}

func (this *Writer) Header() []byte {
	return this.header
}

// Close Flushes the internal line buffer and cleans up.
// Does NOT close the underlying file handle / io.Writer.
func (this *Writer) Close() error {

	var err = this.writeBuffer(true)
	if err != nil {
		return err
	}

	if this.gzw != nil {
		err = this.gzw.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *Writer) WriteObject(psvItem PSVer) error {
	var err = psvItem.ToPSV(this)

	if err != nil {
		return err
	}

	return this.EndLine()
}

func (this *Writer) WriteNonEmptyString(value string) *Writer {

	if len(value) != 0 {
		this.buff = append(this.buff, value...)
	}

	return this
}

func (this *Writer) WriteBool(value bool) *Writer {

	var char byte = 'F'
	if value == true {
		char = 'T'
	}

	this.buff = append(this.buff, char)

	return this
}

func (this *Writer) WriteNonZeroFloat(value float64) *Writer {

	if value != 0 {
		this.buff = strconv.AppendFloat(this.buff, value, 'g', -1, 64)
	}

	return this
}

func (this *Writer) WriteNonZeroInt(value int64) *Writer {

	if value != 0 {
		this.buff = strconv.AppendInt(this.buff, value, 10)
	}

	return this
}

func (this *Writer) WriteNonZeroUint(value uint64) *Writer {

	if value != 0 {
		this.buff = strconv.AppendUint(this.buff, value, 10)
	}

	return this
}

func (this *Writer) WriteInt32Array(value []int32) *Writer {

	if len(value) >= 1 {

		this.buff = strconv.AppendInt(this.buff, int64(value[0]), 10)

		for _, cInt := range value[1:] {
			this.buff = append(this.buff, ';')
			this.buff = strconv.AppendInt(this.buff, int64(cInt), 10)
		}
	}

	return this
}

func (this *Writer) WriteIntArray(value []int64) *Writer {

	if len(value) >= 1 {
		this.buff = strconv.AppendInt(this.buff, value[0], 10)

		for _, cInt := range value[1:] {
			this.buff = append(this.buff, ';')
			this.buff = strconv.AppendInt(this.buff, cInt, 10)
		}
	}

	return this
}

func (this *Writer) Sep() *Writer {
	this.buff = append(this.buff, this.sep)
	return this
}

func (this *Writer) EndLine() error {

	if this.didWriteHeader == false {
		var err = this.writeHeader()
		if err != nil {
			return err
		}
	}

	this.buff = append(this.buff, '\n')

	var err = this.writeBuffer(false)
	if err != nil {
		return err
	}

	this.buff = this.buff[:0]

	return nil
}

func (this *Writer) writeBuffer(sync bool) error {

	_, err := this.writer.Write(this.buff)
	if err != nil {
		return err
	}

	if sync == true {
		if this.gzw != nil {
			err = this.gzw.Flush()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (this *Writer) writeHeader() error {

	this.didWriteHeader = true

	if len(this.header) == 0 {
		return nil
	}

	_, err := this.writer.Write(this.header)
	_, err = this.writer.Write(newLineCharBytes)
	if err != nil {
		return err
	}

	return nil

}
