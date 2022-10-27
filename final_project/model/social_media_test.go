package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestSocialMedia_BeforeCreate(t *testing.T) {
	type fields struct {
		ID             int
		Name           string
		SocialMediaUrl string
		UserID         int
		CreatedAt      *time.Time
		UpdatedAt      *time.Time
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
			name: "Name must not be empty",
			fields: fields{
				ID:             1,
				Name:           "",
				SocialMediaUrl: "fdfd.com",
				UserID:         1,
			},
			wantErr: true,
		},
		{
			name: "Social media url invalid",
			fields: fields{
				ID:             1,
				Name:           "fdfdf",
				SocialMediaUrl: "fdfd",
				UserID:         1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SocialMedia{
				ID:             tt.fields.ID,
				Name:           tt.fields.Name,
				SocialMediaUrl: tt.fields.SocialMediaUrl,
				UserID:         1,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}

			if err := sm.BeforeCreate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %v", err, tt.wantErr)
			}
		})
	}
}
