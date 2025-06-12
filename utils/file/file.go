package file

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFileTypeByExt(fileName string) string {
	ext := filepath.Ext(fileName) // 如 ".png"
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream" // 默认类型
	}
	return mimeType
}

// 自动添加后缀防止文件重名（如 file.txt → file(1).txt）
func GetLocalAvailableFileName(path string) (string, string) {
	dir := filepath.Dir(path)             // 获取目录路径
	base := filepath.Base(path)           // 获取文件名
	ext := filepath.Ext(base)             // 获取文件扩展名
	name := strings.TrimSuffix(base, ext) // 获取文件的名称（不带扩展名）

	// 检查是否存在，存在则加后缀
	newPath := path
	newName := base
	i := 1
	for {
		// 判断文件是否存在
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			break
		}
		// 如果文件存在，生成新的文件名
		newName = fmt.Sprintf("%s(%d)%s", name, i, ext)
		newPath = filepath.Join(dir, newName)
		i++
	}

	// 返回新的文件路径和文件名
	return newPath, newName
}

// CleanDuplicateSuffix 去除文件名中的 (1)、(2) 等重复标记
func CleanDuplicateSuffix(filename string) string {
	ext := filepath.Ext(filename)                 // .docx
	base := strings.TrimSuffix(filename, ext)     // test(1)
	re := regexp.MustCompile(`^(.*?)(\(\d+\))?$`) // 匹配尾部 (数字)
	matches := re.FindStringSubmatch(base)

	if len(matches) > 1 {
		base = matches[1] // 去掉 (数字)
	}
	return base + ext
}
