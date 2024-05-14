package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AshishKothariii/gotalkapi/controller"
	"github.com/AshishKothariii/gotalkapi/db"
	"github.com/AshishKothariii/gotalkapi/middleware"
	"github.com/AshishKothariii/gotalkapi/repository"
	"github.com/AshishKothariii/gotalkapi/services"
	"github.com/AshishKothariii/gotalkapi/websockets"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
        err := godotenv.Load(".env")
        if err != nil {
                fmt.Print(err)
        return}
        // Connect to MongoDB
        client,err := db.Init()
		if err!=nil{
			fmt.Println("db connection failed")
		}
		database :=client.Database(os.Getenv("DB_NAME"))
        
fmt.Print("till user repo up")
        // Initialize repositories
        userRepo := repository.NewUserRepository(database)
fmt.Print("after user repo before service")
        // Initialize services
        userService := services.NewUserService(userRepo)

        // Initialize controllers
        userController := controller.NewUserController(userService)

        // Set up the router
        router := gin.Default()
        router.Use(gin.Logger())

         router.Use(middleware.CORSMiddleware())

        router.GET("/test", func(c *gin.Context) {
                c.JSON(http.StatusOK, "test ok")
        })

        router.POST("/register", userController.RegisterUser)
        router.POST("/login", userController.Login)
        router.POST("/logout", userController.Logout)
        router.GET("/profile",userController.GetProfile)
        router.GET("/users/:username",userController.GetUserByUserName)
         hub := websockets.NewHub()
    go hub.Run()
        webSocketController :=websockets.NewWebSocketController(userService,hub)

        router.GET("/ws",webSocketController.HandleWebSocket)

	// Periodically broadcast online users
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        if err := router.Run(":" + port); err != nil {
                log.Fatal(err)
        }
}