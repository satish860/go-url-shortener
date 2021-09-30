package handler

import (
	"go-url-shortener/shortner"
	"go-url-shortener/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct{
   LongUrl string `json:"long_url" binding:"required"`
   UserId string  `json:"userid" binding:"required"`
}

func CreateShortUrl(g *gin.Context){
	var creationRequest UrlCreationRequest;
	if err := g.ShouldBind(&creationRequest);  err !=nil{
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortId := shortner.GenerateShortLink(creationRequest.LongUrl,creationRequest.UserId)
	store.SaveUrlMapping(shortId,creationRequest.LongUrl,creationRequest.UserId)

	host := "http://localhost:9808/"
	g.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortId,
	})
}

func HandleShortUrlRedirect(g *gin.Context){
	shortUrl := g.Param("shortUrl")
	longUrl := store.RetrieveInitialUrl(shortUrl)
	g.Redirect(302,longUrl)
}