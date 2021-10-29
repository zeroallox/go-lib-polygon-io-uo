package polypsv

import (
    "github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
    "time"
)

var pipeCharBytes = []byte("|")
var pipeCharByte byte = '|'
var newLineCharBytes = []byte("\n")
var newLineCharByte byte = '\n'

var localeTimeZoneMap = map[polyconst.Locale]*time.Location{
    polyconst.LOC_USA: polyconst.NYCTime,
}
