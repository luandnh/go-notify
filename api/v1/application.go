package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luandnh/go-notify/common/model"
	"github.com/luandnh/go-notify/common/response"
	"github.com/luandnh/go-notify/common/util"
	authMdw "github.com/luandnh/go-notify/middleware/auth"
	"github.com/luandnh/go-notify/service"
)

type ApplicationAPI struct {
	sv service.IApplication
}

func NewApplicationAPI(r *gin.Engine, appSv service.IApplication) {
	handler := &ApplicationAPI{
		sv: appSv,
	}
	Group := r.Group("v1/application")
	{
		Group.GET("", authMdw.AuthMiddleware(), handler.GetApplications)
	}
}

func (h *ApplicationAPI) GetApplications(c *gin.Context) {
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
	code, result := h.sv.GetApplications(c, filter)
	c.JSON(code, result)
}

func (h *ApplicationAPI) PostApplication(c *gin.Context) {
	user, ok := authMdw.GetUser(c)
	if !ok {
		c.JSON(response.Unauthorized())
		return
	} else if user.Level != service.ADMIN {
		c.JSON(response.Forbidden())
		return
	}
	body := map[string]any{}
	if c.BindJSON(&body) != nil {
		c.JSON(response.BadRequest())
		return
	}
	code, result := h.sv.PostApplication(c, body)
	c.JSON(code, result)
}

//TODO: implement PUT, DELETE, GET
