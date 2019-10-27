package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nickrisaro/workshop-go/tweet"
)

// API API REST para acceder a la fucnionalidad de tweeter
type API struct {
	repository *tweet.Repository
}

// New construye una nueva API REST con el repositorio recibido
func New(repository *tweet.Repository) *API {
	api := new(API)
	api.repository = repository
	return api
}

// Start inicia la API, queda escuchando en todas las interfaces en el puerto 8080
func (a *API) Start() {
	router := gin.Default()

	router.GET("/tweets", a.listTweets)
	router.POST("/tweet", a.saveTweet)

	router.Run()
}

func (a *API) listTweets(c *gin.Context) {
	c.JSON(http.StatusOK, a.repository.GetAllTweets())
}

func (a *API) saveTweet(c *gin.Context) {
	newTweet := &tweet.Tweet{}
	c.ShouldBindJSON(newTweet)
	a.repository.Save(newTweet)
	// TODO Si el guardado devolvió un error responder con un Status de error y un mensaje acorde
	c.JSON(http.StatusOK, struct{ Message string }{"Tweet saved"})
}
