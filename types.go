package transloapi

const apiHost string = "https://translo.p.rapidapi.com/" // with trailing slash

type APIResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Message string `json:"message"` // RapidAPI's message, not Translo's
}

type Translation struct {
	APIResponse
	TextLang       string `json:"text_lang"`
	TranslatedText string `json:"translated_text"`
}

type BatchTranslation struct {
	APIResponse
	BatchTranslations []Batch `json:"batch_translations"`
}

type Detection struct {
	APIResponse
	Lang string `json:"lang"`
}

type Batch struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
