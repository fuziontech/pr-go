package pr

import (
	"github.com/fuziontech/pr-idl/pb"
	"time"
)

// This type is used to represent traits in messages that support it.
// It is a free-form object so the application can set any value it sees fit but
// a few helper method are defined to make it easier to instantiate traits with
// common fields.
// Here's a quick example of how this type is meant to be used:
//
//	analytics.Identify{
//		UserId: "0123456789",
//		Traits: analytics.NewTraits()
//			.SetFirstName("Luke")
//			.SetLastName("Skywalker")
//			.Set("Role", "Jedi"),
//	}
//
// The specifications can be found at https://segment.com/docs/spec/identify/#traits

type Traits pb.Traits

func NewTraits() Traits {
	return Traits{}
}

func (t Traits) SetAddress(address string) Traits {
	t.Address = address
	return t
}

func (t Traits) SetAge(age int) Traits {
	t.Age = int32(age)
	return t
}

func (t Traits) SetAvatar(url string) Traits {
	t.Avatar = url
	return t
}

func (t Traits) SetBirthday(date time.Time) Traits {
	//t.Birthday, _ = ptypes.TimestampProto(date)
	return t
}

func (t Traits) SetCreatedAt(date time.Time) Traits {
	//t.CreatedAt, _ = ptypes.TimestampProto(date)
	return t
}

func (t Traits) SetDescription(desc string) Traits {
	t.Description = desc
	return t
}

func (t Traits) SetEmail(email string) Traits {
	t.Email = email
	return t
}

func (t Traits) SetFirstName(firstName string) Traits {
	t.FirstName = firstName
	return t
}

func (t Traits) SetGender(gender string) Traits {
	t.Gender = gender
	return t
}

func (t Traits) SetLastName(lastName string) Traits {
	t.LastName = lastName
	return t
}

func (t Traits) SetName(name string) Traits {
	t.Name = name
	return t
}

func (t Traits) SetPhone(phone string) Traits {
	t.Phone = phone
	return t
}

func (t Traits) SetTitle(title string) Traits {
	t.Title = title
	return t
}

func (t Traits) SetUsername(username string) Traits {
	t.Username = username
	return t
}

func (t Traits) SetWebsite(url string) Traits {
	t.Website = url
	return t
}

func (t Traits) Set(field string, value string) Traits {
	t.ExtraFields[field] = value
	return t
}
