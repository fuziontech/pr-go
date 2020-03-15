package pr

import (
	"github.com/fuziontech/pr-idl/pb"
)

type Context pb.Context

func (c Context) ToPB() *pb.Context {
	out := pb.Context(c)
	return &out
}
