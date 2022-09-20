package model

type NameModel struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type JokeModel struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}
