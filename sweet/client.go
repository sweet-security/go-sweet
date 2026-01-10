package sweet

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	restyClient *resty.Client
	accessToken string
	expires     time.Time
	env         string
	subenv      string
}

type Auth struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

var Version = "build"

const (
	DefaultEnv    = "prod"
	DefaultSubenv = "main"
)

func New(apiKey string, secret string, options ...func(client *ApiClient)) *ApiClient {
	sweetClient := &ApiClient{}
	sweetClient.env = DefaultEnv
	sweetClient.subenv = DefaultSubenv

	for _, o := range options {
		o(sweetClient)
	}

	baseUrl := fmt.Sprintf("https://%s--eapi.%s.sweet.security", sweetClient.subenv, sweetClient.env)
	baseUrl = strings.ReplaceAll(baseUrl, fmt.Sprintf("%s--", DefaultSubenv), "")
	baseUrl = strings.ReplaceAll(baseUrl, fmt.Sprintf(".%s.", DefaultEnv), ".")

	restyClient := resty.New()
	restyClient.SetHeaders(map[string]string{
		"User-Agent": fmt.Sprintf("Go-Sweet/%s", Version),
	})
	restyClient.SetBaseURL(baseUrl)
	restyClient.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.IsError() {
			return fmt.Errorf("error: %s", response.String())
		}
		return nil
	})
	restyClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		reqUrl, err := url.Parse(req.URL)
		if err != nil {
			return err
		}
		if strings.HasPrefix(reqUrl.Path, "/v1/auth/token") {
			// Don't mangle with auth requests
			return nil
		}
		expires := sweetClient.expires
		// we want to make sure the request will work, 10 seconds extra just to start the request should be enough
		if expires.IsZero() && time.Now().After(expires.Add(time.Second*-10)) {
			authRes, err := c.
				R().
				SetBody(map[string]interface{}{"apiKey": apiKey, "secret": secret}).
				SetResult(&Auth{}).
				Post("/v1/auth/token")
			if err != nil {
				return err
			}
			sweetClient.accessToken = authRes.Result().(*Auth).AccessToken
			sweetClient.expires = time.Now().Add(time.Duration(authRes.Result().(*Auth).ExpiresIn) * time.Second)
		}
		req.SetAuthToken(sweetClient.accessToken)

		return nil
	})
	sweetClient.restyClient = restyClient
	return sweetClient
}

func WithEnv(env string) func(*ApiClient) {
	return func(s *ApiClient) {
		s.env = env
	}
}

func WithSubenv(subenv string) func(*ApiClient) {
	return func(s *ApiClient) {
		s.subenv = subenv
	}
}
