package helpers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	filepath2 "path/filepath"
)

type storageHelper struct{}

type S3Helper struct {
	Client     *s3.Client
	BucketName string
}

func NewS3Helper(bucketName string) (*S3Helper, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("your-region"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Helper{
		Client:     client,
		BucketName: bucketName,
	}, nil
}

func NewStorageHelper() *storageHelper {
	return &storageHelper{}
}

func (s *S3Helper) UploadFile(file multipart.File, filename string) (string, error) {
	// Create a unique filename if needed, e.g., with timestamp
	key := fmt.Sprintf("news_category/%s", filename)

	// Upload input
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Generate a URL to access the file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, "your-region", key)
	return url, nil
}

const (
	storageFolder = "storages/app/"
)

type StorageHelper interface {
	UploadFile(file multipart.File, header *multipart.FileHeader, filepath string) error
	UploadCategoryNews(file multipart.File, header *multipart.FileHeader) (error, string)
	GetPullFilePath(directory string, filename string) string
	GetNewsCategoryPath(filename string) string
}

func (h storageHelper) UploadFile(file multipart.File, header *multipart.FileHeader, filepath string) error {
	storagePath := storageFolder + filepath
	outFile, err := os.Create(storagePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		return fmt.Errorf("could not copy file: %w", err)
	}
	return nil
}

func (h storageHelper) UploadFileToS3(file multipart.File, header *multipart.FileHeader, filepath string) error {
	return nil
}

func (h storageHelper) GetNewsCategoryPath(filename string) string {
	return NewsCategoryPath(filename)
}

func (h storageHelper) GetPublicUrl(filepath string) string {
	return os.Getenv("STORAGE_URL") + "/file/" + filepath
}

func (h storageHelper) UploadCategoryNews(file multipart.File, header *multipart.FileHeader) (error, string) {
	filepath := NewsCategoryPath(uuid.NewString() + filepath2.Ext(header.Filename))
	err := h.UploadFile(file, header, filepath)
	if err != nil {
		return err, ""
	}
	return nil, filepath
}

func NewsCategoryPath(filename string) string {
	if filename == "" {
		return "news_category"
	}
	return "news_category/" + filename
}

func (h storageHelper) GetPullFilePath(directory string, filename string) string {
	return storageFolder + "" + directory + "/" + filename
}
