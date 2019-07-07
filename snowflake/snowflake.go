package snowflake

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	timeCustom = "2019-01-01T00:00:00+00:00"
	// Custom Epoch (in milliseconds) (January 1, 2019 Midnight UTC = 2019-01-01T00:00:00Z)
	customEPOCH   int64
	nodeID        int
	lastTimestamp = int64(-1)
	sequence      = int64(0)
	totalBITS     = int(64)
	epochBITS     = int(42)
	nodeIDBITS    = int(10)
	sequenceBITS  = float64(12)
	maxNodeID     = (int)(math.Pow(2, float64(nodeIDBITS)) - 1)
	maxSequence   = (int)(math.Pow(2, sequenceBITS) - 1)
)

//Init the custom epoch
func Init() {
	rand.Seed(time.Now().UnixNano())
	timeMustParse()
	nodeIDGenerator()

}

func timeMustParse() {
	timeObj, err := time.Parse(time.RFC3339, timeCustom)
	if err != nil {
		panic(err)
	}
	customEPOCH = timeObj.UnixNano() / 1e6
}

//Snowflake service ...
type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	sequence      int64
}

// GenerateUniqueSequenceID generates unique id ...
func (s *Snowflake) GenerateUniqueSequenceID() error {
	return errors.New("init")
}
