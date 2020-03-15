package pr

import "github.com/fuziontech/pr-idl/pb"

type Identify pb.Identify

func (msg Identify) internal() {
	panic(unimplementedError)
}

func (msg Identify) Validate() error {
	if len(msg.UserId) == 0 && len(msg.AnonymousId) == 0 {
		return FieldError{
			Type:  "analytics.Identify",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	return nil
}
