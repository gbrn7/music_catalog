package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gbrn7/music_catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		exptectedBody      memberships.LoginResponse
		wantErr            bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("accessToken", nil)
			},
			expectedStatusCode: 200,
			exptectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
			wantErr: false,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 400,
			exptectedBody:      memberships.LoginResponse{},
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()

			h := &Handler{
				Engine:  api,
				Service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endPoint := "/memberships/login"

			model := memberships.LoginRequest{
				Email:    "test@gmail.com",
				Password: "password",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endPoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := memberships.LoginResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.exptectedBody, response)
			}
		})
	}
}
