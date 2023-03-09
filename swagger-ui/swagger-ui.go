package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	default_swagger_url string
	http_port           string
	static_file         string
	warningLogger       = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime)
	errorLogger         = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	infoLogger          = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
)

// //////////////////////////
// Get environment variables
// //////////////////////////
func getDefaultSwaggerFile() string {
	return os.Getenv("SWAGGER_URL")
}

func getHttpPort() string {
	return os.Getenv("HTTP_PORT")
}

func getStaticFile() string {
	return os.Getenv("STATIC_FILE_PATH")
}

// ////////////////////////
// Set default value to variable if empty
// ////////////////////////
func checkHttpPort(port string) string {
	if len(port) != 0 {
		return port
	}
	port = "8080"
	return port
}

func checkDefaultSwaggerUrl(url string) string {
	if len(url) != 0 {
		return url
	}
	url = "https://raw.githubusercontent.com/neuvector/neuvector/main/controller/api/apis.yaml"
	return url
}

func checkStaticFile(path string) string {
	if len(path) != 0 {
		return path
	}
	path = "./static"
	return path
}

// ////////////////////////
// Local api for liveness and readiness
// ////////////////////////

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ////////////////////////
// Logger
// ////////////////////////

func isError(err error) bool {
	return err != nil
}

// ////////////////////////
// Main
// ////////////////////////
func main() {
	// Get environment variables
	default_swagger_url = getDefaultSwaggerFile()
	http_port = getHttpPort()
	static_file = getStaticFile()

	// Set default value to variable if empty
	default_swagger_url = checkDefaultSwaggerUrl(default_swagger_url)
	http_port = checkHttpPort(http_port)
	static_file = checkStaticFile(static_file)

	// Create a Gin router
	router := gin.New()

	// Custom logger: [HTTP] 2023/01/08 18:47:25 | Code: 404 | Method: GET | IP: 127.0.0.1 | Path: /api/test
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[HTTP] %s | Code: %d | Method: %s | IP: %s | Path: %s\n",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.StatusCode,
			param.Method,
			param.ClientIP,
			param.Path,
		)
	}))
	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL(default_swagger_url)))

	// Root "/" redirect to default service route
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Static files
	router.Static("/static", "./static")

	// Local api group endpoints
	api := router.Group("/api")
	{
		api.GET("/ping", ping)
	}

	// Start the server
	srv := &http.Server{
		Addr:         ":" + http_port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	infoLogger.Println("Setup http routers/endpoints were successfully finish")
	infoLogger.Println("Server start at port: " + http_port)
	err := srv.ListenAndServe()

	errorExist := isError(err)
	if errorExist {
		errorLogger.Fatalf("%s: %s", err, "Could not start the webserver")
	}
}
