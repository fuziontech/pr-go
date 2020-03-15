package pr

import "github.com/fuziontech/pr-idl/pb"

type Group pb.Group

func (msg Group) internal() {
	panic(unimplementedError)
}

func (msg Group) Validate() error {
	if len(msg.GroupId) == 0 {
		return FieldError{
			Type:  "analytics.Group",
			Name:  "GroupId",
			Value: msg.GroupId,
		}
	}

	if len(msg.UserId) == 0 && len(msg.AnonymousId) == 0 {
		return FieldError{
			Type:  "analytics.Group",
			Name:  "UserId",
			Value: msg.UserId,
		}
	}

	return nil
}
