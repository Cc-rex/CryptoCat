package models

type KeyModel struct {
	MODEL
	UserID    uint      `gorm:"index;not null" json:"user_id"`        // 创建索引并设置为非空
	PublicKey string    `gorm:"type:text;not null" json:"public_key"` // 公钥数据
	UserModel UserModel `gorm:"foreignKey:UserID"`                    // 根据UserID关联UserModel
}
