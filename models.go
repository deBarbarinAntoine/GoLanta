package GoLanta

import "html/template"

// BaseData stores all basic data used in the base.gohtml template.
type BaseData struct {
	Title      string
	StaticPath string
}

// Character stores all info and content used in the website.
type Character struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Strength     int    `json:"strength"`
	Agility      int    `json:"agility"`
	Stamina      int    `json:"stamina"`
	Vitality     int    `json:"vitality"`
	Initiative   int    `json:"initiative"`
	Intelligence int    `json:"intelligence"`
	Knowledge    int    `json:"knowledge"`
	Fame         int    `json:"fame"`
}

// StatsHTML stores all Character stat converted into template.HTML type.
type StatsHTML struct {
	Strength     StatTemplate
	Agility      StatTemplate
	Stamina      StatTemplate
	Vitality     StatTemplate
	Initiative   StatTemplate
	Intelligence StatTemplate
	Knowledge    StatTemplate
	Fame         StatTemplate
}

// StatTemplate stores any Character 's stat for the template's use.
type StatTemplate struct {
	Line    template.HTML
	StatBar template.HTML
}
