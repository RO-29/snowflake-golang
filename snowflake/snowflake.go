package snowflake

import (
	"math"
	"time"

	"github.com/pkg/errors"
)

var (
	nodeID int

	// Custom Epoch (January 1, 2019 Midnight UTC = 2019-01-01T00:00:00Z)
	customEPOCH int64
	timeCustom  = "2019-01-01T00:00:00+00:00"

	totalBITS    = int(64)
	epochBITS    = int(42)
	nodeIDBITS   = int(10)
	sequenceBITS = int(12)
	maxNodeID    = (int)(math.Pow(2, float64(nodeIDBITS)) - 1)
	maxSequence  = (int)(math.Pow(2, float64(sequenceBITS)) - 1)

	// private volatile long lastTimestamp = -1L;
	// private volatile long sequence = 0L;
)

//Init the custom epoch
func Init() {
	timeMustParse()

}

func timeMustParse() {
	timeObj, err := time.Parse(time.RFC3339, timeCustom)
	if err != nil {
		panic(err)
	}
	customEPOCH = timeObj.Unix()
}

//Snowflake service ...
type Snowflake struct{}

// GenerateUniqueID generates unique id ...
func (s *Snowflake) GenerateUniqueID() error {
	return errors.New("init")
}
