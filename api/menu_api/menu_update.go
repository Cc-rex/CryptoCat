package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// MenuUpdateView
// @Tags Menu Management
// @Summary Update menu details
// @Description Updates the details of an existing menu and manages its associated banners based on the provided input.
// @Accept json
// @Produce json
// @Param id path string true "The ID of the menu to update"
// @Param menu body ctype.MenuRequest true "Menu update request body containing new details and banner information"
// @Router /api/menus{id} [put]
// @Success 200 {object} resp.Response{} "Update menu successfully."
// @Failure 400 {object} resp.Response{} "Invalid input."
// @Failure 404 {object} resp.Response{} "Menu does not exist."
// @Failure 500 {object} resp.Response{} "Failed to update menu."
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr ctype.MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	// 先把之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		resp.FailWithMsg("Menu does not exist", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 如果选择了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			resp.FailWithMsg("Failed to create a menu image", c)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("Failed to update menu", c)
		return
	}

	resp.OkWithMsg("Update menu successfully", c)
}
