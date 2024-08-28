package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils"
	"myServer/utils/encapsulation/resp"
	"os"
	"path"
	"strings"
)

var (
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

// ImageUploadView 上传单个图片，返回图片url
// @Tags Image Management
// @Summary Image Upload
// @Description Used for users to upload images
// @Accept multipart/form-data
// @Produce json
// @Param images formData file true "List of images to upload"
// @Router /api/images/ [post]
// @Success 200 {object} resp.Response{data=[]ctype.FileUploadResponse} "Upload successful"
// @Failure 400 {object} resp.Response{} "Invalid request"
// @Failure 500 {object} resp.Response{} "Internal server error"
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		resp.FailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		resp.FailWithMsg("Image does not exist", c)
		return
	}

	//判断路径是否存在，不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []ctype.FileUploadResponse

	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, ctype.FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "File type error",
			})
			continue
		}

		filePath := path.Join(basePath, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, ctype.FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("The image is too large. The current size is: %.2fMB, The legal size is：%dMB", size, global.Config.Upload.Size),
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		// Check if image exists
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			// find it
			resList = append(resList, ctype.FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "Image already exists",
			})
		} else {
			err = c.SaveUploadedFile(file, filePath)
			if err != nil {
				global.Log.Error(err)
				resList = append(resList, ctype.FileUploadResponse{
					FileName:  file.Filename,
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}
			resList = append(resList, ctype.FileUploadResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "Upload Success",
			})
			//Write to database
			global.DB.Create(&models.BannerModel{
				Path: filePath,
				Hash: imageHash,
				Name: fileName,
			})
		}

	}

	resp.OkWithData(resList, c)
}
