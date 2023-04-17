package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luandnh/go-notify/common/log"
	"github.com/luandnh/go-notify/service"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

var cacheObj libcache.Cache
var strategy union.Union
var tokenStrategy auth.Strategy

type LocalAuthMiddleware struct {
}

type GoAuthInfo interface {
	auth.Info
}

func NewLocalAuthMiddleware() IAuthMiddleware {
	return &LocalAuthMiddleware{}
}

func SetupGoGuardian() {
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 10)
	basicStrategy := basic.NewCached(validateBasicAuth, cacheObj)
	tokenStrategy = token.New(validateTokenAuth, cacheObj)
	appStrategy := token.New(validateAppTokenAuth, cacheObj, token.SetParser(token.XHeaderParser("X-APP-TOKEN")))
	strategy = union.New(tokenStrategy, appStrategy, basicStrategy)
}

func (auth *LocalAuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, user, err := strategy.AuthenticateRequest(c.Request)
		if err != nil {
			log.Error("invalid credentials")
			c.JSON(
				http.StatusUnauthorized,
				map[string]interface{}{
					"error": http.StatusText(http.StatusUnauthorized),
				},
			)
			c.Abort()
			return
		}
		c.Set("user", user)

	}
}

var ClientSvr *service.Client

func validateTokenAuth(ctx context.Context, r *http.Request, tokenString string) (auth.Info, time.Time, error) {
	if tokenString == SECRET_TOKEN {
		id := "2273f762-7ae6-4a0e-a09d-6d5a3c961a50"
		name := SUPERADMIN
		level := SUPERADMIN
		user := NewGoAuthUser(id, name, level, nil, "", "token")
		return user, time.Now(), nil
	}
	user, err := service.UserSvr.FindByUserToken(ctx, tokenString)
	if err != nil {
		log.Error(err)
		return nil, time.Time{}, errors.New("invalid credentials")
	} else if user == nil {
		log.Error("basic auth not found username")
		return nil, time.Time{}, errors.New("invalid credentials")
	}
	return NewGoAuthUser(user.UserId, user.Username, user.Level, nil, user.ApplicationId, "token"), time.Now(), nil
}

func validateBasicAuth(ctx context.Context, r *http.Request, username, password string) (auth.Info, error) {
	user, err := service.UserSvr.FindByUsernameAndPassword(ctx, username, password)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid credentials")
	} else if user == nil {
		log.Error("basic auth not found username")
		return nil, errors.New("invalid credentials")
	}
	return NewGoAuthUser(user.UserId, user.Username, user.Level, nil, user.ApplicationId, "basic"), nil
}

func validateAppTokenAuth(ctx context.Context, r *http.Request, tokenString string) (auth.Info, time.Time, error) {
	application, err := service.ApplicationSvr.FindByApplicationToken(ctx, tokenString)
	if err != nil {
		log.Error(err)
		return nil, time.Time{}, errors.New("invalid credentials")
	} else if application == nil {
		log.Error("app auth not found token")
		return nil, time.Time{}, errors.New("invalid credentials")
	}
	return NewGoAuthUser("", "", "", nil, application.ApplicationId, "app-token"), time.Now(), nil
}
