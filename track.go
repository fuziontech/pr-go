package pr

import (
	"github.com/fuziontech/pr-idl/pb"
)

type Track pb.Track

func (msg Track) internal() {
	panic(unimplementedError)
}

func (msg Track) Validate() error {
	if len(msg.Event) == 0 {
		return FieldError{
			Type:  "analytics.Track",
			Name:  "Event",
			Value: msg.Event,
		}
	}

	if len(msg.UserId) == 0 && len(msg.AnonymousId) == 0 {
		return FieldError{
			Type:  "analytics.Track",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	return nil
}
