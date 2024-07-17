package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils"
	"gvb_server/utils/res"
	"os"
	"path"
	"strings"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImagesUploadView upload multiple images and returns the url of the image
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("upload multiple files fail", c)
		return
	}
	// get a list of images
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("images is not exist", c)
		return
	}
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			res.FailWithMsg(fmt.Sprintf(
				"create dir fail", basePath), c)
		}
		return
	}

	var fileResponse []FileUploadResponse
	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)
		fileSize := float64(file.Size) / float64(1024*1024)
		// 获取文件后缀名
		fileNameList := strings.Split(file.Filename, ".")
		suffix := fileNameList[len(fileNameList)-1]
		flag := utils.In_list(suffix, WhiteImageList)
		if !flag {
			fileResponse = append(fileResponse, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("当前图片类型为%s属于不合法的文件类型", suffix),
			})
			continue
		}
		if fileSize > float64(global.Config.Upload.Size) {
			fileResponse = append(fileResponse, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("The upload file size is %.2f.Set file size is %d", file.Size, global.Config.Upload.Size),
			})
			continue
		}
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			fileResponse = append(fileResponse, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: true,
				Msg:       fmt.Sprintf("upload file success but save file fail"),
			})
			continue
		}
		fileResponse = append(fileResponse, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       fmt.Sprintf("upload file success"),
		})
	}
	res.OkWithData(fileResponse, c)
}
