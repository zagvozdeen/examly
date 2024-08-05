package service

import (
	"fmt"
	"github.com/Den4ik117/examly/internal/model"
	"github.com/Den4ik117/examly/internal/repository"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const FileDir = "files"

type FileService struct {
	repo repository.Files
}

func NewFileService(repo repository.Files) *FileService {
	return &FileService{repo: repo}
}

func (s *FileService) UploadFile(file multipart.File, header *multipart.FileHeader) (*model.File, error) {
	err := os.MkdirAll(FileDir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newUUID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	filePath := fmt.Sprintf(
		"/%s/%s%s",
		FileDir,
		newUUID.String(),
		strings.ToLower(filepath.Ext(header.Filename)),
	)
	dst, err := os.Create("./public" + filePath)
	defer dst.Close()
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		return nil, err
	}

	fileModel := &model.File{
		UUID:       newUUID.String(),
		Content:    filePath,
		Size:       int(header.Size),
		MimeType:   header.Header.Get("Content-Type"),
		OriginName: header.Filename,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	id, err := s.repo.CreateFile(fileModel)
	if err != nil {
		return nil, err
	}

	fileModel.ID = id

	return fileModel, nil
}
