package conf

import (
	"path/filepath"
	"strings"
)

const NMS_UPLOAD_MAX_SIZE=100;

func GetUploadPath() string {
	storage := CommonConfig.GetString("storage.root")
	upload := CommonConfig.GetString("storage.upload")
	if strings.HasPrefix(upload, "/") {
		return upload
	} else {
		return filepath.Join(storage, upload)
	}
}
