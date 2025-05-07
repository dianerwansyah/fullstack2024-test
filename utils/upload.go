package utils

import "fmt"

func UploadToS3(fileName string) string {
	return fmt.Sprintf("https://s3.amazonaws.com/mybucket/%s", fileName)
}
