package service

import (
	"os"
	"pipe-mbx/repo"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Run is the main function that is called from main
func (s *svc) Run() error {

	dataOpts := repo.GetDataOpts{
		RawDataPath: s.opts.RawDataPath,
	}

	dataRepo, err := repo.NewGetDataRepo(s.opts.DataType)
	if err != nil {
		return errors.Wrap(err, "Error getting datarepo")
	}

	log.Info("Fetching Data...")
	geojsonr, err := dataRepo.GetData(dataOpts)
	if err != nil {
		return errors.Wrap(err, "Error getting data from repo")
	}

	log.Info("Saving Data...")
	file, err := os.Create("merged.geojson")
	if err != nil {
		return errors.Wrap(err, "Error creating file")
	}
	if err := s.repo.Save(geojsonr, file); err != nil {
		return errors.Wrap(err, "Error saving data")
	}

	return errors.Wrap(err, "Error coyping reader to writer")
}
