package app

type UpdateWordModel struct {
	UUID           string `json:"uuid"`
	Word           string `json:"word"`
	TranslatedWord string `json:"translated_word"`
}