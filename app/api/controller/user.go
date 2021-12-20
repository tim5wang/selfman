package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tim5wang/selfman/common/web"
	"github.com/tim5wang/selfman/model"
	"github.com/tim5wang/selfman/service/userservice"
)

type UserModule struct {
	userService *userservice.UserService
}

func NewUserModule(userService *userservice.UserService) web.Module {
	return &UserModule{
		userService: userService,
	}
}

func (m *UserModule) Init(r web.Router) {
	g := r.Group("api/user")
	{
		g.GET("/:username", m.GetUser)
		g.POST("", m.CreateUser)
	}
}

type GetUserReq struct {
	UserName string `uri:"username"`
}
type GetUserRes struct {
	UserName    string `json:"username"`
	NickName    string `json:"nickname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avator      string `json:"avator"`
	Description string `json:"description"`
	UserType    int    `json:"user_type"`
}
type CreateUserReq struct {
	NickName string `json:"nickname" form:"nickname"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (m *UserModule) GetUser(ctx *gin.Context, req *GetUserReq) {
	err, user := m.userService.GetUser(req.UserName)
	if err != nil {
		web.Error(ctx, err)
		return
	}
	res := &GetUserRes{
		UserName:    user.UserName,
		NickName:    user.NickName,
		Email:       user.Email,
		Phone:       user.Phone,
		Avator:      user.Avator,
		Description: user.Description,
		UserType:    user.UserType,
	}
	web.Success(ctx, res)
}

func (m *UserModule) CreateUser(ctx *gin.Context, req *CreateUserReq) {
	user := &model.User{}
	user.UserName = req.UserName
	user.Password = req.Password
	user.NickName = req.NickName
	err := m.userService.CreateUser(user)
	web.GeneralResponse(ctx, err)
}
