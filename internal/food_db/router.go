package food_db

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FoodDataRouter(rtx *gin.Engine) *gin.Engine {
	rtx.GET("/foodData", func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		if err != nil {
			panic(err)
		}
		per_page, err := strconv.Atoi(ctx.DefaultQuery("per_page", "5"))
		if err != nil {
			panic(err)
		}
		search_keyword := ctx.Query("search_keyword")

		food_data := GetData(page, per_page, search_keyword)
		ctx.String(http.StatusOK, food_data.toJson())
	})
	return rtx
}
