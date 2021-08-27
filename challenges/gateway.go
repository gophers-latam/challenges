package challenges

import "gorm.io/gorm"

type ChallengeGateway struct {
	*gorm.DB
}

func (g *ChallengeGateway) CreateChallenge(c Challenge) (Challenge, error) {
	err := g.Create(&c).Error
	return c, err
}

func (g *ChallengeGateway) GetChallenges(level Level, challengeType ChallengeType) ([]Challenge, error) {
	var result []Challenge

	if _, ok := Levels[level]; !ok {
		level = "default"
	}

	if _, ok := ChallengeTypes[challengeType]; !ok {
		challengeType = "default"
	}

	err := g.Find(&result, "level = ? and challenge_type = ?", level, challengeType).Error
	return result, err
}
