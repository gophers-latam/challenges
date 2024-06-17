package http

import (
	"gorm.io/gorm"
)

type (
	Level         string
	ChallengeType string
)

const (
	defaultLevel = "easy"
	defaultType  = "backend"
)

var (
	Levels = map[Level]struct{}{
		"easy":   {},
		"medium": {},
		"hard":   {},
	}

	ChallengeTypes = map[ChallengeType]struct{}{
		"backend":     {},
		"algorithm":   {},
		"concurrency": {},
		"database":    {},
		"web":         {},
		"cli":         {},
		"frontend":    {},
	}
)

type CountryData struct {
	Timezone string
	Flag     string
}

var (
	TimeZones = map[string]CountryData{
		"Argentina":            {Timezone: "America/Argentina/Buenos_Aires", Flag: "🇦🇷"},
		"Bolivia":              {Timezone: "America/La_Paz", Flag: "🇧🇴"},
		"Brasil":               {Timezone: "America/Sao_Paulo", Flag: "🇧🇷"},
		"Chile":                {Timezone: "America/Santiago", Flag: "🇨🇱"},
		"Colombia":             {Timezone: "America/Bogota", Flag: "🇨🇴"},
		"Costa_Rica":           {Timezone: "America/Costa_Rica", Flag: "🇨🇷"},
		"Cuba":                 {Timezone: "America/Havana", Flag: "🇨🇺"},
		"El_Salvador":          {Timezone: "America/El_Salvador", Flag: "🇸🇻"},
		"Ecuador":              {Timezone: "America/Guayaquil", Flag: "🇪🇨"},
		"Guatemala":            {Timezone: "America/Guatemala", Flag: "🇬🇹"},
		"Honduras":             {Timezone: "America/Tegucigalpa", Flag: "🇭🇳"},
		"Mexico":               {Timezone: "America/Mexico_City", Flag: "🇲🇽"},
		"Nicaragua":            {Timezone: "America/Managua", Flag: "🇳🇮"},
		"Panama":               {Timezone: "America/Panama", Flag: "🇵🇦"},
		"Paraguay":             {Timezone: "America/Asuncion", Flag: "🇵🇾"},
		"Peru":                 {Timezone: "America/Lima", Flag: "🇵🇪"},
		"Puerto_Rico":          {Timezone: "America/Puerto_Rico", Flag: "🇵🇷"},
		"Republica_Dominicana": {Timezone: "America/Santo_Domingo", Flag: "🇩🇴"},
		"Uruguay":              {Timezone: "America/Montevideo", Flag: "🇺🇾"},
		"Venezuela":            {Timezone: "America/Caracas", Flag: "🇻🇪"},
	}
	FlagToCountry = map[string]string{}
)

type Command struct {
	gorm.Model
	Cmd  string `json:"cmd" gorm:"column:cmd;size:500"`
	Text string `json:"text" gorm:"column:text;size:10000"`
}

type Fact struct {
	gorm.Model
	Text string `json:"text" gorm:"column:text;size:10000"`
}

type Event struct {
	gorm.Model
	Text string `json:"text" gorm:"column:text;size:5000"`
}

type Challenge struct {
	gorm.Model
	Description   string        `json:"description" gorm:"column:description;size:15000"`
	Level         Level         `json:"level" gorm:"column:level"`
	ChallengeType ChallengeType `json:"challengeType" gorm:"column:challenge_type"`
	Active        int           `json:"active" gorm:"column:active"`
}

func (c *Challenge) validate() {
	if _, ok := Levels[c.Level]; !ok {
		c.Level = defaultLevel
	}

	if _, ok := ChallengeTypes[c.ChallengeType]; !ok {
		c.ChallengeType = defaultType
	}
}

func (c Challenge) ChallengeFmt() string {
	m := `[*challenge*]⤵️
		-**Level:** ` + string(c.Level) + ` -**Type:** ` + string(c.ChallengeType) + `
		-**Description: ** ` + c.Description
	return m
}
