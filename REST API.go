package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Note :storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.
// schema/model of data to be used
type PostPayload struct {
	Name   string `json:"name"`
	Status string
}

func main() {
	// Getenv retrieves the value of the environment variable named by the key. It returns the value, which will be empty if the variable is not present. To distinguish between an empty value and an unset value, use LookupEnv.
	// PORT is available in env file
	port := os.Getenv("PORT")

	// if port retrieved from env then ok otherwise allocate port 8000 to it
	if port == "" {
		port = "8000"
	}

	// creating a map with a key value
	post := map[string]string{
		"Akanksha": "Present",
	}

	// initializing gin router using default()
	r := gin.Default()

	// get request called when / is triggered and prints message to the screen
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Home Page",
		})
	})

	getPost := func(c *gin.Context) {
		res := gin.H{
			"message": post,
		}
		c.JSON(http.StatusUnauthorized, res)
	}
	// get all the data of post
	r.GET("/post", getPost)

	createPost := func(c *gin.Context) {
		// var of type struct PostPayload contains name and status as keys
		var p1 PostPayload
		// Call BindJSON to bind the received JSON/request body to p1
		err := c.ShouldBindJSON(&p1)
		if err != nil {
			log.Print("got error", err)
		}
		post[p1.Name] = "Present"
		res := gin.H{
			"message": p1,
		}
		c.JSON(http.StatusOK, res)
	}
	r.POST("/post", createPost)

	deletePost := func(c *gin.Context) {
		// Context.Params to retrieve the id path parameter from the URL
		// key := c.Params.ByName("id")
		// fetching the value of id from url into key
		key := c.Query("id")
		//to delete the key value pair from a map via key
		delete(post, key)
		res := gin.H{
			"message": key + " Deleted",
		}
		c.JSON(http.StatusOK, res)
	}
	r.DELETE("/post", deletePost)

	updatePostById := func(c *gin.Context) {
		id := c.Query("id")

		post[id] = "absent"
		res := gin.H{
			"message": id + " marked absent",
			"status":  "Sleeping",
		}
		c.JSON(http.StatusOK, res)
	}
	// as we are sending data as query no need to specify in the url path {:id}
	r.PATCH("/post", updatePostById) //data will be sent via query

	updateValuebyId := func(c *gin.Context) {
		var p1 PostPayload
		val := c.Query("val")
		// unmarshalling using shoulBindJSON to put data from json to struct
		err := c.ShouldBindJSON(&p1)
		if err != nil {
			log.Print("got error", err)
		}
		post[p1.Name] = val
		res := gin.H{
			"message": p1.Name + " updated",
			"status":  "Awake",
		}
		c.JSON(http.StatusOK, res)
	}
	r.PUT("/post", updateValuebyId) //data will be sent via payload and query

	getPostById := func(c *gin.Context) {
		id := c.Query("id")
		// value of map of id index
		val := post[id]
		res := gin.H{
			"message": id + " " + val,
			"status":  "Wants to sleep",
		}
		c.JSON(http.StatusOK, res)
	}
	r.GET("/postbyid", getPostById)
	// listen and serve on "localhost:8080" for windows

	// to run/start the server
	// Run function to attach the router to an http.Server and start the server
	r.Run(fmt.Sprintf(":%s", port))
}

//get all posts localhost:8080/post
/// delete post by id:-
// url: http://localhost:8080/post?id=rahul query
// url: http://localhost:8080/post/rahul param
// request method: delete

/* add post POST
url: localhost:8080/post
and json as
 {
	   "name":"Akanksha"
  }
*/

// patch  value changes to absent
// url: localhost:8080/post?id=rahul

// get post by id:-
// url: http://localhost:8080/postbyid?id=rahul
// request method: get

// http://localhost:8080/post
// put call to change the value of post to anything passed in query
// url: localhost:8080/post?val=check
// and json as
// {
//     "name": "shivam"
// }
