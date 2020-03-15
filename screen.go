package pr

import (
	"github.com/fuziontech/pr-idl/pb"
)

type Screen pb.Screen

func (msg Screen) internal() {
	panic(unimplementedError)
}

func (msg Screen) Validate() error {
	if len(msg.UserId) == 0 && len(msg.AnonymousId) == 0 {
		return FieldError{
			Type:  "analytics.Screen",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	return nil
}
