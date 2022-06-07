package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNoFileToCopy          = errors.New("no file to copy")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if exists, _ := fileExists(fromPath); !exists {
		return ErrNoFileToCopy
	}

	size, err := fileSize(fromPath)
	if err != nil {
		return err
	}
	if size == 0 {
		return ErrUnsupportedFile
	}

	if size < offset {
		return ErrOffsetExceedsFileSize
	}

	toFile, err := os.Create(toPath)

	if err != nil {
		return err
	}
	defer toFile.Close()

	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	if limit == 0 || (limit > size-offset) {
		limit = size - offset
	}

	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	progress := pb.Full.Start64(limit)

	defer progress.Finish()

	progressReader := progress.NewProxyReader(fromFile)
	_, err = io.CopyN(toFile, progressReader, limit)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func fileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}
