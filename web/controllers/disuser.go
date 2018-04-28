package contorllers

import (
	"encoding/json"
	"log"
	"my-go-web/web/models"
	"strconv"
	"sync"
)

type UserController struct {
	BaseController
	mutex sync.Mutex
}
// @Success 200 {object}
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (u *UserController) Post() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	var user *models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	flag, err := models.AddUser(user)
	if err != nil {
		u.setdata(flag, nil)
		log.Print(err.Error())
		return
	}
	u.setdata(flag, nil)
}


// @Success 200 {object}
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /login [post]
func (u *UserController) Login() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	var user *models.User
	var err error
	user.UserId, err = strconv.Atoi(u.GetString("userid"))
	if err != nil {
		log.Print(err.Error())
		u.setdata(false, nil)
		return
	}
	user.Passwd = u.GetString("passwd")
	user, err = models.QueryByUserID(user)
	if err != nil {
		u.setdata(false, nil)
		log.Println(err.Error())
		return
	}
	u.setdata(true, user)
}
