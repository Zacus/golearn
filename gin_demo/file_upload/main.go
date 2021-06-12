package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	//single upload example
	// router := gin.Default()
	// // Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.Static("/", "./public")
	// router.POST("/upload", func(c *gin.Context) {
	// 	name := c.PostForm("name")
	// 	email := c.PostForm("email")

	// 	// Source
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 		return
	// 	}

	// 	filename := filepath.Base(file.Filename)
	// 	if err := c.SaveUploadedFile(file, filename); err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 		return
	// 	}

	// 	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	// })
	// router.Run(":8080")

	//mulitple upload

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	})
	router.Run(":8080")
}
