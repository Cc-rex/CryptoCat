package article_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/es_service"
	"myServer/utils/encapsulation/resp"
	"time"
)

// ArticleUpdateView
// @Tags Article Management
// @Summary Update an article
// @Description Updates an existing article's details such as title, content, banner, and more based on provided values. Empty fields are ignored in the update process.
// @Accept json
// @Produce json
// @Param req body ctype.ArticleUpdateRequest true "Request body for updating an article including new values for fields like title, abstract, banner ID, and more"
// @Router /api/articles [put]
// @Success 200 {object} resp.Response{} "Confirms that the article has been successfully updated."
// @Failure 400 {object} resp.Response{} "Invalid request body or argument errors"
// @Failure 404 {object} resp.Response{} "Banner or article not found"
// @Failure 500 {object} resp.Response{} "Failed to update the article due to internal server error"
func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ctype.ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		resp.FailWithError(err, &cr, c)
		return
	}
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			resp.FailWithMsg("banner不存在", c)
			return
		}
	}

	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}

	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}

	err = es_service.ArticleUpdate(cr.ID, maps)
	if err != nil {
		resp.FailWithMsg("文章更新失败", c)
		return
	}
	resp.OkWithMsg("更新成功", c)
}
