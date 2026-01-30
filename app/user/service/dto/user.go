package dto

import (
	"go-admin/app/user/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

// UserGetPageReq 列表请求参数
type UserGetPageReq struct {
	dto.Pagination `search:"-"`
	Username       string `form:"username"  search:"type:contains;column:username;table:user" comment:"用户名"`
	NickName       string `form:"nickName"  search:"type:contains;column:nick_name;table:user" comment:"昵称"`
	Phone          string `form:"phone"  search:"type:contains;column:phone;table:user" comment:"手机号"`
	UserOrder
}

type UserOrder struct {
	IdOrder        string `search:"type:order;column:id;table:user" form:"idOrder"`
	UsernameOrder  string `search:"type:order;column:username;table:user" form:"usernameOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:user" form:"createdAtOrder"`
}

func (m *UserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// UserInsertReq 创建请求参数
type UserInsertReq struct {
	Id       int    `json:"-" comment:"编码"`
	Username string `json:"username" comment:"用户名" binding:"required"`
	Password string `json:"password" comment:"密码" binding:"required"`
	NickName string `json:"nickName" comment:"昵称"`
	Phone    string `json:"phone" comment:"手机号"`
	Email    string `json:"email" comment:"邮箱"`
	Status   string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *UserInsertReq) Generate(model *models.User) {
	model.Username = s.Username
	model.Password = s.Password
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.Email = s.Email
	model.Status = s.Status
}

func (s *UserInsertReq) GetId() interface{} {
	return s.Id
}

// UserUpdateReq 更新请求参数
type UserUpdateReq struct {
	Id       int    `uri:"id" comment:"编码"`
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
	NickName string `json:"nickName" comment:"昵称"`
	Phone    string `json:"phone" comment:"手机号"`
	Email    string `json:"email" comment:"邮箱"`
	Status   string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *UserUpdateReq) Generate(model *models.User) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Username = s.Username
	if s.Password != "" {
		model.Password = s.Password
	}
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.Email = s.Email
	model.Status = s.Status
}

func (s *UserUpdateReq) GetId() interface{} {
	return s.Id
}

// UserGetReq 获取请求参数
type UserGetReq struct {
	Id int `uri:"id"`
}

func (s *UserGetReq) GetId() interface{} {
	return s.Id
}

// UserDeleteReq 删除请求参数
type UserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *UserDeleteReq) GetId() interface{} {
	return s.Ids
}
