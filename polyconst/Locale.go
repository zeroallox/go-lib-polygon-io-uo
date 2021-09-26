package polyconst

import "strings"

type Locale uint8

const (
	LOC_Invalid      Locale = 0
	LOC_Global       Locale = 1
	LOC_USA          Locale = 2
	LOC_GreatBritain Locale = 3
	LOC_Canada       Locale = 4
	LOC_Netherlands  Locale = 5
	LOC_Greece       Locale = 6
	LOC_Spain        Locale = 7
	LOC_Germany      Locale = 8
	LOC_Belgium      Locale = 9
	LOC_Denmark      Locale = 10
	LOC_Finland      Locale = 11
	LOC_Ireland      Locale = 12
	LOC_Portugal     Locale = 13
	LOC_India        Locale = 14
	LOC_Mexico       Locale = 15
	LOC_France       Locale = 16
	LOC_China        Locale = 17
	LOC_Switzerland  Locale = 18
	LOC_Sweden       Locale = 19
	loc_max          Locale = iota
)

func (mkt Locale) Code() string {
	switch mkt {
	case LOC_Global:
		return "g"
	case LOC_USA:
		return "us"
	case LOC_GreatBritain:
		return "gb"
	case LOC_Canada:
		return "ca"
	case LOC_Netherlands:
		return "nl"
	case LOC_Greece:
		return "gr"
	case LOC_Spain:
		return "sp"
	case LOC_Germany:
		return "de"
	case LOC_Belgium:
		return "be"
	case LOC_Denmark:
		return "dk"
	case LOC_Finland:
		return "fi"
	case LOC_Ireland:
		return "ie"
	case LOC_Portugal:
		return "pt"
	case LOC_India:
		return "in"
	case LOC_Mexico:
		return "mx"
	case LOC_France:
		return "fr"
	case LOC_China:
		return "cn"
	case LOC_Switzerland:
		return "ch"
	case LOC_Sweden:
		return "se"
	default:
		return "_INVALID_Locale_"
	}
}

func LocaleFromString(str string) Locale {

	str = strings.ToLower(str)

	for i := 0; i < int(loc_max); i++ {
		var cLocale = Locale(i)
		if cLocale.Code() == str {
			return cLocale
		}
	}

	return LOC_Invalid
}
