package main

import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"log"
	"../proto"
	"net/http"
	"strconv"
)

func main(){

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

	g.GET("/add/:a/:b", func(ctx *gin.Context){
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Add(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{ "result": fmt.Sprint("response: ", response.Result)})
		
		}	else{
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("error: ", err.Error())})
		}
		
	})

	g.GET("/multiply/:a/:b", func(ctx *gin.Context){

		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}
		
		if response, err := client.Multiply(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint("response: ", response.Result)})
		
		}	else{
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("error: ", err.Error())})
		}
	})

	if err := g.Run("localhost:8080"); err != nil{
		log.Fatalf("Failed to run server: %v", err)
	}
}