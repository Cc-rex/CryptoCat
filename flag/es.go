package flag

import "myServer/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
