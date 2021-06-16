package start_some_service

import (
	"baotian0506.com/39_config/gocode/mygin/common"
	"baotian0506.com/39_config/gocode/mygin/start_some_service/mysky"
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	language_agent "github.com/SkyAPM/go2sky/reporter/grpc/language-agent"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var tracerV1 = mysky.GetTracer("service_v1")
var tracerV2 = mysky.GetTracer("service_v2")

func loginEndpoint(c *gin.Context) {
	var r = c.Request
	ctx := context.Background()
	span, ctx, err := tracerV1.CreateEntrySpan(ctx, r.Method+"/"+c.Request.URL.Path, func() (s string, e error) {
		log.Println("header = " + c.GetHeader(propagation.Header))
		return c.GetHeader(propagation.Header), nil
	})
	span.SetSpanLayer(language_agent.SpanLayer_Http)
	span.SetComponent(mysky.HttpServerComponentID)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	span.Tag(go2sky.TagStatusCode, strconv.Itoa(200))
	common.Exception(err)
	defer span.End()

	time.Sleep(time.Millisecond * time.Duration((rand.Intn(500) + 1)))
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func loginEndpoint2(c *gin.Context) {
	var r = c.Request
	ctx := context.Background()
	span, ctx, err := tracerV2.CreateEntrySpan(ctx, r.Method+"/"+c.Request.URL.Path, func() (s string, e error) {
		return c.GetHeader(propagation.Header), nil
	})
	span.SetSpanLayer(language_agent.SpanLayer_Http)
	span.SetComponent(mysky.HttpServerComponentID)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	span.Tag(go2sky.TagStatusCode, strconv.Itoa(200))

	common.Exception(err)
	defer span.End()

	time.Sleep(time.Millisecond * time.Duration((rand.Intn(3000) + 100)))
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

var submitEndpoint = loginEndpoint
var readEndpoint = loginEndpoint

func init() {
	go v1Service()
	log.Println("v1Service start")
	go v2Service()
	log.Println("v2Service start")

	time.Sleep(time.Second)
}

func Start() {

}

func v1Service() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	router.Run(":11111")
}

func v2Service() {
	router := gin.Default()

	// Simple group: v1
	v := router.Group("/v2")
	{
		v.GET("/login", loginEndpoint2)
		v.POST("/submit", submitEndpoint)
		v.POST("/read", readEndpoint)
	}
	router.Run(":22222")
}

func some2() {
	gin.Mode()
}
