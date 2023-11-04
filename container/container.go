package container

import (
	"app/internal/program_studi/interfaces/public/restapi"
	"app/internal/program_studi/service"
	restapiUser "app/internal/user/interfaces/public/restapi"
	serviceUser "app/internal/user/service"
	"github.com/gin-gonic/gin"
)

func NewContainer(r *gin.Engine) {
	userApp := serviceUser.NewApplication()
	restapiUser.AddRoute(r, userApp)

	programStudiApp := service.NewApplication()
	restapi.AddRoute(r, programStudiApp)
}
