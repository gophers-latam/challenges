package challenges

import "gorm.io/gorm"

type ChallengeGateway struct {
	*gorm.DB
}

func (g *ChallengeGateway) CreateChallenge(c Challenge) (Challenge, error) {
	err := g.Create(&c).Error
	return c, err
}

func (g *ChallengeGateway) GetChallenge(level, challengeType string) ([]Challenge, error) {
	var result []Challenge
	err := g.Find(&result, "level = ? and challenge_type = ?", level, challengeType).Error

	return result, err
}
