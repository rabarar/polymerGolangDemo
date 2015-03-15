//
//  Created by Robert A. Baruch on 03/14/2015
//  Copyright (c) 2015 Robert A. Baruch, All rights reserved.
//

package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ()

type DemoPages struct {
	pages map[string]DemoPage
}

type DemoPage struct {
	urlString string
	callback  gin.HandlerFunc
	template  string
	context   interface{}
}

type Movies struct {
	Title string `json:"title"`
}

var (
	nullFunc = func(c *gin.Context) {
		c.String(http.StatusOK, "default func - stub - replace with actual function")
	}
)

func main() {

	var dp DemoPages
	var foo map[string]int = map[string]int{"foo": 1, "bar": 2}
	var bar map[string]string = map[string]string{"foo": "Alpha", "bar": "Beta"}
	var grill map[string]string = map[string]string{"foo": "Omega", "bar": "Alpha"}

	dp.pages = map[string]DemoPage{
		"fred":  DemoPage{"/fred", FredFunc, "fred.tmpl", foo},
		"tom":   DemoPage{"/tom", TomFunc, "tom.tmpl", grill},
		"sally": DemoPage{"/sally", SallyFunc, "sally.tmpl", bar},
		"core":  DemoPage{"/feeds/api/rob", coreFunc, "core.tmpl", nil},
	}

	// Use Gin-Gonic for our web framework
	router := gin.Default()
	router.Use(dp.Params())

	// Register a template for each page in the pages table
	var html = template.New("temps")

	// Escape the {{ }} braces used by Polymer and Golang Templates
	html.Delims("{[", "]}")

	// Load all of the templates defined in the pages map
	template.Must(html.ParseFiles(buildTmpArray("./templates", dp.pages)...))
	router.SetHTMLTemplate(html)

	// Serve static content
	router.Static("/static", "./static")

	// Add a default GET for /
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	// for each element in the pages map, add a GET route and corresponding handler
	for _, p := range dp.pages {
		router.GET(p.urlString, p.callback)
	}

	router.Run(":8080")
}

// GET callbacks: Fred
func FredFunc(c *gin.Context) {

	dp, err := GetDPContext(c)

	log.Printf("preparing %s\n", dp.pages["fred"].urlString)
	log.Printf("context = %v\n", dp.pages["fred"].context)

	ctx := (dp.pages["fred"].context).(map[string]int)

	if err != nil {
		log.Printf("failed to get server context\n")
	} else {
		obj := gin.H{"title": dp.pages["fred"].urlString,
			"foo": ctx["foo"],
		}
		c.HTML(http.StatusOK, "fred.tmpl", obj)
	}

}

// GET callbacks; Tom
func TomFunc(c *gin.Context) {

	dp, err := GetDPContext(c)

	ctx := (dp.pages["tom"].context).(map[string]string)

	if err != nil {
		log.Printf("failed to get server context\n")
	} else {
		obj := gin.H{"title": dp.pages["tom"].urlString,
			"foo": ctx["foo"],
		}
		c.HTML(http.StatusOK, "tom.tmpl", obj)
	}

}

// GET callbacks; Sally
func SallyFunc(c *gin.Context) {

	dp, err := GetDPContext(c)

	ctx := (dp.pages["sally"].context).(map[string]string)

	if err != nil {
		log.Printf("failed to get server context\n")
	} else {
		obj := gin.H{"title": dp.pages["sally"].urlString,
			"foo": ctx["foo"],
		}
		c.HTML(http.StatusOK, "sally.tmpl", obj)
	}

}

// GET callbacks; core
func coreFunc(c *gin.Context) {

	// dp, err := GetDPContext(c)
	var core []Movies = []Movies{
		Movies{"yo bro"},
		Movies{"sambo"},
		Movies{"huh"},
	}

	c.JSON(http.StatusOK, core)

}

// Middleware handler for global Context
func (dp *DemoPages) Params() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Get("dp")
		if err != nil {
			log.Printf("Setting context\n")
			c.Set("dp", dp)
		}
		c.Next()
	}
}

// Function to retrieve global Context
func GetDPContext(c *gin.Context) (*DemoPages, error) {
	dp, err := c.Get("dp")
	if err != nil {
		log.Printf("failed to find server parameter context")
		return nil, errors.New("failed to find parameter context")
	}

	return dp.(*DemoPages), nil
}

// return an array of templates to pass to ParseFiles
func buildTmpArray(tdir string, p map[string]DemoPage) []string {
	var temps []string = []string{}

	for _, page := range p {
		pe := fmt.Sprintf("%s/%s", tdir, page.template)
		temps = append(temps, pe)
	}

	return temps
}
