//autogenerated:yes
//nolint:revive,lll
package dynamic_reconfigure

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

const (
	SensorLevels_RECONFIGURE_CLOSE   int8 = 3
	SensorLevels_RECONFIGURE_STOP    int8 = 1
	SensorLevels_RECONFIGURE_RUNNING int8 = 0
)

type SensorLevels struct {
	msg.Package     `ros:"dynamic_reconfigure"`
	msg.Definitions `ros:"byte RECONFIGURE_CLOSE=3,byte RECONFIGURE_STOP=1,byte RECONFIGURE_RUNNING=0"`
}