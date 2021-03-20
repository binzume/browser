package browser

import (
	"net/http"
	"net/http/cookiejar"
)

var DefaultUserAgent = "Mozilla/5.0 WebBrowser"

type WebBrowser struct {
	http.Client
	UserAgent string
}

func NewWebBrowser() (*WebBrowser, error) {
	jar, err := cookiejar.New(nil)
	browser := &WebBrowser{Client: http.Client{Jar: jar}, UserAgent: DefaultUserAgent}
	browser.Transport = browser
	return browser, err
}

func (b *WebBrowser) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", b.UserAgent)
	return http.DefaultTransport.RoundTrip(req)
}
