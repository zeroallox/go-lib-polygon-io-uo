package polypsv

import (
    "errors"
    "github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
    "time"
)


var localeMarketTimeZoneMap = map[polyconst.Locale]map[polyconst.Market]*time.Location{
    polyconst.LOC_USA: {
        polyconst.MKT_Stocks: polyconst.NYCTime,
    },
}

func getLocaleMarketTimeZone(locale polyconst.Locale, market polyconst.Market) (*time.Location, error) {
    if locale == polyconst.LOC_Global {
        return time.UTC, nil
    }

    var tz = localeMarketTimeZoneMap[locale][market]
    if tz == nil {
        return nil, errors.New("unsupported time zone")
    }

    return tz, nil
}


