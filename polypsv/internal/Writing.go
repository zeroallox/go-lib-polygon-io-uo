package internal

import "strconv"

func WriteNonEmptyString(dst *[]byte, value string) {
	if len(value) != 0 {
		*dst = append(*dst, value...)
	}
}

func WriteBool(dst *[]byte, value bool) {
	if value == true {
		*dst = append(*dst, 'T')
	} else {
		*dst = append(*dst, 'F')
	}
}

func WriteNonZeroFloat(dst *[]byte, value float64) {
	if value != 0 {
		*dst = strconv.AppendFloat(*dst, value, 'g', -1, 64)
	}
}

func WriteNonZeroInt(dst *[]byte, value int64) {
	if value != 0 {
		*dst = strconv.AppendInt(*dst, value, 10)
	}
}

func WriteNonZeroUint(dst *[]byte, value uint64) {
	if value != 0 {
		*dst = strconv.AppendUint(*dst, value, 10)
	}
}

func WriteInt32Array(dst *[]byte, intArr []int32) {
	if len(intArr) >= 1 {

		WriteNonZeroInt(dst, int64(intArr[0]))

		for _, cInt := range intArr[1:] {
			*dst = append(*dst, ';')
			WriteNonZeroInt(dst, int64(cInt))
		}
	}
}

func WriteIntArray(dst *[]byte, intArr []int64) {

	if len(intArr) >= 1 {

		WriteNonZeroInt(dst, intArr[0])

		for _, cInt := range intArr[1:] {
			*dst = append(*dst, ';')
			WriteNonZeroInt(dst, cInt)
		}
	}
}

func WriteNewLine(dst *[]byte) {
	*dst = append(*dst, '\n')
}

func WriteSep(dst *[]byte, sep byte) {
	*dst = append(*dst, sep)
}
