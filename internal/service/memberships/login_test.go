package memberships

import (
	"fmt"
	"testing"

	"github.com/gbrn7/music_catalog/internal/configs"
	"github.com/gbrn7/music_catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2a$10$zlp1yJ1MmeTI5Lsz0nlb7uw/S/3YhUWLdbMAj6LUIhOJfC.IGGmsC",
					Username: "rayhan",
				}, nil)
			},
		},
		{
			name: "failed when getUser",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "wrong password",
					Username: "rayhan",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretKey: "abc",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				fmt.Printf("test case: %s\n", tt.name)
				assert.Empty(t, got)
			}
		})
	}
}
