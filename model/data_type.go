package model

import (
	"github.com/pkg/errors"
)

type DataType string

// DataType enums
const (
	DataType土砂災害警戒区域 DataType = "dosha-saigai-keikai-kuiki"
	DataTypeNone     DataType = ""
)

// NewDataType generates a new data type from string if supported
func NewDataType(s string) (DataType, error) {
	switch s {
	case string(DataType土砂災害警戒区域):
		return DataType土砂災害警戒区域, nil
	default:
		return DataTypeNone, errors.Errorf("Unsupported type: %s", s)
	}
}
