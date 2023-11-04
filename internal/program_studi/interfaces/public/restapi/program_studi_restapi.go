package restapi

import (
	"app/internal/common/http_resp"
	"app/internal/program_studi/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProgramStudiRestApi struct {
	app usecase.Application
}

func AddRoute(r *gin.Engine, app usecase.Application) {
	resource := ProgramStudiRestApi{app: app}
	r.GET("program-studi", resource.List)
}

func (u ProgramStudiRestApi) List(ctx *gin.Context) {

	res, err := u.app.List.Handle(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, http_resp.ErrRequest(http.StatusBadRequest, err))
		return
	}

	ctx.JSON(http.StatusOK, http_resp.SuccRequest(map[string]interface{}{
		"program_studi": res,
	}))
}
