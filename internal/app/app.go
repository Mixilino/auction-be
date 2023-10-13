package app

import (
	"auction-be/internal/config"
	"auction-be/internal/databases"
	"auction-be/internal/handlers/auth_handler"
	"auction-be/internal/repos/user_repo"
	"auction-be/internal/services/auth_service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type App struct {
	authHandler auth_handler.Auth
	router      *gin.Engine
}

func NewApp() *App {
	config.Load()
	router := gin.Default()
	db := databases.GetConnection()
	userRepo := user_repo.NewUserSqlImpl(db)
	userService := auth_service.NewAuthService(db, userRepo)
	userHandler := auth_handler.NewAuthHandler(userService)

	return &App{
		authHandler: userHandler,
		router:      router,
	}
}

func (a *App) initializeRoutes() {
	a.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiGroup := a.router.Group("/api")
	{
		apiGroup.POST("/signup", a.authHandler.Signup)
		apiGroup.POST("/login", a.authHandler.Login)
	}

	realTimeGroup := a.router.Group("/realtime")
	{
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		realTimeGroup.GET("/", func(c *gin.Context) {
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				return
			}
			defer conn.Close()
			for {
				conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
				time.Sleep(time.Second)
			}
		})
	}
}

func (a *App) StartServer(port string) error {
	a.initializeRoutes()
	server := &http.Server{
		Handler:        a.router,
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
