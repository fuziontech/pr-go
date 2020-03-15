package pr

import "github.com/fuziontech/pr-idl/pb"

type Page pb.Page

func (msg Page) internal() {
	panic(unimplementedError)
}

func (msg Page) Validate() error {
	if len(msg.UserId) == 0 && len(msg.AnonymousId) == 0 {
		return FieldError{
			Type:  "analytics.Page",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	return nil
}
