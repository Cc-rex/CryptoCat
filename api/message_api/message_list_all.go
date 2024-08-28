package message_api

import (
	"github.com/gin-gonic/gin"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/resp"
)

// MessageListAllView
// @Tags Message Management
// @Summary List all messages
// @Description Retrieves a paginated list of all messages across the platform, including details for each message.
// @Accept json
// @Produce json
// @Param page query int false "Page number of the results"
// @Param size query int false "Number of results per page"
// @Router /api/messages_all [get]
// @Success 200 {array} models.MessageModel "Paginated list of all messages"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or request body"
// @Failure 500 {object} resp.Response{} "Failed to retrieve messages due to internal server error"
func (MessageApi) MessageListAllView(c *gin.Context) {
	var cr ctype.PageInfo
	if err := c.ShouldBindJSON(&cr); err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}
	list, count, _ := common.ListQuery(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})

	resp.OkWithList(list, count, c)
}
