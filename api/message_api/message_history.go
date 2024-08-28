package message_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

// MessageHistoryView
// @Tags Message Management
// @Summary Retrieve message history
// @Description Retrieves the message history between the current user and another specified user, sorting messages by creation date in ascending order and marking them as read.
// @Accept json
// @Produce json
// @Param req body ctype.MessageHistoryRequest true "Request body for retrieving message history including the other user's ID"
// @Router /api/messages/history [get]
// @Success 200 {array} models.MessageModel "List of messages between the current user and the specified other user"
// @Failure 400 {object} resp.Response{} "Invalid request body or user ID"
// @Failure 500 {object} resp.Response{} "Failed to retrieve message history due to internal server error"
func (MessageApi) MessageHistoryView(c *gin.Context) {
	var cr ctype.MessageHistoryRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range _messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}

	// 点开消息，里面的每一条消息，都从未读变成已读

	resp.OkWithData(messageList, c)
}
