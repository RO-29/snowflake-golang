package snowflake

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const timeCustom = "2019-01-01T00:00:00+00:00"

var (
	// Custom Epoch (in milliseconds) (January 1, 2019 Midnight UTC = 2019-01-01T00:00:00Z)
	customEPOCH int64
	nodeID      int
	maxNodeID   = (int)(math.Pow(2, float64(nodeIDBITS)) - 1)
	maxSequence = (int)(math.Pow(2, sequenceBITS) - 1)

	totalBITS    = uint(64)
	epochBITS    = uint(42)
	nodeIDBITS   = uint(10)
	sequenceBITS = float64(12)
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

//NewSnowFlake service init
func NewSnowFlake() *Snowflake {
	return &Snowflake{
		lastTimestamp: -1,
	}
}

//Snowflake service ...
type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	sequence      int64
}

// GenerateUniqueSequenceID generates unique id ...
func (s *Snowflake) GenerateUniqueSequenceID() (int64, error) {

	currentTimeStamp, err := s.generateCurrentTimeSequence()
	if err != nil {
		return 0, errors.Wrap(err, "generate time sequence")
	}

	// first 42 bits of our ID will be filled with the epoch timestamp. left-shift to achieve this
	id := currentTimeStamp << (totalBITS - epochBITS)

	// fill the next 10 bits with the node ID.
	id |= int64(nodeID << (totalBITS - epochBITS - nodeIDBITS))

	// last 12 bits with the local counter.
	id |= s.sequence
	return id, nil

}

func (s *Snowflake) generateCurrentTimeSequence() (int64, error) {

	s.mu.Lock()
	currentTimeStamp, err := s.getCurrentTimeStamp()
	if err != nil {
		err = errors.WithStack(err)
		return 0, errors.Wrap(err, "get current time stamp")
	}
	s.lastTimestamp = currentTimeStamp
	s.mu.Unlock()
	return currentTimeStamp, nil

}

func (s *Snowflake) getCurrentTimeStamp() (int64, error) {
	currentTimeStamp := getTimeStampMilli()

	if currentTimeStamp < s.lastTimestamp {
		return 0, errors.New("invalid system clock")
	}

	if currentTimeStamp > s.lastTimestamp {
		// reset sequence to start with zero for the next millisecond
		s.sequence = 0
		return currentTimeStamp, nil
	}

	s.sequence = (s.sequence + 1) & int64(maxSequence)
	if s.sequence != 0 {
		return currentTimeStamp, nil
	}

	// Sequence Exhausted, wait till next millisecond.
	return s.waitNextMillis(currentTimeStamp), nil

}

// Block and wait till next millisecond
func (s *Snowflake) waitNextMillis(currentTimeStamp int64) int64 {
	for currentTimeStamp == s.lastTimestamp {
		currentTimeStamp = getTimeStampMilli()
	}
	return currentTimeStamp
}
