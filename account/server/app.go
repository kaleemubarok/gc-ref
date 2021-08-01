package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kaleemubarok/gc-ref/account/config"
	"github.com/kaleemubarok/gc-ref/account/resource/account"
	"github.com/kaleemubarok/gc-ref/account/usecase/profiledetail"
	"github.com/kaleemubarok/gc-ref/account/usecase/userauth"
	_ "github.com/lib/pq"
	"github.com/twharmon/gouid"
)

//declare var for usecase used
var userAuthUsecase userauth.IUsecase
var profileDetailUsecase profiledetail.IUsecase

//alias for sqlx
var db *sqlx.DB
var dbResource account.IDB

//standart response
type StandardAPIResponse struct {
	Err     string      `json:"err"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	//config load
	configuration := config.New(".env")

	//database connect
	dbConStr := configuration.Get("POSTGRES_CONNECT")

	dbInit, err := sqlx.Connect("postgres", dbConStr)
	if err != nil {
		log.Fatalln(err)
	}
	dbRsc := account.NewDBResource(dbInit)
	dbResource = dbRsc
	db = dbInit

	//new usecaseinit
	userAuthUsecase = userauth.NewUsecase(dbResource, "s!Gn1ngKEy", configuration)
	profileDetailUsecase = profiledetail.NewUsecase(dbResource, configuration)

	//start gin server
	r := gin.Default()
	r.POST("/register", register)
	r.POST("/login", login)

	r.GET("/user/:user_id", validateSession(getUser))
	r.GET("/user/info", validateSession(getUserInfo))

	r.GET("/uid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": gouid.String(8, gouid.LowerCaseAlphaNum),
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi!",
		})
	})
	r.Run(":7070")
}

func register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	confirmPassword := c.Request.FormValue("confirm_password")

	err := userAuthUsecase.Register(username, password, confirmPassword)
	if err != nil {
		c.JSON(400, StandardAPIResponse{
			Err:     err.Error(),
			Message: "Failed",
		})
		return
	}

	c.JSON(201, StandardAPIResponse{
		Message: "Success create new user",
	})
}

func login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	user, err := userAuthUsecase.Login(username, password)
	if err != nil {
		c.JSON(400, StandardAPIResponse{
			Err:     "Failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, StandardAPIResponse{
		Data: user,
	})
}

func getUserInfo(c *gin.Context) {
	userID := c.GetString("uid")

	user, err := profileDetailUsecase.GetUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, StandardAPIResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, StandardAPIResponse{
		Data: user,
	})
}

func getUser(c *gin.Context) {
	userID := c.Param("user_id")

	user, err := profileDetailUsecase.GetUserByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, StandardAPIResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(200, StandardAPIResponse{
		Data: user,
	})
}

func validateSession(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header["X-Access-Token"]

		if len(accessToken) < 1 {
			c.JSON(403, StandardAPIResponse{
				Err: "No access token provided",
			})
			return
		}

		userID, err := userAuthUsecase.ValidateSession(accessToken[0])
		if err != nil {
			c.JSON(400, StandardAPIResponse{
				Err: "Cannot validate session",
			})
			return
		}
		c.Set("uid", userID)
		handlerFunc(c)
	}
}
