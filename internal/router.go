package internal

import (
	"github.com/gin-gonic/gin"

	"github.com/sunjin7725/api-call-test/internal/food_db"
)

func SetupRouter(rtx *gin.Engine) *gin.Engine {
	main_group := rtx.Group("")
	{
		main_group.GET("foodData", food_db.FoodDataEndpoint)
	}
	return rtx
}
