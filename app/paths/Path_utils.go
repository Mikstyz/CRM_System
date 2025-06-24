package paths

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	once            sync.Once
	exeDir          string
	appData         string
	dbPath          string
	fontDir         string
	docDir          string
	embedTargetPath string
	embedFontPath   string
	backupPath      string
	errInit         error
)

func InitAppData() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")

	if localAppData == "" {
		return "", errors.New("переменная LOCALAPPDATA не установлена")
	}

	appDir := filepath.Join(localAppData, "SrmSystemKs54")

	err := os.MkdirAll(appDir, os.ModePerm)

	if err != nil {
		return "", err
	}

	return appDir, nil
}

func InitPaths() error {
	once.Do(func() {
		var exePath string
		exePath, errInit = os.Executable()
		if errInit != nil {
			return
		}

		appData, _ = InitAppData()

		exeDir = filepath.Dir(exePath)

		dbPath = filepath.Join(appData, "Data", "Sql", "stud.db")

		fontDir = filepath.Join(appData, "Data", "Fonts")
		docDir = filepath.Join(appData, "Document")

		backupPath = filepath.Join(appData, "Data", "backup")

		embedFontPath = filepath.Join(appData, "Data", "Fonts", "DejaVuSans.ttf")
		embedTargetPath = filepath.Join(appData, "Data", "Sql", "stud.db")
	})
	return errInit
}

func GetDBPath() string {
	return dbPath
}

func GetFontDir() string {
	return fontDir
}

func GetDocDir() string {
	return docDir
}

func GetEmbedTarget() string {
	return embedTargetPath
}

func GetEmbedFontPath() string {
	return embedFontPath
}

func GetbackupPath() string {
	return backupPath
}

func GetAppdata() string {
	return appData
}

func GetDocFile(fileName string) string {
	return filepath.Join(docDir, fileName)
}
