package pebble

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

//FeatureDB stores image feature
type FeatureDB struct {
	db *pebble.DB
}

//NewFeatureDB ...
func NewFeatureDB(dataPath string) (*FeatureDB, error) {
	path := fmt.Sprintf("%s", dataPath)
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &FeatureDB{
		db: db,
	}, nil
}

//addFeature add image feature to db
func (f *FeatureDB) AddFeature(imageKey string, pbFeature []byte) error {
	key := []byte(imageKey)
	if err := f.db.Set(key, pbFeature, pebble.NoSync); err != nil {
		return err
	}
	return nil
}

//GetFeature get image feature from db
func (f *FeatureDB) GetFeature(imageKey string) ([]byte, error) {
	key := []byte(imageKey)
	feature, closer, err := f.db.Get(key)
	if err != nil {
		return nil, err
	}
	err = closer.Close()
	return feature, err
}
