package messages

import (
	"errors"
	"github.com/tomiok/challenge-svc/challenges"
	"github.com/tomiok/challenge-svc/storage"
	"math/rand"
)

func GetChallenge(level, topic string) (*challenges.Challenge, error) {
	var res []challenges.Challenge

	err := storage.Get().Find(&res, "level=? and challenge_type=?", level, topic).Error

	if err != nil {
		return nil, err
	}

	l := len(res)

	if l == 0 {
		return nil, errors.New("no results found")
	}

	i := rand.Intn(l)

	return &res[i], err
}
