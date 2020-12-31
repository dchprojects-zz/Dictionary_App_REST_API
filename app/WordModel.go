package app

type WordModel struct {
	ID             int    `json:"id"`
	UUID           string `json:"uuid"`
	Word           string `json:"word"`
	TranslatedWord string `json:"translated_word"`
	CreatedDate    string `json:"created_date"`
}
