package menu_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// MenuCreateView
// @Tags Menu Management
// @Summary Create a new menu
// @Description Adds a new menu entry to the database with optional banner images if provided.
// @Accept json
// @Produce json
// @Param menu body ctype.MenuRequest true "Request body for creating a new menu entry"
// @Router /api/menus [post]
// @Success 200 {object} resp.Response{} "Menu created successfully."
// @Failure 400 {object} resp.Response{} "Invalid request body"
// @Failure 409 {object} resp.Response{} "Duplicate menu entries found"
// @Failure 500 {object} resp.Response{} "Failed to create menu due to internal server error"
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr ctype.MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}

	//重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		resp.FailWithMsg("Duplicate menus", c)
		return
	}

	// 创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		resp.OkWithMsg("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range cr.ImageSortList {
		// 这里也得判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("菜单图片关联失败", c)
		return
	}
	resp.OkWithMsg("菜单添加成功", c)
}
