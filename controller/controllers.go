package controller

import (
	"github.com/aURORA-JC/headerfaker/data"
	"github.com/aURORA-JC/headerfaker/database"
	"github.com/aURORA-JC/headerfaker/model"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"net/http"
	"strings"
)

// IndexHandler Handle index router
func IndexHandler(c *gin.Context) {
	// Get user cookie, if not exist, set a guest cookie
	cookie, cookieError := c.Cookie("hfer-cookie")
	if cookieError != nil {
		cookie = "guest"
		c.SetCookie("hfer-cookie", "guest", 0, "/", "", false, false)
	}

	// Set login status & Get user records
	status := 0
	var tuids []string
	if cookie != "guest" {
		status = 1
		database.DB.Table("records").Select("tuid").Where("uuid = ?", cookie).Scan(&tuids)
	}

	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"status":   status,
		"user":     cookie,
		"testlist": &data.TestData,
		"record":   tuids,
	})
}

// LoginPageHandler Handle the GET request for user login to show the page
func LoginPageHandler(c *gin.Context) {
	// Get user cookie, if not guest, turn to index
	cookie, _ := c.Cookie("hfer-cookie")
	if cookie != "guest" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.HTML(http.StatusOK, "index/user.tmpl", gin.H{
		"method": "login",
		"text":   "Login",
	})
}

// LoginPostHandler Handle the POST request for user login (login API)
func LoginPostHandler(c *gin.Context) {
	// Get user cookie, if not guest, turn to index
	cookie, _ := c.Cookie("hfer-cookie")
	if cookie != "guest" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Get user post form & parse json as struct
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	jsonError := c.BindJSON(&user)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "BadRequest",
		})
		return
	}

	// Authenticated user password
	var password, existUUID string
	database.DB.Table("users").Select("password").Where("name = ?", user.Username).Scan(&password)
	if password == user.Password {
		database.DB.Table("users").Select("uuid").Where("name = ?", user.Username).Scan(&existUUID)
		c.SetCookie("hfer-cookie", existUUID, 86400, "/", "", false, false)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "LoginSuccessful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "LoginFailed",
		})
	}
}

// RegPageHandler Handle the GET request to show the page
func RegPageHandler(c *gin.Context) {
	// Get user cookie, if not guest, turn to index
	cookie, _ := c.Cookie("hfer-cookie")
	if cookie != "guest" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.HTML(http.StatusOK, "index/user.tmpl", gin.H{
		"method": "register",
		"text":   "Register",
	})
}

// RegPostHandler Handle the POST request for user register (Register API)
func RegPostHandler(c *gin.Context) {
	// Get user cookie, if not guest, turn to index
	cookie, _ := c.Cookie("hfer-cookie")
	if cookie != "guest" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Get user post form & parse json as struct
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	jsonError := c.BindJSON(&user)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "BadRequest",
		})
		return
	}

	// Check user submit, make sure the username is unique
	var isExist int
	database.DB.Table("users").Select("COUNT(*)").Where("name = ?", user.Username).Scan(&isExist)
	if isExist != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "UsernameAlreadyExists",
		})
		return
	}

	// Generate a new uuid
	uuidString := uuid.New()
	record := model.User{
		Name:     user.Username,
		Password: user.Password,
		UUID:     uuidString,
	}

	// Write to Database
	result := database.DB.Create(&record)
	if result.Error != nil {
		color.Red("ERROR! Can not write Database!")
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadGateway,
			"message": "ServerError",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "RegisterSuccess",
		})
	}
}

// LogoutHandler Handle the logout request
func LogoutHandler(c *gin.Context) {
	// Set user cookie to guest
	c.SetCookie("hfer-cookie", "guest", 0, "/", "", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// GetResponsesHandler Handle test router
func GetResponsesHandler(c *gin.Context) {
	c.Header("HeaderFaker-Flag", data.FlagData["getResponsesID"].Flag)
	c.HTML(http.StatusOK, "quiz/getresponse.tmpl", nil)
}

// EasyPasswordGETHandler Handle test router
func EasyPasswordGETHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz/easypassword.tmpl", nil)
}

// EasyPasswordPOSTHandler Handle test router
func EasyPasswordPOSTHandler(c *gin.Context) {
	// Get user post form & parse json as struct
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
	// Get key's value from url & check length
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
	// Check request header's referer
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
	// Check User-Agent
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
	// Check X-Forwarded-For
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
	// Get user cookie & auth character
	cookie, cookieError := c.Cookie("hfer-cookie")
	if cookieError != nil || cookie == "guest" {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "NotLoggedIn",
		})
		return
	}

	// Get user post form & parse json as struct
	var check model.Check
	jsonError := c.BindJSON(&check)
	if jsonError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "BadRequest",
		})
		return
	}

	// Read user's records
	flagStruct := data.FlagData[check.CheckTag]
	var count int
	database.DB.Table("records").Select("COUNT(*)").Where("uuid = ? AND tuid = ?", cookie, flagStruct.TUID).Scan(&count)
	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    450,
			"message": "AlreadyDone",
		})
		return
	}

	// Check flag
	if check.CheckFlag == flagStruct.Flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "OK",
		})

		record := model.Record{
			UUID: cookie,
			TUID: flagStruct.TUID,
		}

		result := database.DB.Create(&record)
		if result.Error != nil {
			color.Red("ERROR! Can not write Database!")
		}
		//data.FlagData[check.CheckTag] = util.GenerateFlag(check.CheckTag)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusForbidden,
			"message": "IncorrectFlag",
		})
	}
}
