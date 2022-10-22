package model

type Account struct {
	Uid      int    `json:"uid"`      // 标识
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
