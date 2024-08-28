package menu_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/utils/encapsulation/resp"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuAllResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuImageListView
// @Tags Menu Management
// @Summary List all menus with banners
// @Description Retrieves a complete list of menus along with their associated banners, sorted by the banner sort order.
// @Accept json
// @Produce json
// @Router /api/menus [get]
// @Success 200 {array} MenuAllResponse "A list of all menus including associated banners."
// @Failure 500 {object} resp.Response{} "Internal server error."
func (MenuApi) MenuImageListView(c *gin.Context) {
	//先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Find(&menuList).Select("id").Scan(&menuIDList)
	//查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)

	var menus []MenuAllResponse
	for _, model := range menuList {
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuAllResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	resp.OkWithList(menus, int64(len(menus)), c)
	return
}
