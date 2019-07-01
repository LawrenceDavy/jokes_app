package main 

import (
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Jokes information structure
type Joke struct {
	ID int `json:"id" binding:"required"`
	Likes int `json:"likes"`
	Joke string `json:"joke" binding:"required"`
}

// List of jokes
var jokes = []Joke {
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
  	Joke{3, 0, "How many apples grow on a tree? All of them."},
  	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
  	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
  	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
  	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
  

}

func main() {
	// Default gin router
	router := gin.Default()

	// Serve static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	api.GET("/jokes", JokeHandler) // Gets list of available jokes
	api.POST("/jokes/like/:jokeID", LikeJoke) // Saves likes of a certain joke

	// Start server
	router.Run(":3000")
}

// Gets list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message": "Jokes handler not implemented",
	})
}

// Increments the number of likes a joke gets
func LikeJoke(c *gin.Context) {
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes++
			}
		}

		// Return updated jokes list
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound) // JokeId is invaild
	}
}