package message_api

import (
	"github.com/gin-gonic/gin"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
)

// MessageCreateView 发布消息
// @Tags Message Management
// @Summary Send a new message
// @Description Sends a new message from one user to another, including sender and receiver details along with the content of the message.
// @Accept json
// @Produce json
// @Param req body ctype.MessageRequest true "Request body for sending a new message including sender and receiver user IDs, and the content of the message"
// @Router /api/messages [post]
// @Success 200 {object} resp.Response{} "Confirms that the message has been successfully sent."
// @Failure 400 {object} resp.Response{} "Invalid request body or user IDs"
// @Failure 404 {object} resp.Response{} "Sender or receiver not found"
// @Failure 500 {object} resp.Response{} "Failed to send the message due to internal server error"
func (MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id
	var cr ctype.MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	var sendUser, receiveUser models.UserModel

	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		resp.FailWithMsg("发送人不存在", c)
		return
	}
	err = global.DB.Take(&receiveUser, cr.RevUserID).Error
	if err != nil {
		resp.FailWithMsg("接收人不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  receiveUser.NickName,
		RevUserAvatar:    receiveUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg("消息发送失败", c)
		return
	}
	resp.OkWithMsg("消息发送成功", c)
	return
}
