package model

import "github.com/pkg/errors"

// OutcomeType is a type of outcome
type OutcomeType string

// OutcomeType enums
const (
	OutcomeGeoJSON     OutcomeType = "geojson"
	OutcomeMapboxLayer OutcomeType = "mapbox-layer"
	OutcomeNone        OutcomeType = ""
)

// NewOutcomeType generates a new outcome type from string if supported
func NewOutcomeType(s string) (OutcomeType, error) {
	switch s {
	case string(OutcomeGeoJSON):
		return OutcomeGeoJSON, nil
	case string(OutcomeMapboxLayer):
		return OutcomeMapboxLayer, nil
	default:
		return OutcomeNone, errors.Errorf("Unsupported outcome type: %s", s)
	}
}
