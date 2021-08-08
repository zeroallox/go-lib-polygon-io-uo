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

var regexFileName = regexp.MustCompile(`(?m)(\w+)-(\w+)-(\d+-\d+-\d+)\.(gz|psv)`)

type PSVFile struct {
	locale     polyconst.Locale
	market     polyconst.Market
	date       time.Time
	compressed bool
}

func NewPSVFile(locale polyconst.Locale,
	market polyconst.Market,
	date time.Time, compressed bool) *PSVFile {

	var n = new(PSVFile)

	n.locale = locale
	n.market = market
	n.date = date
	n.compressed = compressed

	return n
}

var ErrInvalidFileName = errors.New("invalid file name")

func NewFileFromPath(path string) (*PSVFile, error) {

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

	date, err := polyutils.TimeFromStringDate(parts[3], localeTimeZoneMap[locale])
	if err != nil || date.IsZero() {
		return nil, ErrInvalidFileName
	}

	var compressed = false
	switch parts[4] {
	case "gz":
		compressed = true
	case "psv":
		compressed = false
	default:
		return nil, ErrInvalidFileName
	}

	var n = new(PSVFile)
	n.market = market
	n.date = date
	n.locale = locale
	n.compressed = compressed

	return n, nil
}

func (file *PSVFile) Locale() polyconst.Locale {
	return file.locale
}

func (file *PSVFile) Market() polyconst.Market {
	return file.market
}

func (file *PSVFile) Date() time.Time {
	return file.date
}

func (file *PSVFile) Compressed() bool {
	return file.compressed
}

func (file *PSVFile) MakeDir() string {
	year, month, _ := file.date.Date()

	return strings.ToLower(fmt.Sprintf("polygon/%v/%v/%04d/%04d-%02d", file.locale.Code(), file.market.Code(),
		year,
		year, month))
}

func (file *PSVFile) MakeFileName(compressed bool) string {
	year, month, day := file.date.Date()

	var ext = "psv"
	if compressed == true {
		ext = "gz"
	}

	return strings.ToLower(fmt.Sprintf("%s-%s-%04d-%02d-%02d.%v",
		file.locale.Code(), file.market.Code(),
		year, month, day, ext))
}

func (file *PSVFile) MakeABSFilePath(compressed bool) string {
	return filepath.Join(file.MakeDir(), file.MakeFileName(compressed))
}
