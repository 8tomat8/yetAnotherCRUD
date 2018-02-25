package entity

import (
	"testing"
	"time"
)

func Test_calculateUserHash(t *testing.T) {
	tests := []struct {
		name          string
		plainPassword string
		want          string
		wantErr       bool
	}{
		{
			"Not empty",
			"SomePassword!@^#&4759a9s8hc9ae",
			"f6c600bdfe1a39c6dc2035c989e0e322", // Hardcoded for particular hash salt
			false,
		},
		{
			"Empty",
			"",
			"92b7e44517577c567b30c54870daaeec", // Hardcoded for particular hash salt
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateUserHash(tt.plainPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateUserHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateUserHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsValid(t *testing.T) {
	type fields struct {
		Username  string
		Password  string
		Firstname string
		Lastname  string
		Sex       string
		Birthdate time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
	}{
		{
			"No mandatory field - Username",
			fields{
				Password:  "Password",
				Firstname: "Firstname",
				Lastname:  "Lastname",
				Sex:       "Sex",
				Birthdate: time.Now(),
			},
			false,
		},
		{
			"No mandatory field - Password",
			fields{
				Username:  "Username",
				Firstname: "Firstname",
				Lastname:  "Lastname",
				Sex:       "Sex",
				Birthdate: time.Now(),
			},
			false,
		},
		{
			"No mandatory field - Firstname",
			fields{
				Username:  "Username",
				Password:  "Password",
				Lastname:  "Lastname",
				Sex:       "Sex",
				Birthdate: time.Now(),
			},
			false,
		},
		{
			"No mandatory field - Lastname",
			fields{
				Username:  "Username",
				Password:  "Password",
				Firstname: "Firstname",
				Sex:       "Sex",
				Birthdate: time.Now(),
			},
			false,
		},
		{
			"No mandatory field - Sex",
			fields{
				Username:  "Username",
				Password:  "Password",
				Firstname: "Firstname",
				Lastname:  "Lastname",
				Birthdate: time.Now(),
			},
			false,
		},
		{
			"No mandatory field - Birthdate",
			fields{
				Username:  "Username",
				Password:  "Password",
				Firstname: "Firstname",
				Lastname:  "Lastname",
				Sex:       "Sex",
			},
			false,
		},
		{
			"Valid",
			fields{
				Username:  "Username",
				Password:  "Password",
				Firstname: "Firstname",
				Lastname:  "Lastname",
				Sex:       "Sex",
				Birthdate: time.Now(),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Username:  tt.fields.Username,
				Password:  tt.fields.Password,
				Firstname: tt.fields.Firstname,
				Lastname:  tt.fields.Lastname,
				Sex:       tt.fields.Sex,
				Birthdate: tt.fields.Birthdate,
			}
			if gotValid := u.IsValid(); gotValid != tt.wantValid {
				t.Errorf("User.IsValid() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
