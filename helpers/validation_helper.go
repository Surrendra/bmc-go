package helpers

import (
	"fmt"
	_ "io"
	"mime/multipart"
	"net/http"
)

type validationHelper struct{}

func NewValidationHelper() *validationHelper {
	return &validationHelper{}
}

type ValidationHelper interface {
	ValidateImage(file multipart.File, fileHeader *multipart.FileHeader) error
}

func (h validationHelper) ValidateImage(file multipart.File, fileHeader *multipart.FileHeader) error {
	const maxSize = 5 << 20
	if fileHeader.Size > maxSize {
		return fmt.Errorf("file size too big: %d maximal is : %d", fileHeader.Size, maxSize)
	}
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed read the file: %w", err)
	}

	contentType := http.DetectContentType(buffer)
	allowTypes := []string{"image/jpg", "image/png", "image/jpeg"}
	isAllowed := false
	for _, t := range allowTypes {
		if contentType == t {
			isAllowed = true
			break
		}
	}
	if isAllowed == false {
		return fmt.Errorf("file type %s is not allowed", contentType)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed seek the file: %w", err)
	}

	//img, _, err := image.Decode(file)
	//if err != nil {
	//	return fmt.Errorf("failed decode the file: %w", err)
	//}
	//
	//width := img.Bounds().Dx()
	//height := img.Bounds().Dy()
	//if width > 2000 || height > 2000 {
	//	return fmt.Errorf("image dimensions exceed the 2000x2000 limit")
	//}
	return nil
}
