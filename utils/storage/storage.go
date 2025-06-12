package storage

type FileStorage interface {
	Upload(localFilePath string, saveName string) (string, string, error)
	Delete(filePath string) error
}

// NewFileStorage 创建文件存储实例
func NewFileStorage(storageType string, basePath string) FileStorage {
	switch storageType {
	case "local":
		return &LocalFileStorage{
			BasePath: basePath,
		}
	default:
		return nil
	}
}
