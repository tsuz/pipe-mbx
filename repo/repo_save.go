package repo

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

func (r *repo) Save(geojsonld io.Reader, file io.Writer) error {

	file.Write([]byte(`
{
	"type": "FeatureCollection",
	"features": [`))

	scanner := bufio.NewScanner(geojsonld)

	var first bool = true
	for scanner.Scan() {
		if !first {
			file.Write([]byte(","))
		}
		if first {
			first = false
		}
		file.Write([]byte(scanner.Text()))
	}

	file.Write([]byte(`
	]
}`))

	_, err := io.Copy(file, geojsonld)
	return errors.Wrap(err, "Error copying bytes")
}
