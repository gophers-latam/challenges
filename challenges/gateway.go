package challenges

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

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
		level = defaultValue
	}

	if _, ok := ChallengeTypes[challengeType]; !ok {
		challengeType = defaultValue
	}

	err := g.Find(&result, "level = ? and challenge_type = ?", level, challengeType).Error
	return result, err
}

func (g *ChallengeGateway) GetChallengeById(id int) (*Challenge, error) {
	var result Challenge
	err := g.Find("id=?", id).First(&result).Error

	if err != nil {
		zap.L().Error("cannot get challenge", zap.Error(err))
		return nil, err
	}

	return &result, nil
}
