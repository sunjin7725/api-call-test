package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/sunjin7725/api-call-test/docs" // 프로젝트 내부의 docs 패키지를 꼭 참조할것
	"github.com/sunjin7725/api-call-test/internal"
)

//	@title			Food DB API
//	@version		1.0
//	@description	Sample API for food db.

//	@contact.name	Seon Jin Kim
//	@contact.url	https://sunjin7725.github.io
//	@contact.email	sunjin7725@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.basic	BasicAuth
func main() {
	/* Custom Recovery behavior
	r := gin.New()
	r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	*/
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r = internal.SetupRouter(r)
	r.Run("0.0.0.0:8080")
}
