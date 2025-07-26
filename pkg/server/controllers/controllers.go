package controllers

import (
	"net/http"

	"github.com/bandanascripts/tondru/pkg/client/token"
	"github.com/bandanascripts/tondru/pkg/core"
	"github.com/bandanascripts/tondru/pkg/service/middleware"
	"github.com/gin-gonic/gin"
)

func GenerateTokenHandler(c *gin.Context) {

	var payLoad map[string]any

	if err := c.ShouldBindJSON(&payLoad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, _, err := token.GenerateTokens(c.Request.Context(), payLoad)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accesstoken": accessToken})
}

func InspectTokenHandler(c *gin.Context) {

	userClaim, err := middleware.ExtractClaim(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := core.LogUserSearch(c.Request.Context(), userClaim, c.ClientIP()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userclaim": userClaim})
}

func UserHistoryHandler(c *gin.Context) {

	userClaims, err := core.FetchUserSearch(c.Request.Context(), c.ClientIP())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userhistory": userClaims})
}
