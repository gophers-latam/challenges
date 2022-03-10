package challenges

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChallengeGateway struct {
	*gorm.DB
}

func (g *ChallengeGateway) CreateChallenge(c Challenge) (Challenge, error) {
	c.validate()
	c.Active = true
	err := g.Create(&c).Error
	return c, err
}

func (g *ChallengeGateway) GetChallenges(level Level, challengeType ChallengeType) ([]Challenge, error) {
	var result []Challenge

	if _, ok := Levels[level]; !ok {
		level = defaultLevel
	}

	if _, ok := ChallengeTypes[challengeType]; !ok {
		challengeType = defaultType
	}

	err := g.Find(&result, "level = ? and challenge_type = ? and active = ?", level, challengeType, true).Error

	return result, err
}

func (g *ChallengeGateway) GetChallengeById(id int) (*Challenge, error) {
	var result *Challenge
	err := g.Model(Challenge{}).Find(&result, "id=?", id).Error

	if err != nil {
		zap.L().Error("cannot get challenge", zap.Error(err))
		return nil, err
	}

	if result.ID == 0 {
		return nil, nil
	}

	return result, nil
}
