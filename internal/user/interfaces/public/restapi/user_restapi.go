package restapi

import (
	"app/internal/common/http_resp"
	"app/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRestapi struct {
	app usecase.Application
}

func AddRoute(r *gin.Engine, app usecase.Application) {
	resource := UserRestapi{app: app}
	r.POST("auth/login", resource.Login)
}

func (u UserRestapi) Login(ctx *gin.Context) {
	var loginReq usecase.LoginDTO
	err := ctx.Bind(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, http_resp.ErrRequest(http.StatusBadRequest, err))
		return
	}

	res, err := u.app.Login.Handle(ctx, loginReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, http_resp.ErrRequest(http.StatusBadRequest, err))
		return
	}

	ctx.JSON(http.StatusOK, http_resp.SuccRequest(map[string]interface{}{
		"user": res,
	}))
}
