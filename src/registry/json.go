package registry

import (
	"errors"
	"net/http"
	"regexp"

	resty "gopkg.in/resty.v1"
)

var (
	// ErrNoMorePages is the error returned when there are no more pages left in request
	ErrNoMorePages = errors.New("No more pages")
)

// Matches an RFC 5988 (https://tools.ietf.org/html/rfc5988#section-5)
// Link header. For example,
//
//    <http://registry.example.com/v2/_catalog?n=5&last=tag5>; type="application/json"; rel="next"
//
// The URL is _supposed_ to be wrapped by angle brackets `< ... >`,
// but e.g., quay.io does not include them. Similarly, params like
// `rel="next"` may not have quoted values in the wild.
var nextLinkRE = regexp.MustCompile(`^ *<?([^;>]+)>? *(?:;[^;]*)*; *rel="?next"?(?:;.*)?`)

func getNextLink(resp *resty.Response) (string, error) {
	for _, link := range resp.Header()[http.CanonicalHeaderKey("Link")] {
		parts := nextLinkRE.FindStringSubmatch(link)
		if parts != nil {
			return parts[1], nil
		}
	}
	return "", ErrNoMorePages
}
