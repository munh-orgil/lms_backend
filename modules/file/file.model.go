package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(fileHeader *multipart.FileHeader, c *fiber.Ctx) (string, error) {
	_, err := openOrMkdir(ROOT_DIR)
	if err != nil {
		return "", err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	// if !isFileAllowed(fileBytes) {
	// 	return "", fmt.Errorf(DISALLOWED_FILE_SIZE)
	// }

	ufile := strings.Split(fileHeader.Filename, ".")
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	newFileName := uuid
	if len(ufile) > 1 {
		newFileName = uuid + "." + ufile[len(ufile)-1]
	}

	filePath := filepath.Join(ROOT_DIR, newFileName)
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	if _, err := f.Write(fileBytes); err != nil {
		return "", err
	}

	return newFileName, nil
}

func openOrMkdir(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		os.Mkdir(path, 0777)
		if file, err = os.Open(path); err != nil {
			return nil, err
		}
	}
	return file, nil
}

func isFileAllowed(fileBytes []byte) bool {
	size := len(fileBytes)
	return size <= MAX_FILE_SIZE
}
