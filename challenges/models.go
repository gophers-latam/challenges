package challenges

import "gorm.io/gorm"

type Level string
type ChallengeType string

const defaultLevel = "easy"
const defaultType = "backend"

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

type Challenge struct {
	gorm.Model
	Description   string        `json:"description" gorm:"column:description;size:15000"`
	Level         Level         `json:"level" gorm:"column:level"`
	ChallengeType ChallengeType `json:"challenge_type" gorm:"column:challenge_type"`
	Active        bool          `json:"active" gorm:"column:active"`
}

func (c *Challenge) validate() {
	if _, ok := Levels[c.Level]; !ok {
		c.Level = defaultLevel
	}

	if _, ok := ChallengeTypes[c.ChallengeType]; !ok {
		c.ChallengeType = defaultType
	}
}
