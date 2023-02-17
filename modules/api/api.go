package api

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/Mrpye/evaluate-js-templates-api/docs"
	"github.com/Mrpye/evaluate-js-templates-api/modules/body_types"
	"github.com/Mrpye/evaluate-js-templates-api/modules/js"
	"github.com/Mrpye/evaluate-js-templates-api/modules/template"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var web_ip string
var web_port string

// @Summary Evaluate JS Code and return result
// @ID post-evaluate-js-code
// @Produce json
// @Param request body body_types.Evaluate.request true true "query params"
// @Success 200 {object}  string "tar file Exported"
// @Failure 404 {string}  string "error"
// @Router /evaluate [post]
func postEvaluate(c *gin.Context) {
	var importRequest body_types.Evaluate

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	result, err := js.EvaluateJSCode(importRequest.Code, importRequest.Model)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, result)

}

// @Summary Builds a template and returns the result
// @ID 	post-build-template
// @Produce json
// @Param request body body_types.Template.request true true "query params"
// @Success 200 {object}  string "tar file Exported"
// @Failure 404 {string}  string "error"
// @Router /template [post]
func postTemplate(c *gin.Context) {
	var importRequest body_types.Template

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&importRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Bad payload")
		return
	}

	result, err := template.ParseTemplate(importRequest)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	//output as text
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, result)

	//c.IndentedJSON(http.StatusCreated, result)

}

// @Summary Check API Endpoint
// @ID check-api-endpoint
// @Produce json
// @Success 200 {string}  string "ok"
// @Router / [get]
func getOK(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "OK")
}

// Function to start web server
func StartWebServer(ip string, port string) {
	//****************
	//Set the variable
	//****************
	web_ip = ip
	web_port = port

	fmt.Println("Starting Web-Server")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/evaluate", postEvaluate)
	router.POST("/template", postTemplate)
	router.GET("/", getOK)

	//**********************************
	//Set up the environmental variables
	//**********************************
	if os.Getenv("WEB_IP") != "" {
		web_ip = os.Getenv("WEB_IP")
	}
	if os.Getenv("WEB_PORT") != "" {
		web_port = os.Getenv("WEB_PORT")
	}

	//****************
	//Start the server
	//****************
	fmt.Printf("Web-Server started %s:%s", web_ip, web_port)
	router.Run(fmt.Sprintf("%s:%s", web_ip, web_port))

}
