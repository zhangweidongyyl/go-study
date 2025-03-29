package commonfile

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

// DelTempFile 删除临时文件
func DelTempFile(ctx *gin.Context, localFilePath string) bool {
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return true
	}
	if err := os.Remove(localFilePath); err != nil {
		fmt.Printf("delTempFile has err:%+v", err)
		return false
	}
	return true
}

// GetFileNameByUrl 根据url获取文件名
func GetFileNameByUrl(ctx *gin.Context, url string) string {
	if url == "" {
		return ""
	}
	return filepath.Base(url)

}
