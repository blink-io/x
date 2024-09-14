package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestMy_Type_1(t *testing.T) {
	type Profile struct {
		Location  string
		CreatedAt time.Time
		Interval  time.Duration
	}
	type User struct {
		Name       string
		Age        int
		Profile    Profile
		ProfilePtr *Profile
	}

	var tt = time.Now()

	var u1 = User{
		Name: "ABC",
		Age:  23,
		Profile: Profile{
			Location:  "GZ, Guangdong",
			CreatedAt: tt,
		},
		ProfilePtr: &Profile{
			Location:  "NY, USA",
			CreatedAt: time.Now().UTC(),
		},
	}

	u2 := u1

	fmt.Printf("u1 name ptr: %p\n", &u1.Name)
	fmt.Printf("u2 name ptr: %p\n", &u2.Name)
	fmt.Printf("u1 profile ptr: %p\n", &u1.Profile)
	fmt.Printf("u2 profile ptr: %p\n", &u2.Profile)
}
