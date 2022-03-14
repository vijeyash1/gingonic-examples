// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Why Gin?", Content: `One of the best features of Go is its built-in net/http library that allows you to create an HTTP server effortlessly. However, it is also less flexible and requires some boilerplate code to implement.

	There is no built-in support in Go to handle routes based on a regular expression or a "pattern". You can write code to add this functionality. However, as the number of your applications grows, it is quite likely that you will either repeat such code everywhere or create a library for reuse.
	
	This is the crux of what Gin offers. It contains a set of commonly used functionalities, e.g. routing, middleware support, rendering, that reduce boilerplate code and make writing web applications simpler.`},
	article{ID: 2, Title: "Designing the Application", Content: "Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response"},
	article{ID: 4, Title: "Installing the Dependency", Content: "go get -u github.com/gin-gonic/gin"},
	article{ID: 5, Title: "Create the router", Content: "router := gin.Default()"},
	article{ID: 6, Title: " Load the templates", Content: `router.LoadHTMLGlob("templates/*")`},
	article{ID: 7, Title: "Define the route handler", Content: `router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"index.html",
				// Pass the data that the page uses (in this case, 'title')
				gin.H{
						"title": "Home Page",
				},
		)
	
	})`},
	article{ID: 8, Title: "Start the application", Content: "router.Run()"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
