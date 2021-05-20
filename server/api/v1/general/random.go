package general

import (
	"AynaAPI/server/app"
	"AynaAPI/server/app/e"
	"AynaAPI/utils/vrand"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
)

func GetRandomNumber(context *gin.Context) {
	appG := app.AppGin{C: context}
	start := appG.GetIntQueryWithDefault("start", 0)
	end := appG.GetIntQueryWithDefault("end", 100)
	appG.MakeResponse(http.StatusOK, e.API_OK, map[string]int{
		"value": rand.Intn(end+1) + start,
	})
}

func GetRandomSeparation(context *gin.Context) {
	appG := app.AppGin{C: context}
	k := appG.GetIntQueryWithDefault("k", 2)
	if k <= 0 {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_INVALID_PARAMETER, "k can be less or equal to zero")
		return
	}
	membersS, b := appG.C.GetQuery("names")
	if !b {
		appG.MakeResponse(http.StatusBadRequest, e.API_ERROR_REQUIRE_PARAMETER, "need names")
		return
	}
	members := strings.Split(membersS, ",")
	vrand.ShuffleSlice(members)
	teams := make([][]string, k, k)
	var j int = 0
	for _, val := range members {
		if j >= k {
			j = 0
		}
		teams[j] = append(teams[j], val)
		j++
	}
	appG.MakeResponse(http.StatusOK, e.API_OK, teams)
	return
}
