package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestComment_BeforeCreate(t *testing.T) {
	type fields struct {
		ID        int
		UserID    int
		PhotoID   int
		Message   string
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
			name: "Message must not be empty",
			fields: fields{
				ID:      1,
				UserID:  1,
				PhotoID: 1,
				Message: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comment{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				PhotoID:   tt.fields.PhotoID,
				Message:   tt.fields.Message,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			if err := c.BeforeCreate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %v", err, tt.wantErr)
			}
		})
	}
}
