package message_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
)

type MessageGroup map[uint]*ctype.Message

// MessageListView
// @Tags Message Management
// @Summary List messages grouped by conversation
// @Description Retrieves a list of messages for the current user, grouped by conversations with each contact, sorted by creation date in ascending order. Each group includes a summary of the conversation, like the number of messages.
// @Accept json
// @Produce json
// @Router /api/messages/list [get]
// @Success 200 {array} ctype.Message "List of grouped messages showing conversation summaries with each contact"
// @Failure 400 {object} resp.Response{} "Invalid authentication or session details"
// @Failure 500 {object} resp.Response{} "Failed to list messages due to internal server error"
func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []ctype.Message

	global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		message := ctype.Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			// 不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	resp.OkWithData(messages, c)
	return
}
