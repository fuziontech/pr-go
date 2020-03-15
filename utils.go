package pr

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func TsToPbTs(ts time.Time) *timestamp.Timestamp {
	pts, _ := ptypes.TimestampProto(ts)
	return pts
}
