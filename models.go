package GoLanta

// BaseData stores all basic data used in the base.gohtml template.
type BaseData struct {
	Title      string
	StaticPath string
}

// Character stores all info and content used in the website.
type Character struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Nationality  string `json:"nationality"`
	Strength     int    `json:"strength"`
	Agility      int    `json:"agility"`
	Stamina      int    `json:"stamina"`
	Vitality     int    `json:"vitality"`
	Initiative   int    `json:"initiative"`
	Intelligence int    `json:"intelligence"`
	Knowledge    int    `json:"knowledge"`
	Fame         int    `json:"fame"`
}
