package polypsv

type PSVer interface {
	ToPSV(writer *PSVWriter) error
	FromPSV([][]byte) error
}
