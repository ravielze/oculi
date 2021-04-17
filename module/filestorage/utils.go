package filestorage

import (
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GenerateFileName(ext string) string {
	rBase := int(time.Now().Unix() % 26)
	head := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), (11 + rBase)))
	return head + "-" + uuid.NewV4().String() + ext
}
