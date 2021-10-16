package polypsv

import (
	"os"
	"path/filepath"
)

// CreateLocalPSVFile creates the correct directory structure within baseDir
// and returns a handle to an empty io.File.
func CreateLocalPSVFile(baseDir string, file *FileInfo) (*os.File, error) {
	var err = os.MkdirAll(filepath.Join(baseDir, MakeDirPath(file)), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(filepath.Join(baseDir, MakeABSFilePath(file)))
}

// OpenLocalPSVFile returns a handle to file within baseDir.
//  Example:
//  If you store your CSV's in /data and want to open a us-stocks-trade file
//  we generate the correct path where the file should be and returns a handle
//  to it.
func OpenLocalPSVFile(baseDir string, file *FileInfo) (*os.File, error) {
	return os.Open(filepath.Join(baseDir, MakeABSFilePath(file)))
}
