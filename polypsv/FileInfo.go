package polypsv

import (
	"errors"
	"fmt"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var regexFileName = regexp.MustCompile(`(?m)(\w+)-(\w+)-(\w+)-(\d+-\d+-\d+)\.(gz|psv)`)

type FileInfo struct {
	locale     polyconst.Locale
	market     polyconst.Market
	dataType   polyconst.DataType
	date       time.Time
	compressed bool
}

func NewFileInfo(locale polyconst.Locale,
	market polyconst.Market,
	dataType polyconst.DataType,
	date time.Time,
	compressed bool) *FileInfo {

	var n = new(FileInfo)

	n.locale = locale
	n.market = market
	n.dataType = dataType
	n.date = date
	n.compressed = compressed

	return n
}

var ErrInvalidFileName = errors.New("invalid file name")

func NewFileFromPath(path string) (*FileInfo, error) {

	var fileName = filepath.Base(path)

	var parts = regexFileName.FindStringSubmatch(fileName)

	if len(parts) != 6 {
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

	date, err := polyutils.TimeFromStringDate(parts[4], localeTimeZoneMap[locale])
	if err != nil || date.IsZero() {
		return nil, ErrInvalidFileName
	}

	var compressed = false
	switch parts[5] {
	case "gz":
		compressed = true
	case "psv":
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

func (file *FileInfo) Market() polyconst.Market {
	return file.market
}

func (file *FileInfo) Date() time.Time {
	return file.date
}

func (file *FileInfo) Compressed() bool {
	return file.compressed
}

// MakeDirPath generates the directory for file.
//  Example:
//  polygon/us/stocks/trades/2000/2000-01
func MakeDirPath(file *FileInfo) string {

	year, month, _ := file.date.Date()

	return strings.ToLower(fmt.Sprintf("polygon/%v/%v/%v/%04d/%04d-%02d",
		file.locale.Code(),
		file.market.Code(),
		file.dataType.Code(),
		year,
		year, month))

}

// MakeFileName generates a file name for file.
func MakeFileName(file *FileInfo) string {
	return makeFileName(file, file.compressed)
}

// makeFileName returns the file name for a FileInfo. Compressed is
// specified separately and overrides file.compressed.
func makeFileName(file *FileInfo, compressed bool) string {

	year, month, day := file.date.Date()

	var ext = "psv"
	if compressed == true {
		ext = "gz"
	}

	return strings.ToLower(fmt.Sprintf("%s-%s-%s-%04d-%02d-%02d.%v",
		file.locale.Code(),
		file.market.Code(),
		file.dataType.Code(),
		year, month, day, ext))

}

// MakeABSFilePath returns the absolute file path for file.
//  Example:
//  polygon/us/stocks/trades/2000/2000-01/us-stocks-trades-2000-01-01.gz
func MakeABSFilePath(file *FileInfo) string {
	return filepath.Join(MakeDirPath(file), MakeFileName(file))
}
