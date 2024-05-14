package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AshishKothariii/gotalkapi/services"
	"github.com/AshishKothariii/gotalkapi/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct {
	ID       string `json:"_id,omitempty"`
	Username string `json:"username,omitempty"`
}

// UserController handles web requests related to Users
type UserController struct {
        service services.UserService
}

// NewUserController creates a new UserController
func NewUserController(s services.UserService) *UserController {
        return &UserController{service: s}
}

// RegisterUser handles POST /register endpoint
func (uc *UserController) RegisterUser(c *gin.Context) {
        var userInfo struct {
                Username string `json:"username"`
                Email    string `json:"email"`
                Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&userInfo); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
                return
        }

        user, err := uc.service.CreateUser(c, userInfo.Username, userInfo.Email, userInfo.Password)
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
                return
        }

        c.JSON(http.StatusCreated, gin.H{"id": user.ID.Hex()})
}

// Login handles POST /login endpoint
func (uc *UserController) Login(c *gin.Context) {
        var credentials struct {
                Username string `json:"username"`
                Password string `json:"password"`
        }

        if err := c.ShouldBindJSON(&credentials); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
                return
        }

        user, err := uc.service.CheckUserCredentials(c, credentials.Username, credentials.Password)
        if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
                return
        }

        token, err := utils.GenerateJWT(user.ID.Hex(), user.Username)
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
                return
        }
                c.SetCookie("token", token, 3600, "/", os.Getenv("CLIENT_URL"), false, true)
                c.JSON(http.StatusAccepted,gin.H{
                        "username": user.Username,
                        "Email":    user.Email,
                        "LoginStatus": true,
                })
}


func (uc *UserController)  Logout(c *gin.Context) {
        c.SetCookie("token", "", -1, "/", os.Getenv("CLIENT_URL"), false, true)
        c.JSON(http.StatusOK, "ok")
}
func (uc *UserController) GetProfile(c *gin.Context) {
                token, _ := c.Cookie("token")
               ans,_ := utils.ParseJWT(token)
               fmt.Println(ans)
               userid,_ :=primitive.ObjectIDFromHex(ans)
               user,err :=uc.service.GetUserByID(c,userid)
               if err!=nil{
                fmt.Print(err)
               }
          c.JSON(http.StatusOK, gin.H{
                        "username": user.Username,
                        "Email":    user.Email,
                })               


}

func (uc *UserController) GetUserByUserName(c *gin.Context){
                username := c.Param("username")
                user ,_ :=uc.service.GetUserByUserName(c,username)
                c.JSON(http.StatusOK, gin.H{
                         "username": user.Username,
                         "Email":    user.Email,
                })               
}