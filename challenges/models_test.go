package challenges

import (
	"testing"
)

func TestChallenge_validate(t *testing.T) {
	type fields struct {
		Description   string
		Level         Level
		ChallengeType ChallengeType
		Active        bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "super easy",
			fields: fields{
				Description:   "1+1",
				Level:         "super easy",
				ChallengeType: "frontend",
			},
		},
		{
			name: "fullstack",
			fields: fields{
				Description:   "1+1",
				Level:         "medium",
				ChallengeType: "fullstack",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Challenge{
				Description:   tt.fields.Description,
				Level:         tt.fields.Level,
				ChallengeType: tt.fields.ChallengeType,
				Active:        tt.fields.Active,
			}
			c.validate()
			if tt.name == "super easy" {
				if c.Level != "easy" {
					t.Errorf("got: %s, expected: %s", c.Level, "easy")

				}
				if c.ChallengeType != "frontend" {
					t.Errorf("got: %s, expected: %s", c.ChallengeType, "frontend")

				}
				return
			}

			if tt.name == "fullstack" {
				if c.Level != "medium" {
					t.Errorf("got: %s, expected: %s", c.Level, "medium")

				}
				if c.ChallengeType != "backend" {
					t.Errorf("got: %s, expected: %s", c.ChallengeType, "backend")

				}
				return
			}
		})
	}
}
