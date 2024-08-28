package menu_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// MenuNameListView
// @Tags Menu Management
// @Summary List menu names
// @Description Retrieves a list of all menus, providing basic details like ID, title, and path.
// @Accept json
// @Produce json
// @Router /api/menu_names [get]
// @Success 200 {array} ctype.MenuNameResponse "A list of menu names including IDs and paths."
// @Failure 500 {object} resp.Response "Internal server error."
func (MenuApi) MenuNameListView(c *gin.Context) {
	//先查菜单
	var menuNameList []ctype.MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	resp.OkWithData(menuNameList, c)

}
