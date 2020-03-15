package pr

import (
	"github.com/fuziontech/pr-idl/pb"
)

type Alias pb.Alias

func (msg Alias) internal() {
	panic(unimplementedError)
}

func (msg Alias) Validate() error {
	if len(msg.UserId) == 0 {
		return FieldError{
			Type:  "analytics.Alias",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	if len(msg.PreviousId) == 0 {
		return FieldError{
			Type:  "analytics.Alias",
			Name:  "PreviousId",
			Value: msg.PreviousId,
		}
	}

	return nil
}
