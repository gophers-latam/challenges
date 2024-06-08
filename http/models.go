package http

import "gorm.io/gorm"

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

var TimeZones = map[string]string{
	"Argentina":            "America/Argentina/Buenos_Aires",
	"Bolivia":              "America/La_Paz",
	"Brasil":               "America/Sao_Paulo",
	"Chile":                "America/Santiago",
	"Colombia":             "America/Bogota",
	"Costa_Rica":           "America/Costa_Rica",
	"Cuba":                 "America/Havana",
	"El_Salvador":          "America/El_Salvador",
	"Ecuador":              "America/Guayaquil",
	"Guatemala":            "America/Guatemala",
	"Honduras":             "America/Tegucigalpa",
	"Mexico":               "America/Mexico_City",
	"Nicaragua":            "America/Managua",
	"Panama":               "America/Panama",
	"Paraguay":             "America/Asuncion",
	"Peru":                 "America/Lima",
	"Puerto_Rico":          "America/Puerto_Rico",
	"Republica_Dominicana": "America/Santo_Domingo",
	"Uruguay":              "America/Montevideo",
	"Venezuela":            "America/Caracas",
}

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
