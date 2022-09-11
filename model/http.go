package model

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gopkg.in/olahol/melody.v1"
)

type Request struct {
	url    string
	data   []byte
	params map[string]string
}

func Do(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r Request) Get() ([]byte, error) {
	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return nil, err
	}
	return Do(req)
}

func (r Request) Put() ([]byte, error) {
	req, err := http.NewRequest("PUT", r.url, bytes.NewBuffer(r.data))
	if err != nil {
		return nil, err
	}
	if len(r.params) > 0 {
		params := req.URL.Query()
		for key, value := range r.params {
			params.Add(key, value)
		}
		req.URL.RawQuery = params.Encode()
	}
	return Do(req)
}

func PutPiece(s *melody.Session, point Point, boardId int) ([]byte, error) {

	var req Request
	req.url = "http://localhost:8000/api/v1/boards/" + strconv.Itoa(boardId) + "/pieces/"
	req.params = map[string]string{
		"raw":    strconv.Itoa(point.X),
		"column": strconv.Itoa(point.Y),
	}
	orderType := s.MustGet("order").(string)
	userName := s.MustGet("userName").(string)
	player := Player{User{userName}, Order{orderType}}
	playerJson, err := json.Marshal(player)
	if err != nil {
		return nil, err
	}
	req.data = playerJson
	return req.Put()
}
