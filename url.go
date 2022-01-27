package url

import (
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

type Error string

func (e Error) Error() string { return string(e) }

const missingScheme = Error("missing scheme")

func Parse(rawURL string) (*URL, error) {
	i := strings.Index(rawURL, "://")

	if i < 0 {
		return nil, missingScheme
	}

	scheme, rest := rawURL[:i], rawURL[i+3:]
	host, path := rest, ""
	if i := strings.Index(rest, "/"); i >= 0 {
		host, path = rest[:i], rest[i+1:]
	}
	return &URL{Scheme: scheme, Host: host, Path: path}, nil
}
