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
	"Argentina": "America/Argentina/Buenos_Aires",
	"Brasil":    "America/Sao_Paulo",
	"Chile":     "America/Santiago",
	"Colombia":  "America/Bogota",
	"Mexico":    "America/Mexico_City",
	"Peru":      "America/Lima",
	"Venezuela": "America/Caracas",
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
