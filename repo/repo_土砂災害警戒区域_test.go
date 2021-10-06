package repo

import (
	"bufio"
	"encoding/json"
	"testing"

	testdata "pipe-mbx/testdata"

	geojson "github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

func Test土砂災害警戒区域(t *testing.T) {

	// setup a mock zip file with contents
	var path string = "./zipdata.zip"

	if err := testdata.Create土砂災害警戒区域Data(path); err != nil {
		t.Fatal(err, "Error creating test data")
	}

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
	cmpFeature([]byte(geojsonld[0]), []byte(testdata.F1土砂災害警戒区域))
	cmpFeature([]byte(geojsonld[1]), []byte(testdata.F2土砂災害警戒区域))

	if err := testdata.Tear土砂災害警戒区域Data(path); err != nil {
		t.Fatal(err, "Error tearing down test data")
	}
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
