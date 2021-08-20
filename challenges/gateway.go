package challenges

import "gorm.io/gorm"

type ChallengeGateway struct {
	*gorm.DB
}

func (g *ChallengeGateway) CreateChallenge(c Challenge) (Challenge, error) {
	err := g.Create(&c).Error
	return c, err
}
