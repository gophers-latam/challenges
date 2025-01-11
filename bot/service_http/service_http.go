package service_http

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	svg "github.com/ajstarks/svgo"
	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	chg "github.com/gophers-latam/challenges/http"
	"github.com/gophers-latam/challenges/storage"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetChallenge(level, topic string) (*chg.Challenge, error) {
	var res []chg.Challenge

	err := storage.Get().Find(&res, "level = ? and challenge_type = ? and active = ?", level, topic, 1).Error
	if err != nil {
		return &chg.Challenge{}, err
	}

	l := len(res)
	if l == 0 {
		return &chg.Challenge{}, sql.ErrNoRows
	}

	i, err := helpers.IntnCrypt(l)

	return &res[i], err
}

func GetFact() (*chg.Fact, error) {
	var res []chg.Fact
	var count int64

	storage.Get().Model(&chg.Fact{}).Count(&count)

	i, err := helpers.IntnCrypt(int(count))
	if i == 0 {
		i++
	}

	err = storage.Get().Find(&res, "id = ?", i).Error
	if err != nil {
		return &chg.Fact{}, err
	}

	if len(res) == 0 {
		return &chg.Fact{}, sql.ErrNoRows
	}

	return &res[0], err
}

func GetEvents() (*[]chg.Event, error) {
	var res []chg.Event

	err := storage.Get().Find(&res).Error
	if err != nil {
		return &res, err
	}

	l := len(res)
	if l == 0 {
		return &res, sql.ErrNoRows
	}

	return &res, err
}

func GetCommand(cmd string) (*chg.Command, error) {
	var res []chg.Command

	err := storage.Get().Find(&res, "cmd = ?", cmd).Error
	if err != nil {
		return &chg.Command{}, err
	}

	l := len(res)
	if l == 0 {
		return &chg.Command{}, sql.ErrNoRows
	}

	return &res[0], err
}

func GetHours(hour, country string) (string, error) {
	var b bytes.Buffer
	args := strings.Split(hour, ":")
	if len(args) != 2 {
		return "", errors.New("invalid time format. Please use HH:MM format")
	}

	h, err := strconv.Atoi(args[0])
	if err != nil {
		return "", errors.New("invalid hour format")
	}
	m, err := strconv.Atoi(args[1])
	if err != nil {
		return "", errors.New("invalid minute format")
	}

	// Check if country has 1 characters and look up in FlagToCountry to assign the country
	if utf8.RuneCountInString(country) == 2 {
		if newCountry, ok := chg.FlagToCountry[country]; ok {
			country = newCountry
		}
	}

	countryCase := cases.Title(language.Und).String(country)
	timeZoneInfo, ok := chg.TimeZones[countryCase]
	if !ok {
		return "", errors.New("unknown country")
	}

	loc, err := time.LoadLocation(timeZoneInfo.Timezone)
	if err != nil {
		return "", errors.New("unable to load timezone")
	}

	now := time.Now().UTC()
	inTime := time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, loc)
	originTime := inTime.In(loc)

	tzones := make([]string, 0, len(chg.TimeZones))
	for key := range chg.TimeZones {
		tzones = append(tzones, key)
	}
	sort.Strings(tzones)

	b.WriteString(fmt.Sprintf("🕒 %s **%s**: `%s` hrs\n", timeZoneInfo.Flag, countryCase, inTime.Format("15:04")))
	for _, tz := range tzones {
		if tz == countryCase {
			continue
		}
		loc, err := time.LoadLocation(chg.TimeZones[tz].Timezone)
		if err != nil {
			continue
		}
		lTime := originTime.In(loc)
		b.WriteString(fmt.Sprintf("🕒 %s **%s**: `%s` hrs\n", chg.TimeZones[tz].Flag, tz, lTime.Format("15:04")))
	}

	return b.String(), nil
}

