package model

type User struct {
	ID        int    `json:"id"`
	USER_NAME string `json:"user_name"`
}

type Word struct {
	USER_ID          int    `json:"user_id"`
	ID               int    `json:"id"`
	Word             string `json:"word"`
	Word_Description string `json:"word_description"`
	Word_Language    string `json:"word_language"`
	Created_At       string `json:"created_at"`
	Updated_At       string `json:"updated_at"`
}

type DeleteWord struct {
	USER_ID int `json:"user_id"`
	ID      int `json:"id"`
}

type UpdateWord struct {
	USER_ID          int    `json:"user_id"`
	ID               int    `json:"id"`
	Word             string `json:"word"`
	Word_Description string `json:"word_description"`
}
