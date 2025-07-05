package model

import (
	"testing"
)

func TestUser_IsValidName(t *testing.T) {
	type fields struct {
		ID    string
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid name",
			fields: fields{
				Name: "John Doe",
			},
			want: true,
		},
		{
			name: "invalid name - too long",
			fields: fields{
				Name: "John Jacob Jingleheimer Schmidt",
			},
			want: false,
		},
		{
			name: "invalid name - empty",
			fields: fields{
				Name: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := u.IsValidName(); got != tt.want {
				t.Errorf("User.IsValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsValidEmail(t *testing.T) {
	type fields struct {
		ID    string
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid email",
			fields: fields{
				Email: "john.doe@example.com",
			},
			want: true,
		},
		{
			name: "invalid email - too long",
			fields: fields{
				Email: "john.doe.john.doe.john.doe@example.com",
			},
			want: false,
		},
		{
			name: "invalid email - empty",
			fields: fields{
				Email: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := u.IsValidEmail(); got != tt.want {
				t.Errorf("User.IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
