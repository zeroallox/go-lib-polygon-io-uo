package polypsv

import (
	"os"
	"path/filepath"
)

// CreateLocalPSVFile creates the correct directory structure within baseDir
// and returns a handle to an empty io.File
func CreateLocalPSVFile(baseDir string, file *PSVFile) (*os.File, error) {
	var err = os.MkdirAll(filepath.Join(baseDir, file.MakeDir()), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(filepath.Join(baseDir, file.MakeABSFilePath(file.compressed)))
}

// OpenLocalFile returns a handle to
func OpenLocalPSVFile(baseDir string, file *PSVFile) (*os.File, error) {
	return os.Open(filepath.Join(baseDir, file.MakeABSFilePath(file.compressed)))
}
