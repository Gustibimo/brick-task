package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCustomMiddleware_ValidateSignature(t *testing.T) {
	handlerFunc := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	tests := []struct {
		name       string
		headerKey  string
		headerVal  string
		wantStatus int
	}{
		{
			name:       "ValidSignature",
			headerKey:  "X-Signature-Key",
			headerVal:  "webhook-123-xyz",
			wantStatus: http.StatusOK,
		},
		{
			name:       "InvalidSignature",
			headerKey:  "X-Signature-Key",
			headerVal:  "invalid-key",
			wantStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(tt.headerKey, tt.headerVal)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)
			m := &CustomMiddleware{}
			handler := m.ValidateSignature(handlerFunc)

			err := handler(c)

			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					if he.Code != tt.wantStatus {
						t.Errorf("Unexpected status code: got %d, want %d", he.Code, tt.wantStatus)
					}
					return
				}
			}

			if rec.Code != tt.wantStatus {
				t.Errorf("Unexpected status code: got %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestNewCustomMiddleware(t *testing.T) {
	want := &CustomMiddleware{}
	got := NewCustomMiddleware()
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("NewCustomMiddleware() returned unexpected type: got %T, want %T", got, want)
	}
}
