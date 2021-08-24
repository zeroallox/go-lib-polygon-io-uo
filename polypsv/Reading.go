package polypsv

import (
	"bytes"
	"errors"
	"strconv"
	"unsafe"
)

func ReadString(src []byte) string {
	return string(src)
}

var ErrBadSourceData = errors.New("bad source data")

func ReadBool(src []byte) (bool, error) {

	if len(src) != 1 {
		return false, ErrBadSourceData
	}

	switch src[0] {
	case 'T':
		return true, nil
	case 'F':
		return false, nil
	default:
		return false, ErrBadSourceData
	}

}

func ReadFloat(src []byte) (float64, error) {

	if len(src) == 0 {
		return 0, nil
	}

	return strconv.ParseFloat(fastByteToString(src), 64)
}

func ReadInt(src []byte) (int64, error) {

	if len(src) == 0 {
		return 0, nil
	}

	return strconv.ParseInt(fastByteToString(src), 10, 64)
}

func ReadUint(src []byte) (uint64, error) {

	if len(src) == 0 {
		return 0, nil
	}

	return strconv.ParseUint(fastByteToString(src), 10, 64)
}

var pipeChar = []byte("|")
var semiColon = []byte(";")

func ReadIntArray(src []byte) ([]int64, error) {

	var arr []int64

	for _, cInt := range bytes.Split(src, semiColon) {

		if len(cInt) == 0 {
			continue
		}

		value, err := strconv.ParseInt(fastByteToString(cInt), 10, 64)
		if err != nil {
			return nil, err
		}

		arr = append(arr, value)

	}

	return arr, nil
}

func ReadInt32Array(src []byte) ([]int32, error) {

	var arr []int32

	for _, cInt := range bytes.Split(src, semiColon) {

		if len(cInt) == 0 {
			continue
		}

		value, err := strconv.ParseInt(fastByteToString(cInt), 10, 32)
		if err != nil {
			return nil, err
		}

		arr = append(arr, int32(value))

	}

	return arr, nil
}

func fastByteToString(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}
