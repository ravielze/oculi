package dvlp

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	storage "github.com/ramadani/go-filestorage"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"github.com/ravielze/fuzzy-broccoli/middleware"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DevModule struct {
}

func (DevModule) Name() string {
	return "Development Module"
}

type A struct {
	Attachment *multipart.FileHeader `json:"attachment,omitempty" form:"attachment" binding:"required"`
}

func NewDevModule(db *gorm.DB, g *gin.Engine) DevModule {
	config := &storage.Config{
		Root: "storage",
	}
	localStorage := storage.NewStorage(config)
	devGroup := g.Group("/development")
	{
		devGroup.GET("/test", middleware.GetStaticTokenMiddleware())
		devGroup.POST("/addfile", func(c *gin.Context) {
			var data A
			err := c.Bind(&data)
			if err != nil {
				fmt.Println(err)
				utils.AbortAndResponse(c, http.StatusBadRequest, err.Error())
				return
			}
			random := int(time.Now().Unix() % 26)
			fileName := strings.ToUpper(strconv.FormatInt(time.Now().Unix(), (11+random))) + "_" + uuid.NewV4().String() + filepath.Ext(data.Attachment.Filename)
			localStorage.PutFileAs("testing", data.Attachment, fileName)
			fmt.Println(fileName)
			x, _ := data.Attachment.Open()
			mime, _ := mimetype.DetectReader(x)
			fmt.Println(mime)
			fmt.Println(x.Seek(0, 2))
			x.Close()
			utils.OKAndResponseData(c, mime.String())
		})
	}
	return DevModule{}
}
