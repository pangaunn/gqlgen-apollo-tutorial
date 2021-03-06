// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Launch struct {
	ID       string   `json:"id"`
	Site     *string  `json:"site"`
	Mission  *Mission `json:"mission"`
	Rocket   *Rocket  `json:"rocket"`
	IsBooked bool     `json:"isBooked"`
}

type LaunchConnection struct {
	Cursor   string    `json:"cursor"`
	HasMore  bool      `json:"hasMore"`
	Launches []*Launch `json:"launches"`
}

type Mission struct {
	Name         *string `json:"name"`
	MissionPatch *string `json:"missionPatch"`
}

type Rocket struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type TripUpdateResponse struct {
	Success  bool      `json:"success"`
	Message  *string   `json:"message"`
	Launches []*Launch `json:"launches"`
}

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	ProfileImage *string   `json:"profileImage"`
	Trips        []*Launch `json:"trips"`
}

type PatchSize string

const (
	PatchSizeSmall PatchSize = "SMALL"
	PatchSizeLarge PatchSize = "LARGE"
)

var AllPatchSize = []PatchSize{
	PatchSizeSmall,
	PatchSizeLarge,
}

func (e PatchSize) IsValid() bool {
	switch e {
	case PatchSizeSmall, PatchSizeLarge:
		return true
	}
	return false
}

func (e PatchSize) String() string {
	return string(e)
}

func (e *PatchSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PatchSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PatchSize", str)
	}
	return nil
}

func (e PatchSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
