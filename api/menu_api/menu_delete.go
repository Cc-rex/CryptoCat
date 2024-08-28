package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// MenuDeleteView deletes specified menu entries along with associated banners.
// @Tags Menu Management
// @Summary Delete menus
// @Description Deletes menu entries and their associated banners based on the provided list of menu IDs.
// @Accept json
// @Produce json
// @Param delete body ctype.DeleteRequest true "Request body for deleting menus, includes list of menu IDs"
// @Router /api/menus [delete]
// @Success 200 {object} resp.Response{} "Deletion successful with the number of menus deleted."
// @Failure 400 {object} resp.Response{} "Invalid request due to bad input"
// @Failure 404 {object} resp.Response{} "Menu not found"
// @Failure 500 {object} resp.Response{} "Failed to delete menu due to internal server error"
func (MenuApi) MenuDeleteView(c *gin.Context) {
	var cr ctype.DeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		resp.FailWithMsg("menu does not exist", c)
		return
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("Fail to delete menu", c)
		return
	}
	resp.OkWithMsg(fmt.Sprintf("Deleted %d menus", count), c)
}
