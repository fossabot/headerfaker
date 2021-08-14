package controller

import (
	"github.com/aURORA-JC/headerfaker/data"
	"github.com/aURORA-JC/headerfaker/model"
	"github.com/aURORA-JC/headerfaker/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// IndexHandler Handle index router
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"testlist": &data.TestData,
	})
}

// GetResponsesHandler Handle test router
func GetResponsesHandler(c *gin.Context) {
	c.Header("HeaderFaker-Flag", data.FlagData["getResponsesID"])
	c.HTML(http.StatusOK, "quiz/getresponse.tmpl", nil)
}

// EasyPasswordHandler Handle test router
func EasyPasswordHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "TestBuilding",
		"title":   "EasyPassword",
	})
}

// GetHandler Handle test router (GET)
func GetHandler(c *gin.Context) {
	key := c.Query("key")
	if len(key) >= 128 {
		c.JSON(http.StatusRequestURITooLong, gin.H{
			"code":    http.StatusRequestURITooLong,
			"message": "RequestURITooLong",
			"title":   "GET / POST",
		})
	} else {
		if key == "w8f0bxSqIpb4uO8lPN6BCFUe8XFDPjIOX4XePKYyuxACeHKr9YJ5sgBVRvULboFekvI5ZKIlSem3yRQbYjZlsBBVLEHlO70xm4vhemJYZ6DeIAjfGST4ybncfHTFLI3k" {
			c.HTML(http.StatusOK, "quiz/quiz.tmpl", gin.H{
				"code":    http.StatusOK,
				"message": data.FlagData["getOrPostID"],
				"title":   "GET / POST",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "InCorrectKey",
				"title":   "GET / POST",
			})
		}
	}
}

// PostHandler Handle test router (POST)
func PostHandler(c *gin.Context) {
	key := c.Query("key")
	if key == "w8f0bxSqIpb4uO8lPN6BCFUe8XFDPjIOX4XePKYyuxACeHKr9YJ5sgBVRvULboFekvI5ZKIlSem3yRQbYjZlsBBVLEHlO70xm4vhemJYZ6DeIAjfGST4ybncfHTFLI3k" {
		c.HTML(http.StatusOK, "quiz/quiz.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["getOrPostID"],
			"title":   "GET / POST",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "InCorrectKey",
			"title":   "GET / POST",
		})
	}
}

// FakeReferHandler Handle test router
func FakeReferHandler(c *gin.Context) {
	if c.Request.Referer() == "https://cust.team" || c.Request.Referer() == "https://cust.team/" {
		c.HTML(http.StatusOK, "quiz/quiz.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["fakeReferID"],
			"title":   "Fake Refer",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "InCorrectRefer",
			"title":   "Fake Refer",
			"required": gin.H{
				"host": "cust.team",
				"tls":  "necessary",
			},
		})
	}
}

// FakeAgentHandler Handle test router
func FakeAgentHandler(c *gin.Context) {
	if strings.Contains(c.Request.UserAgent(), "iPhone") && strings.Contains(c.Request.UserAgent(), "NetType 2G/3G/4G/5G") {
		c.HTML(http.StatusOK, "quiz/quiz.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["fakeAgentID"],
			"title":   "Fake Agent",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "InCorrectDevice",
			"title":   "Fake Agent",
			"required": gin.H{
				"key":       "NetType",
				"value":     "2G/3G/4G/5G",
				"delimiter": " ",
			},
		})
	}
}

// CheckHandler Handle check routers
func CheckHandler(c *gin.Context) {
	var check model.Check
	jsonError := c.BindJSON(&check)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "BadRequest",
		})
	}

	if check.CheckFlag == data.FlagData[check.CheckTag] {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
		})
		data.FlagData[check.CheckTag] = util.GenerateFlag(check.CheckTag)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "IncorrectFlag",
		})
	}
}
