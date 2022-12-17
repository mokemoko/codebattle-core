package openapi

import (
	"github.com/mokemoko/codebattle-core/server/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	userIdKey = "userId"
)

var authMiddleware *jwt.GinJWTMiddleware

func setup() {
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "github", nil
	}
	goth.UseProviders(github.New(
		os.Getenv("GITHUB_CLIENT_KEY"),
		os.Getenv("GITHUB_CLIENT_SECRET"),
		os.Getenv("GITHUB_CALLBACK_URL"),
		"repo",
	))
	_authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "cdbt",
		Key:        []byte(os.Getenv("SESSION_SECRET")),
		SendCookie: true,
		//SecureCookie: true,
		CookieHTTPOnly: true,
		Timeout:        time.Hour * 24 * 30, // one month
		TokenLookup:    "cookie:jwt",
		IdentityKey:    userIdKey,
		IdentityHandler: func(c *gin.Context) interface{} {
			return jwt.ExtractClaims(c)[userIdKey]
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if userId, ok := data.(string); ok {
				return jwt.MapClaims{
					userIdKey: userId,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			if userId, ok := c.Get(userIdKey); ok {
				return userId, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.Redirect(http.StatusMovedPermanently, os.Getenv("FRONT_DOMAIN"))
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	authMiddleware = _authMiddleware
}

func addRoutes(router *gin.Engine) {
	router.GET("/login", func(c *gin.Context) {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})
	router.GET("/callback", func(c *gin.Context) {
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		record := models.User{
			ID:    user.UserID,
			Name:  user.NickName,
			Icon:  user.AvatarURL,
			Token: null.StringFrom(user.AccessToken),
		}
		if err := record.UpsertG(c, true, []string{"id"}, boil.Infer(), boil.Infer()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set(userIdKey, record.ID)
		authMiddleware.LoginHandler(c)
	})
}

func customRouter(baseRouter *gin.Engine) (*gin.RouterGroup, *gin.RouterGroup) {
	setup()

	baseRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONT_DOMAIN")},
		AllowCredentials: true,
		AllowMethods:     []string{"PUT"},
		AllowHeaders:     []string{"content-type"},
	}))

	unauthorizedRouter := baseRouter.Group("")
	authorizedRouter := baseRouter.Group("")
	authorizedRouter.Use(authMiddleware.MiddlewareFunc())

	addRoutes(baseRouter)

	return unauthorizedRouter, authorizedRouter
}
