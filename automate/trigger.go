package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"os/exec"
	"github.com/joho/godotenv"
	"path/filepath"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}

	router.GET("/task", func(c *gin.Context) {
		scriptDir := os.Getenv("SCRIPT_DIR")
		if scriptDir == "" {
			c.String(500, "Script directory not specified")
			return
		}

		scriptPath := filepath.Join(scriptDir, "trigger.sh")

		cmd := exec.Command("bash", scriptPath)
		err := cmd.Run()
		if err != nil {
			c.String(500, "Error executing the script")
			return
		}

		c.String(200, "Script executed successfully")
	})

	router.Run(":8080")
}

