package polypsv

import (
    "bytes"
    "github.com/klauspost/compress/gzip"
    "io"
    "os"
    "path/filepath"
)

// CreateLocalPSVFile creates the correct directory structure within baseDir
// and returns a handle to an empty io.File.
func CreateLocalPSVFile(baseDir string, file *FileInfo, compressed bool) (*os.File, error) {
    var err = os.MkdirAll(filepath.Join(baseDir, MakeDirPath(file)), os.ModePerm)
    if err != nil {
        return nil, err
    }

    return os.Create(filepath.Join(baseDir, MakeABSFilePath(file, compressed)))
}

// OpenLocalPSVFile returns a handle to file within baseDir.
//  Example:
//  If you store your CSV's in /data and want to open a us-stocks-trade file
//  we generate the correct path where the file should be and returns a handle
//  to it.
func OpenLocalPSVFile(baseDir string, file *FileInfo, compressed bool) (*os.File, error) {
    return os.Open(filepath.Join(baseDir, MakeABSFilePath(file, compressed)))
}

// GetLineCount returns the number of lines in the PSV file including the header.
func GetLineCount(hFile io.Reader, file *FileInfo) (uint, error) {

    var reader io.Reader
    if file.Compressed() == true {
        gzr, err := gzip.NewReader(hFile)
        if err != nil {
            return 0, err
        }
        reader = gzr
    } else {
        reader = hFile
    }

    count, err := countLines(reader)
    if err != nil {
        return 0, err
    }

    return uint(count), nil
}

// Stolen from: https://stackoverflow.com/questions/24562942
func countLines(r io.Reader) (int, error) {
    var buf = make([]byte, 32*1024)
    var count = 0

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], newLineCharBytes)

        switch {
        case err == io.EOF:
            return count, nil
        case err != nil:
            return count, err
        }
    }
}
