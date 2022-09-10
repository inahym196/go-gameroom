package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
	"gopkg.in/olahol/melody.v1"
)

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
	return byteArray, nil
}

func GetMRouter() *melody.Melody {
	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {

		url := "http://localhost:8000/api/v1/boards/0/"
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
	return m
}
