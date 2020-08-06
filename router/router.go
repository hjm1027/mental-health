package router

import (
	"net/http"

	"github.com/mental-health/handler/hole"
	"github.com/mental-health/handler/mood"
	"github.com/mental-health/handler/sd"
	"github.com/mental-health/handler/user"
	"github.com/mental-health/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//用户认证和登录
	g.POST("/api/v1/login", user.Login)

	//服务器健康检查
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// User路由组
	u := g.Group("/api/v1/user/")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/info/", user.GetInfo)
		u.POST("/info/", user.PostInfo)
	}

	// Mood路由组
	Mood := g.Group("/api/v1/mood/")
	Mood.Use(middleware.AuthMiddleware())
	{
		Mood.GET("/score/", mood.GetMoodScore)
		Mood.GET("/note/", mood.GetMoodNote)
		Mood.POST("/new/", mood.NewMood)
	}

	//Hole路由组
	Hole := g.Group("/api/v1/hole/")
	Hole.Use(middleware.AuthMiddleware())
	{
		Hole.GET("/list/", hole.GetHoleList)
		Hole.GET("/info/:id/", hole.GetHoleInfo)
		Hole.POST("/new/", hole.New)
		Hole.PUT("/:id/like/", hole.LikeHole)
		Hole.PUT("/:id/favorite/", hole.FavoriteHole)
		Hole.GET("/collection/favorite/", hole.GetFavoriteCollection)
	}

	return g
}
