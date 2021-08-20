package challenges

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	Description   string `json:"description" gorm:"column:description;size:5000"`
	Level         string `json:"level" gorm:"column:level"`
	ChallengeType string `json:"challenge_type" gorm:"column:challenge_type"`
}
