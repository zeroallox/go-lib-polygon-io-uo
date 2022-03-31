package polypsv

import (
	"fmt"
	"path/filepath"
	"strings"
)

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

// MakeFileName generates a file name for FileInfo. compressed overrides FileInfo.Compressed.
func MakeFileName(file *FileInfo, compressed bool) string {

	year, month, day := file.date.Date()

	var ext = "psv"
	if compressed == true {
		ext = ext + ".gz"
	}

	return strings.ToLower(fmt.Sprintf("%s-%s-%s-%04d-%02d-%02d.%v",
		file.locale.Code(),
		file.market.Code(),
		file.dataType.Code(),
		year, month, day, ext))

}

// MakeABSFilePath returns the absolute file path for file. compressed overrides FileInfo.Compressed.
//  Example:
//  polygon/us/stocks/trades/2000/2000-01/us-stocks-trades-2000-01-01.psv.gz
func MakeABSFilePath(file *FileInfo, compressed bool) string {
	return filepath.Join(MakeDirPath(file), MakeFileName(file, compressed))
}
