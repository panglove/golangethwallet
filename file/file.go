package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

const (
	FILE = 1
	DIR  = 2

	NONE = 0
)

func ReadFileString(path string) (string, error) {

	res := ""

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	buffer := bufio.NewReader(f)

	for {

		str, err := buffer.ReadString('\n')

		res += str

		if err != nil {
			if err == io.EOF {

				break
			} else {

				return "", err
			}
		}
	}

	return res, nil

}
func WriteFileString(path string, content string) error {

	CreateFile(path,true)

	f, err := os.OpenFile(path,os.O_RDWR,os.ModePerm)

	defer f.Close()

	if err != nil {
		return err
	}

	_, err2 := f.WriteString(content)


	return err2

}
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func GetPathType(path string) int {

	if !PathExists(path) {
		return NONE
	}

	st, _ := os.Stat(path)

	if st.IsDir() {
		return DIR
	} else {
		return FILE
	}

}

func RemovePath(path string) bool {

	fileType := GetPathType(path)

	if fileType == NONE {
		return false
	}
	if fileType == FILE {

		err := os.Remove(path)

		return err == nil

	}

	if fileType == DIR {
		err := os.RemoveAll(path)
		return err == nil
	}
	return false
}

func GetFilePath(path string) (string, bool) {

	if GetPathType(path) == FILE {

		abPath, err := filepath.Abs(filepath.Dir(path))

		if err == nil {
			return abPath, true
		}

	}

	return "", false

}
func GetFileName(path string) string {

	if GetPathType(path) == FILE {

		_, file := filepath.Split(path)

		return file

	}

	return ""

}

//CREATE file

func CreateFile(filePath string, isCover bool) bool {

	if PathExists(filePath) {
		if isCover {

			isRm := RemovePath(filePath)

			if !isRm {
				return false
			}

		} else {
			return false
		}

	}
	_,err := os.Create(filePath)

	if err !=nil {
		return false
	}
	return true

}
func CreateDir(filePath string) bool {

	err := os.MkdirAll(filePath,0777)

	return  err==nil

}
