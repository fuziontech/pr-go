package pr

import (
	"encoding/json"
	"testing"
	"github.com/fuziontech/pr-idl/pb"
)

func TestContextMarshalJSONLibrary(t *testing.T) {
	c := Context{
		Library: &pb.LibraryInfo{
			Name: "testing",
		},
	}

	if b, err := json.Marshal(c); err != nil {
		t.Error("marshalling context object failed:", err)

	} else if s := string(b); s != `{"library":{"name":"testing"}}` {
		t.Error("invalid marshaled representation of context:", s)
	}
}

func TestContextMarshalJSONExtra(t *testing.T) {
	c := Context{
		ExtraFields: map[string]string{
			"answer": "42",
		},
	}

	if b, err := json.Marshal(c); err != nil {
		t.Error("marshalling context object failed:", err)

	} else if s := string(b); s != `{"answer":42}` {
		t.Error("invalid marshaled representation of context:", s)
	}
}
