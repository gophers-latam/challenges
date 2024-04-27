package bot

import (
	"errors"
	"math/rand"

	chg "github.com/gophers-latam/challenges/http"
	"github.com/gophers-latam/challenges/storage"
)

func GetChallenge(level, topic string) (*chg.Challenge, error) {
	var res []chg.Challenge

	err := storage.Get().Find(&res, "level = ? and challenge_type = ? and active = ?", level, topic, 1).Error
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

func GetFact() (*chg.Fact, error) {
	var res []chg.Fact

	err := storage.Get().Find(&res).Error
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

func GetCommand(cmd string) (*chg.Command, error) {
	var res []chg.Command

	err := storage.Get().Find(&res, "cmd = ?", cmd).Error
	if err != nil {
		return nil, err
	}

	l := len(res)

	if l == 0 {
		return nil, errors.New("no results found")
	}

	return &res[0], err
}
