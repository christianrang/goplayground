package main

import (
	"bufio"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	monitor := Monitor{Count: 0}

	go func() {
		cmd := exec.Command("go", "run", "cmd/child/main.go")

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			m := scanner.Text()
			msg := stdoutLine(m)
			if msg.Value == "asdf" {
				monitor.Count++
			}
			log.Printf("we have seen %d, 'asdf' messages", monitor.Count)
			log.Printf("%v", msg)
		}

		cmd.Wait()
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": monitor.Count,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type Monitor struct {
	Count int
}

type Message struct {
	Value      string
	StatusCode string
	Date       time.Time
}

func stdoutLine(line string) Message {
	newLine := strings.Split(line, " ")

	date, _ := time.Parse("2006/01/02", newLine[0])

	return Message{
		Date:       date,
		Value:      newLine[1],
		StatusCode: newLine[2],
	}
}
