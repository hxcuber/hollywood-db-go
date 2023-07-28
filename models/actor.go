package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	ActorIdWidth   = 4
	FirstNameWidth = 15
	LastNameWidth  = 15
	GenderWidth    = 7
	DOBWidth       = 11
)

type Actor struct {
	ActorId   int       `gorm:"column=id:primaryKey"`
	FirstName string    `gorm:"column=first_name"`
	LastName  string    `gorm:"column=last_name"`
	Gender    string    `gorm:"column=gender"`
	DOB       time.Time `gorm:"column=dob"`
}

type Actors []Actor

func (_ Actor) TableName() string {
	return "actors"
}

func (a Actor) String() string {
	return fmt.Sprintf("Actor ID: %-*d; First Name: %-*s; Last Name: %-*s; Gender: %-*s; DOB %-*s;\n", ActorIdWidth, a.ActorId, FirstNameWidth, a.FirstName, LastNameWidth, a.LastName, GenderWidth, a.Gender, DOBWidth, a.DOB.Format(time.DateOnly))
}

func (as Actors) String() string {
	var sb strings.Builder
	for _, a := range as {
		sb.WriteString(a.String())
	}
	return sb.String()
}
