package ctype

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"Please enter a title" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"Please enter a menu path" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      Array       `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                                     // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                                         // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"Please enter the menu number" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                                               // 具体图片的顺序
}

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}
