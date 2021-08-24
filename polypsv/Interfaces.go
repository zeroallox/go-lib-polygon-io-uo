package polypsv

type PSVer interface {
	ToPSV(writer *Writer) error
	FromPSV([][]byte) error
}
