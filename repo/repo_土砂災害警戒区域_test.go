package repo

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"os"
	"testing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

func Test土砂災害警戒区域(t *testing.T) {

	// setup a mock zip file with contents
	path := "./zipdata.zip"
	outFile, err := os.Create(path)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "Error creating file: %+v", path))
	}
	defer outFile.Close()
	w := zip.NewWriter(outFile)
	f1path := path + "/GeoJSON/file1.geojson"
	f1, err := w.Create(f1path)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "Error creating file: %+v", f1path))
	}
	f1str := `{
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
	f1.Write([]byte(f1str))
	f2path := path + "/GeoJSON/file2.geojson"
	f2, err := w.Create(f2path)
	if err != nil {
		t.Fatal(errors.Wrapf(err, "Error creating file: %+v", f2path))
	}
	f2str := `
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
	f2.Write([]byte(f2str))

	w.Close()

	// run the implementation
	opts := GetDataOpts{
		RawDataPath: path,
	}
	r, _ := New土砂災害警戒区域Repo()
	geojsonr, err := r.GetData(opts)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error getting data"))
	}

	scanner := bufio.NewScanner(geojsonr)
	geojsonld := make([]string, 0)
	for scanner.Scan() {
		txt := scanner.Text()
		geojsonld = append(geojsonld, txt)
	}

	if len(geojsonld) != 2 {
		t.Errorf("Expected 2 features but got %d", len(geojsonld))
	}
	cmpFeature([]byte(geojsonld[0]), []byte(f1str))
	cmpFeature([]byte(geojsonld[1]), []byte(f2str))
}

func cmpFeature(geojsonld, exp []byte) (bool, error) {
	f, err := geojson.UnmarshalFeature(geojsonld)
	if err != nil {
		return false, errors.Wrapf(err, "Expected observed feature to be valid but got %+v", err)
	}
	expf, err := geojson.UnmarshalFeature(exp)
	if err != nil {
		return false, errors.Wrapf(err, "Expected expected feature to be valid but got %+v", err)
	}

	obsstr, err := json.Marshal(f)
	if err != nil {
		return false, errors.Wrap(err, "Expected observed geojsonld to be marshaled")
	}

	expstr, err := json.Marshal(expf)
	if err != nil {
		return false, errors.Wrap(err, "Expected observed geojsonld to be marshaled")
	}

	if string(obsstr) != string(expstr) {
		return false, errors.Errorf("Expected %s but got %s",
			string(expstr),
			string(obsstr))
	}
	return true, nil
}
