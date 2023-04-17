package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luandnh/go-notify/common/model"
	"github.com/luandnh/go-notify/common/response"
	"github.com/luandnh/go-notify/common/util"
	authMdw "github.com/luandnh/go-notify/middleware/auth"
	"github.com/luandnh/go-notify/service"
)

type UserAPI struct {
	sv service.IUser
}

func NewUserAPI(r *gin.Engine, appSv service.IUser) {
	handler := &UserAPI{
		sv: appSv,
	}
	Group := r.Group("v1/user")
	{
		Group.GET("", authMdw.AuthMiddleware(), handler.GetUsers)
	}
}

func (h *UserAPI) GetUsers(c *gin.Context) {
	user, ok := authMdw.GetUser(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	} else if user.Level != service.ADMIN {
		c.JSON(response.Forbidden())
		return
	}
	filter := model.GeneralFilter{
		PageSize: util.ParsePageSize(c.Query("page_size")),
		Page:     util.ParsePage(c.Query("page")),
	}
	_ = filter
	// code, result := h.sv.GetUsers(c, filter)
	// c.JSON(code, result)
}

func (h *UserAPI) PostUser(c *gin.Context) {
	user, ok := authMdw.GetUser(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	}
	body := map[string]any{}
	if c.BindJSON(&body) != nil {
		c.JSON(response.BadRequest())
		return
	}
	_ = user
	// code, result := h.sv.PostUser(c, body)
	// c.JSON(code, result)
}

//TODO: implement PUT, DELETE, GET
