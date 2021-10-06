package repo

import (
	"bufio"
	"io"
	"pipe-mbx/model"

	"github.com/pkg/errors"
)

type repo struct{}

// GetDataRepo is a repository of data models
type GetDataRepo interface {
	GetData(opts GetDataOpts) (io.Reader, error)
}

// SaveRepo saves data
type SaveRepo interface {
	Save(r io.Reader, w io.Writer) error
	Upload() error
}

// NewSaveRepo creates a new save repo
func NewSaveRepo() (SaveRepo, error) {
	return &repo{}, nil
}

type GetDataOpts struct {
	RawDataPath string
}

func (r *repo) Save(geojsonld io.Reader, file io.Writer) error {

	file.Write([]byte(`
{
	"type": "FeatureCollection",
	"features": [`))

	scanner := bufio.NewScanner(geojsonld)

	for scanner.Scan() {
		file.Write([]byte(scanner.Text()))
		file.Write([]byte(","))
	}

	file.Write([]byte(`
	]
}`))

	_, err := io.Copy(file, geojsonld)
	return errors.Wrap(err, "Error copying bytes")
}

func (r *repo) Upload() error {
	return nil
}

// NewGetDataRepo returns an appropriate data repo from data type
func NewGetDataRepo(dt string) (GetDataRepo, error) {

	switch dt {
	case string(model.DataType土砂災害警戒区域):
		return New土砂災害警戒区域Repo()
	default:
		return nil, errors.Errorf("Unsupported type: %s", dt)
	}
}
