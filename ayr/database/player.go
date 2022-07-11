package database

type Player struct {
	Rating      int
	Position    string
	ClubImage   string
	Image       string
	RareType    int
	FullName    string
	UrlName     string
	Id          int
	NationImage string
	Rare        bool
	Version     string
}

type PlayerPlain struct {
	Rating      string `json:"rating"`
	Position    string `json:"position"`
	ClubImage   string `json:"club_image"`
	Image       string `json:"image"`
	RareType    string `json:"rare_type"`
	FullName    string `json:"full_name"`
	UrlName     string `json:"url_name"`
	Id          string `json:"id"`
	NationImage string `json:"nation_image"`
	Rare        string `json:"rare"`
	Version     string `json:"version"`
}