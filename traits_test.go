package pr

import (
	"github.com/golang/protobuf/ptypes"
	"reflect"
	"testing"
	"time"
)

func TestTraitsSimple(t *testing.T) {
	date := time.Now()
	text := "ABC"
	number := 42

	pbdate, _ := ptypes.TimestampProto(date)

	tests := map[string](struct {
		ref Traits
		run func(Traits)
	}){
		"address":     {Traits{Address: text}, func(t Traits) { t.SetAddress(text) }},
		"age":         {Traits{Age: int32(number)}, func(t Traits) { t.SetAge(number) }},
		"avatar":      {Traits{Avatar: text}, func(t Traits) { t.SetAvatar(text) }},
		"birthday":    {Traits{Birthday: pbdate}, func(t Traits) { t.SetBirthday(date) }},
		"createdAt":   {Traits{CreatedAt: pbdate}, func(t Traits) { t.SetCreatedAt(date) }},
		"description": {Traits{Description: text}, func(t Traits) { t.SetDescription(text) }},
		"email":       {Traits{Email: text}, func(t Traits) { t.SetEmail(text) }},
		"firstName":   {Traits{FirstName: text}, func(t Traits) { t.SetFirstName(text) }},
		"lastName":    {Traits{LastName: text}, func(t Traits) { t.SetLastName(text) }},
		"gender":      {Traits{Gender: text}, func(t Traits) { t.SetGender(text) }},
		"name":        {Traits{Name: text}, func(t Traits) { t.SetName(text) }},
		"phone":       {Traits{Phone: text}, func(t Traits) { t.SetPhone(text) }},
		"title":       {Traits{Title: text}, func(t Traits) { t.SetTitle(text) }},
		"username":    {Traits{Username: text}, func(t Traits) { t.SetUsername(text) }},
		"website":     {Traits{Website: text}, func(t Traits) { t.SetWebsite(text) }},
	}

	for name, test := range tests {
		traits := NewTraits()
		test.run(traits)

		if !reflect.DeepEqual(traits, test.ref) {
			t.Errorf("%s: invalid traits produced: %#v\n", name, traits)
		}
	}
}

func TestTraitsMulti(t *testing.T) {
	t0 := Traits{FirstName: "Luke", LastName: "Skywalker"}
	t1 := NewTraits().SetFirstName("Luke").SetLastName("Skywalker")

	if !reflect.DeepEqual(t0, t1) {
		t.Errorf("invalid traits produced by chained setters:\n- expected %#v\n- found: %#v", t0, t1)
	}
}
