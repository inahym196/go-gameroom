package model

type User struct {
	Order string `json:"order"`
	Name  string `json:"name"`
}

type Board struct {
	Id      int        `json:"id"`
	Pieces  [][]string `json:"pieces"`
	Turn    int        `json:"turn"`
	Status  string     `json:"status"`
	Players struct {
		First User `json:"first"`
		Draw  User `json:"draw"`
	} `json:"players"`
	Winner       string `json:"winner"`
	LastPutPoint string `json:"last_put_point"`
}
