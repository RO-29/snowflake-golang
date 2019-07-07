package snowflake

import "github.com/pkg/errors"

//Snowflake service ...
type Snowflake struct{}

// GenerateUniqueID generates unique id ...
func (s *Snowflake) GenerateUniqueID() error {
	return errors.New("init")
}