func GetGopher() *discordgo.File {
	var buf bytes.Buffer
	cv := svg.New(&buf)
	width := 500
	height := 500
	cv.Start(width, height)
	cv.Rect(0, 0, width, height, "fill:gray")

	// Draw Gopher's body
	body := helpers.RandColor()
	cv.Ellipse(width/2, height/2, 150, 200, fmt.Sprintf("fill:rgb(%d,%d,%d)", body.R, body.G, body.B))
	// Ears
	cv.Ellipse(width/2-70, height/2-170, 30, 40, fmt.Sprintf("fill:rgb(%d,%d,%d)", body.R, body.G, body.B))
	cv.Ellipse(width/2+70, height/2-170, 30, 40, fmt.Sprintf("fill:rgb(%d,%d,%d)", body.R, body.G, body.B))
	// Draw Gopher's eyes
	eye := color.RGBA{255, 255, 255, 255} // White color
	cv.Circle(width/2-50, height/2-80, 30, fmt.Sprintf("fill:rgb(%d,%d,%d)", eye.R, eye.G, eye.B))
	cv.Circle(width/2+50, height/2-80, 30, fmt.Sprintf("fill:rgb(%d,%d,%d)", eye.R, eye.G, eye.B))
	// Draw Gopher's pupils
	pupil := helpers.RandColor()
	cv.Circle(width/2-50, height/2-80, 10, fmt.Sprintf("fill:rgb(%d,%d,%d)", pupil.R, pupil.G, pupil.B))
	cv.Circle(width/2+50, height/2-80, 10, fmt.Sprintf("fill:rgb(%d,%d,%d)", pupil.R, pupil.G, pupil.B))
	// Draw Gopher's teeth
	teeth := color.RGBA{255, 255, 255, 255} // White color
	cv.Rect(width/2-18, height/2-20, 15, 30, fmt.Sprintf("fill:rgb(%d,%d,%d)", teeth.R, teeth.G, teeth.B))
	cv.Rect(width/2+3, height/2-20, 15, 30, fmt.Sprintf("fill:rgb(%d,%d,%d)", teeth.R, teeth.G, teeth.B))
	// Draw Gopher's nose
	nose := color.RGBA{0, 0, 0, 255} // Black color
	cv.Circle(width/2, height/2-60, 15, fmt.Sprintf("fill:rgb(%d,%d,%d)", nose.R, nose.G, nose.B))
	// Draw Gopher's feet
	feet := color.RGBA{252, 208, 180, 255}
	cv.Ellipse(width/2-70, height/2+180, 40, 20, fmt.Sprintf("fill:rgb(%d,%d,%d)", feet.R, feet.G, feet.B))
	cv.Ellipse(width/2+70, height/2+180, 40, 20, fmt.Sprintf("fill:rgb(%d,%d,%d)", feet.R, feet.G, feet.B))
	// Draw Gopher's arms
	cv.Ellipse(width/2-110, height/2, 20, 50, fmt.Sprintf("fill:rgb(%d,%d,%d)", feet.R, feet.G, feet.B))
	cv.Ellipse(width/2+110, height/2, 20, 50, fmt.Sprintf("fill:rgb(%d,%d,%d)", feet.R, feet.G, feet.B))

	cv.End()

	icon, _ := oksvg.ReadIconStream(bytes.NewReader([]byte(buf.String())))
	icon.SetTarget(0, 0, float64(width), float64(height))
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	icon.Draw(rasterx.NewDasher(width, height, rasterx.NewScannerGV(width, height, rgba, rgba.Bounds())), 1)

	var newBuf bytes.Buffer
	_ = png.Encode(&newBuf, rgba)

	file := &discordgo.File{
		Name:        "gopher.png",
		ContentType: "image/png",
		Reader:      bytes.NewReader(newBuf.Bytes()),
	}

	return file
}

func GetWaifu() (*chg.Waifu, error) {
	var res []chg.Waifu
	var count int64

	storage.Get().Model(&chg.Waifu{}).Count(&count)

	i, err := helpers.IntnCrypt(int(count))
	if i == 0 {
		i++
	}

	err = storage.Get().Find(&res, "id = ?", i).Error
	if err != nil {
		return &chg.Waifu{}, err
	}

	if len(res) == 0 {
		return &chg.Waifu{}, sql.ErrNoRows
	}

	return &res[0], err
}
