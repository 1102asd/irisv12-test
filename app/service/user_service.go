package service

import (
	"irisv12-test/app/middleware"
	"irisv12-test/app/models"
	"irisv12-test/app/repo"
	"irisv12-test/app/utils"
)

type UserService interface {
	Login(m map[string]string) (result models.Result)
	Save(user models.User) (result models.Result)
}
type userServices struct {
	userRepo repo.UserRepository
}

func NewUserServices(repo repo.UserRepository) UserService {
	return &userServices{userRepo: repo}
}

/*
登录
*/
func (u *userServices) Login(m map[string]string) (result models.Result) {
	if m["username"] == "" {
		result.Code = -1
		result.Msg = "请输入用户名！"
		return
	}
	if m["password"] == "" {
		result.Code = -1
		result.Msg = "请输入密码！"
		return
	}
	user := u.userRepo.GetUserByUserNameAndPwd(m["username"], utils.GetMD5String(m["password"]))
	if user.ID == 0 {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
		return
	}
	user.Token = middleware.GenerateToken(user)
	result.Code = 0
	result.Data = user
	result.Msg = "登录成功"
	return result
}

/*
保存
*/
func (u *userServices) Save(user models.User) (result models.Result) {
	//添加
	if user.ID == 0 {
		agen := u.userRepo.GetUserByUsername(user.Username)
		if agen.ID != 0 {
			result.Msg = "登录名重复,保存失败"
			return
		}
	}
	code, p := u.userRepo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 0
	result.Data = p
	return
}
