package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @Summary Swagger 설정 테스트...
// @Description Swagger 설정 되었는지, API통신 가능한지 Test중
// @name testGet
// @Router /testGet [get]
// @Success 200 {string} {string} ok, 잘 됩니다.
// @Failure 400
func ControllerTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "잘 됩니다.",
	})
}
