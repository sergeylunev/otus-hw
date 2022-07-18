package storagefabric

import (
	"errors"

	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var (
	ErrStorageProblem = errors.New("problem with storage occured")
)

func Create(conf storage.StorageConf) (storage.Storage, error) {
	var storage storage.Storage
	storageType := conf.Type
	if storageType == "MEMO" {
		storage = memorystorage.New()
	} else if storageType == "DB" {
		storage, err := sqlstorage.New(
			conf.User,
			conf.Pass,
			conf.Name,
			conf.Port,
			conf.Host,
		)
		if err != nil {
			return storage, err
		}
	} else {
		return storage, ErrStorageProblem
	}

	return storage, nil
}
