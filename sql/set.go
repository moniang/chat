package sql

// 用户表结构体
type Set struct {
	ID        int    `gorm:"column:set_id;primary_key"`
	UserId    int    `gorm:"column:set_user_id"`
	FontColor string `gorm:"column:set_font_color"`
}

// 返回表名
func (Set) TableName() string {
	return "chat_set"
}
