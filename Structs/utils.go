package Structs

type Book struct {
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Result struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
}

type Country struct {
	ISO3166Alpha2 string `json:"ISO3166_1_Alpha_2"`
	Official_Name string `json:"official_name"`
}

type LanguageData struct {
	Language string  `json:"language"`
	Authors  int     `json:"authors"`
	Books    int     `json:"books"`
	Fraction float32 `json:"fraction"`
}

type Population struct {
	Population int `json:"population"`
}

type ReadershipData struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Books       int    `json:"books"`
	Authors     int    `json:"authors"`
	Readership  int    `json:"readership"`
}
