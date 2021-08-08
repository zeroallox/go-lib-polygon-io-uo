package polymodels

import (
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

// Compares two int32 arrays for equality.
func equalIntArr(src []int32, other []int32) bool {

	if (src == nil) != (other == nil) {
		return false
	}

	if len(src) != len(other) {
		return false
	}

	for i := range src {
		if src[i] != other[i] {
			return false
		}
	}

	return true
}

// hashInt32 Computes a simple hash from an array of int32's.
// Stolen from Java's HashCode function.
func hashInt32(ints []int32) uint64 {

	if ints == nil {
		return 0
	}

	var chk = uint64(0)
	for idx, cInt := range ints {
		chk = (31 * chk) + uint64(cInt+1) + uint64(idx)
	}

	return chk
}

// Struct fields to exclude or rename when generating PSV headers.
// The map is map[FieldValue] "-" (to exclude) or "replacement"
var replacements = map[string]string{
	"model": "-",      // Remove base model
	"trfid": "trf_id", // Rename "trfid" to "trf_id" for HistoricEquityTrade
}

// MakePSVHeader generates a PSV header for the passed in struct.
func MakePSVHeader(ifx interface{}) []byte {

	val := reflect.ValueOf(ifx)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var fields []string
	for i := 0; i < val.NumField(); i++ {

		var field = strcase.ToSnake(val.Type().Field(i).Name)
		var r = replacements[field]

		switch {
		case r == "-": // Exclude the field
			continue
		case r != "": // Replace with replacement
			field = r
			break
		}

		fields = append(fields, field)
	}

	return []byte(strings.Join(fields, "|"))
}
