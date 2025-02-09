package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Ambil nilai dari .env
	ipAddress := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, ipAddress, dbPort, dbName)
	// fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("Connection to database is good")

	/* CALL REPOSITORY */
	userRepository := user.NewRepository(db) // call func NewRepository dari /user/repository
	campaignRepository := campaign.NewRepository(db)
	/* END CALL REPOSITORY */

	/* CALL SERVICE */
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	campaignService := campaign.NewService(campaignRepository)
	/* END CALL SERVICE */

	/* CALL HANDLER */
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	/* END CALL HANDLER */

	/* ROUTING */
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	// api.POST("/avatars", authMiddleware, userHandler.UploadAvatar) // passing function authMiddleware
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar) // passing nilai kembalian dari function authMiddleware

	api.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run()
	/* END ROUTING */
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// validation
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // middleware ada di tengah, jika ada error/ terkena validasi, proses seharusnya dihentikan
			return
		}

		// Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // middleware ada di tengah, jika ada error/ terkena validasi, proses seharusnya dihentikan
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // middleware ada di tengah, jika ada error/ terkena validasi, proses seharusnya dihentikan
			return
		}

		// ambil id user
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // middleware ada di tengah, jika ada error/ terkena validasi, proses seharusnya dihentikan
			return
		}

		c.Set("currentUser", user)
	}
}
