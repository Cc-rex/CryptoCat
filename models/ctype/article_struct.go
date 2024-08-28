package ctype

type ArticleRequest struct {
	Title    string `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string `json:"abstract"`                                // 文章简介
	Content  string `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string `json:"category"`                                // 文章分类
	Source   string `json:"source"`                                  // 文章来源
	Link     string `json:"link"`                                    // 原文链接
	BannerID uint   `json:"banner_id"`                               // 文章封面id
	Tags     Array  `json:"tags"`                                    // 文章标签
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`     // 文章标题
	Abstract string   `json:"abstract"`  // 文章简介
	Content  string   `json:"content"`   // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
	ID       string   `json:"id"`
}

type DeleteIDList struct {
	IDList []string `json:"id_list"`
}

type ArticleSearchRequest struct {
	PageInfo
	Tag string `json:"tag" form:"tag"`
}
