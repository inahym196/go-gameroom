package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Board struct {
	Id           int        `json:"id"`
	Pieces       [][]string `json:"pieces"`
	Turn         int        `json:"turn"`
	Status       string     `json:"status"`
	Players      _Players   `json:"players"`
	Winner       string     `json:"winner"`
	LastPutPoint string     `json:"last_put_point"`
}

type _Players struct {
	First string `json:"first"`
	Draw  string `json:"draw"`
}

func getBoard(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var board Board
	json.Unmarshal(byteArray, &board)
	return byteArray, nil
}

func main() {

	url := "http://localhost:8000/api/v1/boards/0/"

	engine := gin.Default()
	m := melody.New()
	engine.Static("/public", "./public")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Go GameRoom",
		})
	})
	engine.GET("/rooms/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(http.StatusOK, "room.tmpl", gin.H{
			"title": "Go GameRoom",
			"id":    id,
		})
	})
	engine.GET("/rooms/:id/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {

		request := gjson.Get(string(msg), "request")
		switch request.String() {
		case "board":
			byteArray, err := getBoard(url)
			if err != nil {
				log.Fatal(err)
			}
			s.Write(byteArray)
		case "put":
			byteArray, err := getBoard(url)
			if err != nil {
				log.Fatal(err)
			}
			m.Broadcast(byteArray)
		}

	})
	engine.Run("127.0.0.1:3000")
}
