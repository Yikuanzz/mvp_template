package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"mvp/utils/file"
)

type LocalFileStorage struct {
	BasePath string
}

// Upload 上传文件
func (l *LocalFileStorage) Upload(srcPath, saveName string) (string, string, error) {
	dstPath := filepath.Join(l.BasePath, saveName)
	dstPath, fileName := file.GetLocalAvailableFileName(dstPath)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", "", fmt.Errorf("无法打开源文件: %w", err)
	}
	defer srcFile.Close()

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
		return "", "", fmt.Errorf("创建目录失败: %w", err)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return "", "", fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", "", fmt.Errorf("文件写入失败: %w", err)
	}

	return dstPath, fileName, nil
}

// Delete 删除文件
func (l *LocalFileStorage) Delete(path string) error {
	return os.Remove(path)
}
