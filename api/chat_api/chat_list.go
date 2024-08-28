package chat_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/service/common"
	"myServer/utils/encapsulation/resp"
)

// ChatListView
// @Tags Chat Management
// @Summary List chat sessions
// @Description Retrieves a paginated list of chat sessions, specifically for group chats, sorted by creation date in descending order.
// @Accept json
// @Produce json
// @Param page query int false "Page number of the results"
// @Param size query int false "Number of results per page"
// @Router /api/chat_groups/records [get]
// @Success 200 {array} models.ChatModel "Paginated list of chat sessions with creation dates and other metadata"
// @Failure 400 {object} resp.Response{} "Invalid query parameters or binding errors"
// @Failure 500 {object} resp.Response{} "Failed to retrieve chat sessions due to internal server error"
func (ChatApi) ChatListView(c *gin.Context) {
	var cr ctype.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		resp.FailWithCode(resp.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"
	list, count, _ := common.ListQuery(models.ChatModel{IsGroup: true}, common.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		resp.OkWithList(list, count, c)
		return
	}
	resp.OkWithList(data, count, c)
}
