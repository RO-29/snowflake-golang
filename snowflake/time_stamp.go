package snowflake

import (
	"time"
)

// Get the current timestamp in milliseconds, adjust for the custom epoch.
func getTimeStampMilli() int64 {
	return time.Now().UnixNano() - customEPOCH
}
