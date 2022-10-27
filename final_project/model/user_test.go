package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestUser_BeforeCreate(t *testing.T) {
	type fields struct {
		ID        int
		Username  string
		Email     string
		Password  string
		Age       int
		CreatedAt *time.Time
		UpdatedAt *time.Time
	}

	type args struct {
		db *gorm.DB
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Username must not be empty",
			fields: fields{
				ID:       1,
				Username: "",
				Email:    "ada@ada.com",
				Password: "12345678",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Email must not be empty",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "",
				Password: "12345678",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Email invalid",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "ada@a",
				Password: "12345678",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Password must not be empty",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "ada@ada.com",
				Password: "",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Password length at least have 6 characters",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "ada@ada.com",
				Password: "123",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Age must not be empty",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "ada@ada.com",
				Password: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Age min 8",
			fields: fields{
				ID:       1,
				Username: "ada",
				Email:    "ada@ada.com",
				Password: "12345678",
				Age:      8,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				Username:  tt.fields.Username,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				Age:       tt.fields.Age,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			if err := u.BeforeCreate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %v", err, tt.wantErr)
			}
		})
	}
}
