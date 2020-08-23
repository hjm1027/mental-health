package router

import (
	"net/http"

	"github.com/mental-health/handler/course"
	"github.com/mental-health/handler/hole"
	"github.com/mental-health/handler/message"
	"github.com/mental-health/handler/mood"
	"github.com/mental-health/handler/poster"
	"github.com/mental-health/handler/reserve"
	"github.com/mental-health/handler/sd"
	"github.com/mental-health/handler/search"
	"github.com/mental-health/handler/upload"
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
		u.GET("/info/:id/", user.GetInfoBySid)
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
		Hole.POST("/comment/:id/", hole.NewParentComment)
		Hole.POST("/comment/:id/reply/", hole.Reply)
		Hole.PUT("/:id/comment/like/", hole.CommentLike)
		Hole.GET("/comments/:id/", hole.GetComments)
	}

	//Course路由组
	Course := g.Group("/api/v1/course/")
	Course.Use(middleware.AuthMiddleware())
	{
		Course.GET("/info/:id/", course.GetInfo)
		Course.PUT("/like/:id/", course.LikeCourse)
		Course.PUT("/favorite/:id/", course.FavoriteCourse)
		Course.GET("/collection/like/", course.GetLikeCollection)
		Course.GET("/collection/favorite/", course.GetFavoriteCollection)
	}

	//Search路由组
	Search := g.Group("/api/v1/search/")
	Search.Use(middleware.AuthMiddleware())
	{
		Search.GET("/course/", search.SearchCourse)
		Search.GET("/hole/", search.SearchHole)
	}

	//Message路由组
	Message := g.Group("/api/v1/message/")
	Message.Use(middleware.AuthMiddleware())
	{
		Message.GET("/all/", message.Get)
		Message.GET("/count/", message.Count)
		Message.GET("/read/", message.ReadAll)
	}

	//Upload路由组
	Upload := g.Group("/api/v1/upload/")
	Upload.Use(middleware.AuthMiddleware())
	{
		Upload.POST("/image/", upload.Image)
		Upload.POST("/video/", upload.Video)
		Upload.POST("/videolink/", course.NewCourse)
	}

	//Poster路由组
	Poster := g.Group("/api/v1/poster/")
	Poster.Use(middleware.AuthMiddleware())
	{
		Poster.GET("/info/", poster.GetPosterInfo)
		Poster.POST("/info/", poster.PostPosterInfo)
	}

	//Reserve路由组
	Reserve := g.Group("/api/v1/reserve/")
	Reserve.Use(middleware.AuthMiddleware())
	{
		Reserve.GET("/query/", reserve.QueryReserve)
		Reserve.GET("/teacherlist/", reserve.TeacherList)
		Reserve.POST("/new/", reserve.Reserve)
	}

	return g
}
