package service

import (
	"io/ioutil"
	"os"
	"pipe-mbx/repo"
	"pipe-mbx/testdata"
	"testing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
)

func TestRun(t *testing.T) {
	// setup a mock zip file with contents
	var path string = "./zipdata.zip"

	if err := testdata.Create土砂災害警戒区域Data(path); err != nil {
		t.Fatal(err, "Error creating test data")
	}

	newpath := "overwrite_path.geojson"
	sr, _ := repo.NewSaveRepo()
	o := Opts{
		DataType:    "dosha-saigai-keikai-kuiki",
		RawDataPath: path,
		SavePath:    newpath,
	}
	s, err := NewService(sr, o)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error creating service"))
	}
	if err := s.Run(); err != nil {
		t.Fatal(errors.Wrap(err, "Error running service.run"))
	}
	file, err := os.Open(newpath)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error opening new file"))
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(errors.Wrap(err, "Error reading files"))
	}
	if _, err := geojson.UnmarshalFeatureCollection(b); err != nil {
		t.Fatal(errors.Wrap(err, "Error creating feature collection"))
	}
	if err := os.Remove(newpath); err != nil {
		t.Fatal(errors.Wrap(err, "Error removing file"))
	}

	if err := testdata.Tear土砂災害警戒区域Data(path); err != nil {
		t.Fatal(err, "Error creating test data")
	}
}
