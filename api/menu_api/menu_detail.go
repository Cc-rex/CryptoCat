package menu_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/utils/encapsulation/resp"
)

// MenuDetailView
// @Tags Menu Management
// @Summary Get menu details
// @Description Retrieves detailed information about a specific menu including its associated banners sorted by their order.
// @Accept json
// @Produce json
// @Param id path string true "Menu ID"
// @Router /api/menus/{id} [get]
// @Success 200 {object} MenuAllResponse "Detailed information about the menu including associated banners."
// @Failure 404 {object} resp.Response{} "Menu does not exist."
func (MenuApi) MenuDetailView(c *gin.Context) {
	// 先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		resp.FailWithMsg("menu does not exist", c)
		return
	}
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuAllResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	resp.OkWithData(menuResponse, c)
	return
}
