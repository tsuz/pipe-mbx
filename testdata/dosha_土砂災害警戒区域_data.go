package testdata

import (
	"archive/zip"
	"os"

	"github.com/pkg/errors"
)

var F1土砂災害警戒区域 string = `{
	"type": "FeatureCollection",
	"features": [{ 
		"type": "Feature",
		"properties": {
		"A33_001": "3",
		"A33_002": "3", 
		"A33_003": "01", 
		"A33_004": "3-9-206-669-0022", 
		"A33_005": "霧里6", 
		"A33_006": "釧路市", 
		"A33_007": "9999/1/1", 
		"A33_008": "0"
		}, 
	"geometry": { 
		"type": "MultiPolygon", 
		"coordinates": [ [ [ 
			[ 143.799973406, 42.963843186 ], [ 143.799562832, 42.963540014 ] 
		] ] ]
	}
}`
var F2土砂災害警戒区域 string = `
{
	"type": "FeatureCollection",
	"features": [{ 
		"type": "Feature",
		"properties":  { 
			"A33_001": "3", 
			"A33_002": "3", 
			"A33_003": "01", 
			"A33_004": "9-10-422", 
			"A33_005": "入境学", 
			"A33_006": "釧路郡釧路町入境学", 
			"A33_007": "9999/1/1", 
			"A33_008": "0"
		}, 
		"geometry": { 
		"type": "MultiPolygon", 
		"coordinates": [ [ [ 
			[ 144.677679851, 42.939062999 ], [ 144.677874066, 42.939025220 ] 
		] ] ]
	}
}
`

func Create土砂災害警戒区域Data(path string) error {

	outFile, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "Error creating file: %+v", path)
	}
	defer outFile.Close()
	w := zip.NewWriter(outFile)
	f1path := path + "/GeoJSON/file1.geojson"
	f1, err := w.Create(f1path)
	if err != nil {
		return errors.Wrapf(err, "Error creating file: %+v", f1path)
	}
	f1.Write([]byte(F1土砂災害警戒区域))
	f2path := path + "/GeoJSON/file2.geojson"
	f2, err := w.Create(f2path)
	if err != nil {
		return errors.Wrapf(err, "Error creating file: %+v", f2path)
	}
	f2.Write([]byte(F2土砂災害警戒区域))

	return w.Close()
}

func Tear土砂災害警戒区域Data(path string) error {
	return os.Remove(path)
}
