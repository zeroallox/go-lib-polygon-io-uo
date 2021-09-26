package polyconst

import "strings"

type DataType uint8

const (
	DT_Invalid DataType = 0
	DT_Trades  DataType = 1
	DT_BBO     DataType = 2
	DT_NBBO    DataType = 3
	dt_max     DataType = iota
)

func (dt DataType) String() string {
	switch dt {
	case DT_Trades:
		return "DT_Trades"
	case DT_BBO:
		return "DT_BBO"
	case DT_NBBO:
		return "DT_NBBO"
	default:
		return "DT_Invalid"
	}
}

func (dt DataType) Code() string {
	switch dt {
	case DT_Trades:
		return "trades"
	case DT_BBO:
		return "bbo"
	case DT_NBBO:
		return "nbbo"
	default:
		return "invalid"
	}
}

func DataTypeFromCode(str string) DataType {
	str = strings.ToLower(str)

	for i := 0; i < int(dt_max); i++ {
		var cDataType = DataType(i)
		if cDataType.Code() == str {
			return cDataType
		}
	}

	return DT_Invalid
}
