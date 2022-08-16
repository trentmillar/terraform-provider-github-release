package github

import "net/http"

// etagTransport allows saving API quota by passing previously stored Etag
// available via context to request headers
type etagTransport struct {
	transport http.RoundTripper
}

func (e etagTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func NewEtagTransport(rt http.RoundTripper) *etagTransport {
	return &etagTransport{transport: rt}
}
