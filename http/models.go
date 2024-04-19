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
		"backend":    {},
		"frontend":   {},
		"algorithms": {},
		"services":   {},
		"design":     {},
	}
)

type Command struct {
	gorm.Model
	Cmd  string `json:"cmd" gorm:"column:cmd;size:500"`
	Text string `json:"text" gorm:"column:text;size:10000"`
}

type Challenge struct {
	gorm.Model
	Description   string        `json:"description" gorm:"column:description;size:15000"`
	Level         Level         `json:"level" gorm:"column:level"`
	ChallengeType ChallengeType `json:"challenge_type" gorm:"column:challenge_type"`
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
