package sql

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/moniang/chat/lib"
	"time"
)

// 用户表结构体
type User struct {
	ID           int       `gorm:"column:user_id;primary_key"`
	Nick         string    `gorm:"column:user_nick"`
	Pass         string    `gorm:"column:user_pass"`
	User         string    `gorm:"column:user_user"`
	Salt         string    `gorm:"column:user_salt"`
	Token        string    `gorm:"column:user_token"`
	Vip          int       `gorm:"column:user_vip"`
	TokenEndTime int64     `gorm:"column:user_token_end_time"`
	AddTime      time.Time `gorm:"column:user_add_time"`
	Set          Set       `gorm:"ForeignKey:set_user_id"`
}

// 返回表名
func (User) TableName() string {
	return "chat_user"
}

// 获取用户
func GetUser(userWhere *User) (user []User, err error) {
	err = DB.Where(userWhere).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

// 判断用户登录
func CheckUserLogin(user string, pass string) (userList []User, result bool) {
	userInfo := &User{
		User: user,
	}
	userList, _ = GetUser(userInfo)
	if len(userList) == 0 {
		return userList, false
	}

	if makePass(pass, userList[0].Salt) != userList[0].Pass {
		return userList, false
	}
	return userList, true
}

// 添加用户
func AddUser(user string, pass string, nick string) error {
	addUser := &User{}
	addUser.Salt = lib.GetRandomString(8)
	addUser.Pass = makePass(pass, addUser.Salt)
	addUser.User = user
	addUser.Nick = nick
	addUser.AddTime = time.Now().UTC()
	DB.NewRecord(addUser)
	return DB.Create(&addUser).Error
}

// 生成Token
func MakeToken(user *User) (string, error) {
	h := md5.New()
	h.Write([]byte(lib.GetRandomString(20)))
	token := hex.EncodeToString(h.Sum(nil))

	return token, DB.Model(user).Update("user_token", token, "user_token_end_time", time.Now().AddDate(0, 0, 1).Unix()).Error
}

// 检查Token，并获取用户信息
func CheckToken(token string) (User, bool) {
	if len(token) != 32 {
		return User{}, false
	}
	user, err := GetUser(&User{Token: token})
	if len(user) <= 0 || err != nil {
		return User{}, false
	}
	return user[0], true
}

// 生成密码
func makePass(pass string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt + "chat" + pass + salt))
	return hex.EncodeToString(h.Sum(nil))
}
