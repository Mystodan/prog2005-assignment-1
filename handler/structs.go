package handler

type Language struct {
	Lang_code string `json:"lang_code"`
	Lang_name string `json:"lang_name"`
}

type Universities struct {
	Name      string     `json:"name"`
	Country   string     `json:"country,omitempty"` // Suppress field in JSON output if it is empty
	Isocode   string     `json:"code"`
	Webpages  []string   `json:"webpages"`
	Languages []Language `json:"languages"`
	Map       string     `json:"map"`
}
