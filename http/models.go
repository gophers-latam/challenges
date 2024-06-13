package http

import (
	"strings"

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

var TimeZones = map[string]CountryData{
	"argentina":            {Timezone: "America/Argentina/Buenos_Aires", Flag: "🇦🇷"},
	"bolivia":              {Timezone: "America/La_Paz", Flag: "🇧🇴"},
	"brasil":               {Timezone: "America/Sao_Paulo", Flag: "🇧🇷"},
	"chile":                {Timezone: "America/Santiago", Flag: "🇨🇱"},
	"colombia":             {Timezone: "America/Bogota", Flag: "🇨🇴"},
	"costa_Rica":           {Timezone: "America/Costa_Rica", Flag: "🇨🇷"},
	"cuba":                 {Timezone: "America/Havana", Flag: "🇨🇺"},
	"el_Salvador":          {Timezone: "America/El_Salvador", Flag: "🇸🇻"},
	"ecuador":              {Timezone: "America/Guayaquil", Flag: "🇪🇨"},
	"guatemala":            {Timezone: "America/Guatemala", Flag: "🇬🇹"},
	"honduras":             {Timezone: "America/Tegucigalpa", Flag: "🇭🇳"},
	"mexico":               {Timezone: "America/Mexico_City", Flag: "🇲🇽"},
	"nicaragua":            {Timezone: "America/Managua", Flag: "🇳🇮"},
	"panama":               {Timezone: "America/Panama", Flag: "🇵🇦"},
	"paraguay":             {Timezone: "America/Asuncion", Flag: "🇵🇾"},
	"peru":                 {Timezone: "America/Lima", Flag: "🇵🇪"},
	"puerto_Rico":          {Timezone: "America/Puerto_Rico", Flag: "🇵🇷"},
	"republica_Dominicana": {Timezone: "America/Santo_Domingo", Flag: "🇩🇴"},
	"uruguay":              {Timezone: "America/Montevideo", Flag: "🇺🇾"},
	"venezuela":            {Timezone: "America/Caracas", Flag: "🇻🇪"},
}

var FlagToCountry = map[string]string{}

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

func init() {
	for country, data := range TimeZones {
		FlagToCountry[strings.ToLower(data.Flag)] = country
	}
}
