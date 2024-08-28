package article_api

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"myServer/global"
	"myServer/models"
	"myServer/models/ctype"
	"myServer/utils/encapsulation/resp"
	"myServer/utils/jwt"
	"strings"
	"time"
)

// ArticleCreateView
// @Tags Article Management
// @Summary Create a new article
// @Description Adds a new article entry to the ES. The process includes content sanitization to remove potentially harmful scripts, auto-generating an abstract if not provided, and assigning a random banner if none specified.
// @Accept json
// @Produce json
// @Param req body ctype.ArticleRequest true "Request body for creating a new article including title, content, and optional banner ID"
// @Router /api/articles [post]
// @Success 200 {object} resp.Response{} "Confirms that the article has been successfully added to the database."
// @Failure 400 {object} resp.Response{} "Invalid request body or failure in content sanitization"
// @Failure 404 {object} resp.Response{} "Banner or user not found"
// @Failure 409 {object} resp.Response{} "Duplicate article entries found"
// @Failure 500 {object} resp.Response{} "Failed to create article due to internal server error"
func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ctype.ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		resp.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	// 校验content  xss

	// 处理content
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	// 是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//fmt.Println(doc.Text())
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstract == "" {
		// 汉字的截取不一样
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传banner_id,后台就随机去选择一张
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			resp.FailWithMsg("没有banner数据", c)
			return
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIDList[rand.Intn(len(bannerIDList))]
	}

	// 查banner_id下的ubanner_rl
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		resp.FailWithMsg("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	//应该去判断文章标题是否存在
	if article.ISExistData() {
		resp.FailWithMsg("文章已存在", c)
		return
	}

	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		resp.FailWithMsg(err.Error(), c)
		return
	}
	resp.OkWithMsg("文章发布成功", c)

}
