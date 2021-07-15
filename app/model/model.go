package model

type Word struct {
	UUID            string `json:"uuid"`
	Word            string `json:"word"`
	WordDescription string `json:"word_description"`
	WordLanguage    string `json:"word_language"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type DeleteWord struct {
	UUID string `json:"uuid"`
}

type UpdateWord struct {
	UUID            string `json:"uuid"`
	Word            string `json:"word"`
	WordDescription string `json:"word_description"`
}
