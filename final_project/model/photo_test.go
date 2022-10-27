package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestPhoto_BeforeCreate(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Caption   string
		PhotoUrl  string
		UserID    int
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
			name: "Title must not be empty",
			fields: fields{
				ID:       1,
				Title:    "",
				Caption:  "",
				PhotoUrl: "https://fdfd.png",
				UserID:   1,
			},
			wantErr: true,
		},
		{
			name: "Photo url must not be empty",
			fields: fields{
				ID:       1,
				Title:    "fdfdf",
				Caption:  "",
				PhotoUrl: "",
				UserID:   1,
			},
			wantErr: true,
		},
		{
			name: "Photo url invalid",
			fields: fields{
				ID:       1,
				Title:    "dfdfd",
				Caption:  "",
				PhotoUrl: "hdfd",
				UserID:   1,
			},
			wantErr: true,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Photo{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Caption:   tt.fields.Caption,
				PhotoUrl:  tt.fields.PhotoUrl,
				UserID:    tt.fields.UserID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}

			if err := p.BeforeCreate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %v", err, tt.wantErr)
			}
		})
	}
}
