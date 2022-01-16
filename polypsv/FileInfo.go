package polypsv

import (
    "errors"
    "github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
    "github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
    "path/filepath"
    "regexp"
    "time"
)

var regexFileName = regexp.MustCompile(`(?m)(\w+)-(\w+)-(\w+)-(\d+-\d+-\d+)\.psv`)

type FileInfo struct {
    locale     polyconst.Locale
    market     polyconst.Market
    dataType   polyconst.DataType
    date       time.Time
    compressed bool
}

func NewFileInfo(locale polyconst.Locale, market polyconst.Market, dataType polyconst.DataType, date time.Time) *FileInfo {

    var n = new(FileInfo)

    n.locale = locale
    n.market = market
    n.dataType = dataType
    n.date = date

    return n
}

var ErrInvalidFileName = errors.New("invalid file name")

func NewFileFromPath(path string) (*FileInfo, error) {

    var fileName = filepath.Base(path)

    var parts = regexFileName.FindStringSubmatch(fileName)

    if len(parts) != 5 {
        return nil, ErrInvalidFileName
    }

    var locale = polyconst.LocaleFromString(parts[1])
    if locale == polyconst.LOC_Invalid {
        return nil, ErrInvalidFileName
    }

    var market = polyconst.MarketFromString(parts[2])
    if market == polyconst.MKT_Invalid {
        return nil, ErrInvalidFileName
    }

    var dataType = polyconst.DataTypeFromCode(parts[3])
    if dataType == polyconst.DT_Invalid {
        return nil, ErrInvalidFileName
    }

    tz, err := getLocaleMarketTimeZone(locale, market)
    if err != nil {
        return nil, err
    }

    date, err := polyutils.TimeFromStringDate(parts[4], tz)
    if err != nil || date.IsZero() {
        return nil, ErrInvalidFileName
    }

    var compressed = false
    switch filepath.Ext(path) {
    case ".gz":
        compressed = true
    case ".psv":
        compressed = false
    default:
        return nil, ErrInvalidFileName
    }

    var n = new(FileInfo)
    n.locale = locale
    n.market = market
    n.dataType = dataType
    n.date = date
    n.locale = locale
    n.compressed = compressed

    return n, nil
}

func (file *FileInfo) Locale() polyconst.Locale {
    return file.locale
}

func (file *FileInfo) DataType() polyconst.DataType {
    return file.dataType
}

func (file *FileInfo) Market() polyconst.Market {
    return file.market
}

func (file *FileInfo) Date() time.Time {
    return file.date
}

func (file *FileInfo) Compressed() bool {
    return file.compressed
}
