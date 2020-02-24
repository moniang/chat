package main

import (
	"github.com/jinzhu/gorm"
	"github.com/moniang/chat/lib"
	"github.com/moniang/chat/service"
	"github.com/moniang/chat/sql"
	"net/http"
)

type Register struct {
	User string `json:"user"` // 注册账号
	Pass string `json:"pass"` // 注册密码
	Nick string `json:"nick"` // 用户昵称
}

type UserInfo struct {
	Token     string `json:"token"`      // 用户登录成功后的凭证
	Nick      string `json:"nick"`       // 用户昵称
	FontColor string `json:"font_color"` // 字体颜色
}

// 首页界面
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

// 登录界面及登录函数
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // 这里用来处理传递上来的登录信息
		postUser := r.PostFormValue("user")
		postPass := r.PostFormValue("pass")

		w.Header().Add("content-type", "application/json;charset=utf-8")

		if lib.IsEmpty(postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号或者密码不可为空", nil))
			return
		}

		if !lib.IsAlphaNum(postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号密码只能由字母和数字组成", nil))
			return
		}

		if !lib.Len(6, 20, postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号密码长度为6~20位", nil))
			return
		}

		user, result := sql.CheckUserLogin(postUser, postPass)
		if !result {
			w.Write(lib.MakeReturnJson(502, "账号或者密码错误", nil))
			return
		}
		token, err := sql.MakeToken(&user[0])
		if err != nil {
			w.Write(lib.MakeReturnJson(503, "登录失败", nil))
			return
		}
		set := &sql.Set{} // 获取用户设置
		setErr := sql.DB.Model(user[0]).Related(&set).Error
		if setErr != nil {
			set.FontColor = "#000000"
		}
		userInfo := &UserInfo{
			Token:     token,
			Nick:      user[0].Nick,
			FontColor: set.FontColor,
		}
		w.Write(lib.MakeReturnJson(200, "登录成功", userInfo))
		return
	}
}

// 注册界面及注册函数
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // 这里用来处理传递上来的注册信息
		var registerInfo = Register{
			User: r.PostFormValue("user"),
			Pass: r.PostFormValue("pass"),
			Nick: r.PostFormValue("nick"),
		}
		w.Header().Add("content-type", "application/json;charset=utf-8")
		if lib.IsEmpty(registerInfo.Nick, registerInfo.Pass, registerInfo.User) {
			w.Write(lib.MakeReturnJson(501, "需要填写全部参数", nil))
			return
		}
		if !lib.IsChsAlphaNum(registerInfo.Nick) {
			w.Write(lib.MakeReturnJson(501, "昵称只能由汉字、字母和数字组成", nil))
			return
		}

		if !lib.Len(3, 15, registerInfo.Nick) {
			w.Write(lib.MakeReturnJson(501, "昵称长度为1~5个汉字(15个字符)", nil))
			return
		}

		if !lib.Len(6, 20, registerInfo.Pass, registerInfo.User) {
			w.Write(lib.MakeReturnJson(501, "账号以及密码长度为6~20位", nil))
			return
		}

		if !lib.IsAlphaNum(registerInfo.User, registerInfo.Pass) {
			w.Write(lib.MakeReturnJson(501, "账号密码只能由字母和数字组成", nil))
			return
		}

		user, _ := sql.GetUser(&sql.User{User: registerInfo.User})
		if len(user) != 0 {
			w.Write(lib.MakeReturnJson(502, "账号已存在", registerInfo.User))
			return
		}

		user, _ = sql.GetUser(&sql.User{Nick: registerInfo.Nick})
		if len(user) != 0 {
			w.Write(lib.MakeReturnJson(502, "昵称不可重复", registerInfo.Nick))
			return
		}
		err := sql.AddUser(registerInfo.User, registerInfo.Pass, registerInfo.Nick)
		if err != nil {
			w.Write(lib.MakeReturnJson(503, "注册失败", nil))
			return
		}
		w.Write(lib.MakeReturnJson(200, "注册成功", nil))
	}
}

// 修改昵称
func reviseName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nick := r.PostFormValue("nick")
		w.Header().Add("content-type", "application/json;charset=utf-8")
		if !lib.IsChsAlphaNum(nick) {
			w.Write(lib.MakeReturnJson(501, "昵称只能由汉字、字母和数字组成", nil))
			return
		}
		if !lib.Len(3, 15, nick) {
			w.Write(lib.MakeReturnJson(501, "昵称长度为1~5个汉字(15个字符)", nil))
			return
		}
		token := r.Header.Get("token")
		user, result := sql.CheckToken(token)
		if result == false {
			w.Write(lib.MakeReturnJson(503, "登录失效，请重新登录", nil))
			return
		}
		if user.Nick == nick { // 当前昵称和想修改的昵称一致，回复治愈消息
			w.Write(lib.MakeReturnJson(200, "昵称修改成功", nil))
			return
		}

		nickUser, _ := sql.GetUser(&sql.User{Nick: nick})
		if len(nickUser) != 0 {
			w.Write(lib.MakeReturnJson(502, "当前昵称已存在", nil))
			return
		}
		err := sql.DB.Model(user).Update("user_nick", nick).Error
		if err != nil {
			w.Write(lib.MakeReturnJson(504, "修改昵称失败", nil))
			return
		}
		// 修改目前登录信息中缓存的值
		v, ok := service.SocketList.Load(user.ID)
		if ok {
			client := v.(service.Client)
			client.Name = nick
			service.SocketList.Store(user.ID, client)
		}
		w.Write(lib.MakeReturnJson(200, "昵称修改成功", nil))
		return
	}
}

// 修改字体颜色
func reviseFontColor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		color := r.PostFormValue("fontColor")
		w.Header().Add("content-type", "application/json;charset=utf-8")
		if !lib.IsColorHex(color) {
			w.Write(lib.MakeReturnJson(501, "请输入正确的颜色值", nil))
			return
		}
		token := r.Header.Get("token")
		user, result := sql.CheckToken(token)
		if result == false {
			w.Write(lib.MakeReturnJson(503, "登录失效，请重新登录", nil))
			return
		}

		set := &sql.Set{}
		err := sql.DB.Model(user).Related(&set).Error
		if err == gorm.ErrRecordNotFound {
			set.FontColor = color
			set.UserId = user.ID
			err = sql.DB.Create(&set).Error
		} else {
			set.FontColor = color
			err = sql.DB.Save(&set).Error
		}
		if err != nil {
			w.Write(lib.MakeReturnJson(504, "修改字体颜色失败", nil))
			return
		}
		// 修改目前登录信息中缓存的值
		v, ok := service.SocketList.Load(user.ID)
		if ok {
			client := v.(service.Client)
			client.FontColor = color
			service.SocketList.Store(user.ID, client)
		}
		w.Write(lib.MakeReturnJson(200, "字体颜色修改成功", nil))
		return
	}
}
