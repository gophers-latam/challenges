package challenges

import "gorm.io/gorm"

type Level string
type ChallengeType string

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
	Description   string        `json:"description" gorm:"column:description;size:5000"`
	Level         Level         `json:"level" gorm:"column:level"`
	ChallengeType ChallengeType `json:"challenge_type" gorm:"column:challenge_type"`
}
