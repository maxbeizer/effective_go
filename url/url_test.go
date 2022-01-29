package url

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tt := []struct {
		name    string
		rawURL  string
		wantURL *URL
		wantErr error
	}{
		{
			name:    "missing scheme",
			rawURL:  "example.com/path",
			wantURL: nil,
			wantErr: missingScheme,
		},
		// {
		// 	name:    "missing host",
		// 	rawURL:  "http:///path",
		// 	wantURL: nil,
		// 	wantErr: errors.New("missing host"),
		// },
		{
			name:    "missing path",
			rawURL:  "http://example.com",
			wantURL: &URL{Scheme: "http", Host: "example.com", Path: ""},
			wantErr: nil,
		},
		{
			name:    "valid",
			rawURL:  "http://example.com/path",
			wantURL: &URL{Scheme: "http", Host: "example.com", Path: "path"},
			wantErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gotURL, gotErr := Parse(tc.rawURL)

			if gotErr != tc.wantErr {
				t.Errorf("got error %v, want %v", gotErr, tc.wantErr)
			}

			if !reflect.DeepEqual(gotURL, tc.wantURL) {
				t.Errorf("got URL %v, want %v", gotURL, tc.wantURL)
			}
		})
	}
}

func TestURLPort(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "no port",
			in:   "example.com",
			want: "",
		},
		{
			name: "port number",
			in:   "example.com:80",
			want: "80",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			u := &URL{Host: tc.in}

			if got := u.Port(); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
