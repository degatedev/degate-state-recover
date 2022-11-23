package util

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil{
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadTxtFile(path string) (string, error) {
	content, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
	return string(content), err
}

func WriteTxtFile(info string, path string) error {
	fi, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fi.Close()

	_, err = fi.WriteString(info)
	if err != nil {
		return err
	}
	return nil
}

func FileRename(originName, targetName string) error {
	originExist, _ := PathExists(originName)
	if !originExist {
		return errors.New("in FileRename origin name is not exist.")
	}
	return os.Rename(originName, targetName)
}