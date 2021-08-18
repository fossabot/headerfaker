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

// EasyPasswordGETHandler Handle test router
func EasyPasswordGETHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz/easypassword.tmpl", nil)
}

// EasyPasswordPOSTHandler Handle test router
func EasyPasswordPOSTHandler(c *gin.Context) {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	jsonError := c.BindJSON(&login)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "BadRequest",
		})
		return
	}

	if login.Username == "admin" && login.Password == "admin" {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["easyPasswordID"],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "WrongAuthenticationInformation",
		})
	}

}

// GetHandler Handle test router (GET)
func GetHandler(c *gin.Context) {
	getKey := c.Query("key")
	if len(getKey) >= 128 {
		c.JSON(http.StatusRequestURITooLong, gin.H{
			"code":    http.StatusRequestURITooLong,
			"message": "RequestURITooLong",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "InCorrectKey",
		})
	}
}

// PostHandler Handle test router (POST)
func PostHandler(c *gin.Context) {
	postKey := c.Query("key")

	key := "w8f0bxSqIpb4uO8lPN6BCFUe8XFDPjIOX4XePKYyuxACeHKr9YJ5sgBVRvULboFekvI5ZKIlSem3yRQbYjZlsBBVLEHlO70xm4vhemJYZ6DeIAjfGST4ybncfHTFLI3k"
	if postKey == key {
		c.HTML(http.StatusOK, "quiz/flag.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["getOrPostID"],
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "InCorrectKey",
		})
	}
}

// FakeReferHandler Handle test router
func FakeReferHandler(c *gin.Context) {
	if c.Request.Referer() == "https://cust.team" || c.Request.Referer() == "https://cust.team/" {
		c.HTML(http.StatusOK, "quiz/flag.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["fakeReferID"],
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "InCorrectRefer",
			"required": gin.H{
				"host": "cust.team",
				"tls":  "necessary",
			},
		})
	}
}

// FakeAgentHandler Handle test router
func FakeAgentHandler(c *gin.Context) {
	if strings.Contains(c.Request.UserAgent(), "iPhone") && strings.Contains(c.Request.UserAgent(), "AppleWebKit") && strings.Contains(c.Request.UserAgent(), "NetType 2G/3G/4G/5G") {
		c.HTML(http.StatusOK, "quiz/flag.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["fakeAgentID"],
		})
	} else if strings.Contains(c.Request.UserAgent(), "iPhone") == false {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "ForbiddenDevice",
		})
	} else if strings.Contains(c.Request.UserAgent(), "AppleWebKit") == false {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "InCorrectUserAgent",
		})
	} else if strings.Contains(c.Request.UserAgent(), "NetType 2G/3G/4G/5G") == false {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "InCorrectUserAgent",
			"required": gin.H{
				"key":       "NetType",
				"value":     "2G/3G/4G/5G",
				"delimiter": " ",
			},
		})
	}
}

func FakeIPHandler(c *gin.Context) {
	if strings.Contains(c.Request.Header.Get("X-Forwarded-For"), "1.1.1.1") {
		c.HTML(http.StatusOK, "quiz/flag.tmpl", gin.H{
			"code":    http.StatusOK,
			"message": data.FlagData["fakeIPID"],
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "ForbiddenSource",
			"policy": gin.H{
				"mode":   "whitelist",
				"client": "Cloudflare DNS Center",
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
		return
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
