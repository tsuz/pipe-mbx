package repo

import (
	"archive/zip"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/bcicen/jstream"
	geojson "github.com/paulmach/go.geojson"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type 土砂災害警戒区域 struct{}

// New土砂災害警戒区域Repo is a data repository for dosha-saigai-keikai-kuiki
func New土砂災害警戒区域Repo() (GetDataRepo, error) {
	return &土砂災害警戒区域{}, nil
}

func (s *土砂災害警戒区域) GetData(opts GetDataOpts) (io.Reader, error) {
	r, w := io.Pipe()

	zf, err := zip.OpenReader(opts.RawDataPath)
	if err != nil {
		return r, errors.Wrap(err, "Error opening zip reader")
	}

	go func() {

		defer zf.Close()

		for _, file := range zf.File {

			if strings.Contains(file.Name, "/GeoJSON/") &&
				strings.HasSuffix(file.Name, ".geojson") {

				rc, err := file.Open()
				if err != nil {
					log.Fatal(errors.Wrap(err, "Error opening file"))
				}
				defer rc.Close()

				decoder := jstream.NewDecoder(rc, 2)
				for mv := range decoder.Stream() {

					b, err := json.Marshal(mv.Value)
					if err != nil {
						log.Fatal(errors.Wrap(err, "Error json.Marshal file"))
					}

					// Because data manipulation needs to be done, we need to unmarshal and then re-marshall
					var m geojson.Feature
					err = json.Unmarshal(b, &m)
					if err != nil {
						log.Fatal(errors.Wrap(err, "Error unmarshaling file"))
					}
					// We expect A33_001 to be a string but sometimes we get an integer
					a31, ok := m.Properties["A33_001"].(int)
					if ok {
						log.Info("Updating A33_001")
						m.Properties["A33_001"] = strconv.Itoa(a31)
					}
					// We expect A33_002 to be a string but sometimes we get an integer
					a32, ok := m.Properties["A33_002"].(int)
					if ok {
						log.Info("Updating A33_002")
						m.Properties["A33_002"] = strconv.Itoa(a32)
					}

					b, err = json.Marshal(m)
					if err != nil {
						log.Fatal(errors.Wrap(err, "Error re-marshaling"))
					}

					w.Write(b)
					w.Write([]byte("\n"))
				}
			}

		}
		w.Close()
	}()
	return r, nil
}
