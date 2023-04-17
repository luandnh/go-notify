package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/v2/auth"
)

const (
	SUPERADMIN = "superadmin"
	ADMIN      = "admin"
	USER       = "user"
	CLIENT     = "client"
)

var SECRET_TOKEN = "THIS_IS_A_SECRET_TOKEN_!@#"

type IAuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}

var AuthMdw IAuthMiddleware

func AuthMiddleware() gin.HandlerFunc {
	return AuthMdw.AuthMiddleware()
}

func GetUser(c *gin.Context) (*GoAuthUser, bool) {
	tmp, isExist := c.Get("user")
	if isExist {
		user, ok := tmp.(*GoAuthUser)
		return user, ok
	} else {
		return nil, false
	}
}

func GetUserId(c *gin.Context) (string, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.Id, true
	}
}

func GetUserLevel(c *gin.Context) (string, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.Level, true
	}
}

func GetUserName(c *gin.Context) (string, bool) {
	user, ok := GetUser(c)
	if !ok {
		return "", false
	} else {
		return user.Name, true
	}
}

type GoAuthUser struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Level         string   `json:"level"`
	Scopes        []string `json:"scopes"`
	Extensions    auth.Extensions
	Groups        []string `json:"groups"`
	ApplicationId string   `json:"application_id"`
	AuthType      string   `json:"auth_type"`
}

func NewGoAuthUser(id, name string, level string, scopes []string, applicationId, authType string) GoAuthInfo {
	user := &GoAuthUser{
		Level:         level,
		Scopes:        scopes,
		ApplicationId: applicationId,
		AuthType:      authType,
	}
	user.Id = id
	return user
}

func (d *GoAuthUser) GetID() string {
	return d.Id
}

func (d *GoAuthUser) SetID(id string) {
	d.Id = id
}

func (d *GoAuthUser) GetUserName() string {
	return d.Name
}

func (d *GoAuthUser) SetUserName(name string) {
	d.Name = name
}

func (d *GoAuthUser) GetExtensions() auth.Extensions {
	if d.Extensions == nil {
		d.Extensions = auth.Extensions{}
	}
	return d.Extensions
}

func (d *GoAuthUser) SetExtensions(exts auth.Extensions) {
	d.Extensions = exts
}

func (a *GoAuthUser) SetLevel(level string) {
	a.Level = level
}

func (a *GoAuthUser) GetLevel() string {
	return a.Level
}

func (a *GoAuthUser) SetScopes(scopes []string) {
	a.Scopes = scopes
}

func (a *GoAuthUser) GetScopes() []string {
	return a.Scopes
}

func (d *GoAuthUser) GetGroups() []string {
	return d.Groups
}

func (d *GoAuthUser) SetGroups(groups []string) {
	d.Groups = groups
}
func (d *GoAuthUser) GetApplicationId() string {
	return d.ApplicationId
}

func (d *GoAuthUser) SetApplicationId(applicationId string) {
	d.ApplicationId = applicationId
}
func (d *GoAuthUser) GetAuthType() string {
	return d.AuthType
}

func (d *GoAuthUser) SetAuthType(authType string) {
	d.ApplicationId = authType
}
