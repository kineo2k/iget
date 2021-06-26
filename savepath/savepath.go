package savepath

import (
	"fmt"
	"os"
	"path"
)

type SavePath struct {
	savePath string
}

func New(dirName string) *SavePath {
	savePath, err := savePath(dirName)
	if err != nil {
		panic(err)
	}

	return &SavePath{savePath}
}

func savePath(dirName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s%cDownloads%c%s", homeDir, os.PathSeparator, os.PathSeparator, dirName), nil
}

func (s *SavePath) WithUrl(urlString string) string {
	return fmt.Sprintf("%s%c%s", s.savePath, os.PathSeparator, path.Base(urlString))
}

func (s *SavePath) Create() error {
	if _, err := os.Stat(s.savePath); os.IsNotExist(err) {
		return os.Mkdir(s.savePath, os.ModePerm)
	}

	return nil
}
