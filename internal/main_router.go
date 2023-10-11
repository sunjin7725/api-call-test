package internal

import (
	"github.com/gin-gonic/gin"

	"github.com/sunjin7725/api-call-test/internal/food_db"
)

func SetupRouter() *gin.Engine {
	rtx := gin.Default()
	rtx = food_db.FoodDataRouter(rtx)
	return rtx
}
