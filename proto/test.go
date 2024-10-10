package proto

import (
	"database/sql/driver"
	"errors"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

func (t *Test) Scan(value interface{}) error {
	logrus.Info("calling scanner")

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return protojson.Unmarshal(bytes, t)
}

func (t *Test) Value() (driver.Value, error) {
	logrus.Info("calling valuer")

	return protojson.Marshal(t)
}
