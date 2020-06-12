package garoon

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	HttpClient *http.Client
	ApiBase    string
	User       string
	Password   string
}

type ErrorResponse struct {
	Error struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
		Cause     string `json:"cause"`
	} `json:"error"`
}

func NewClient(subdomain, user, password string) (*Client, error) {
	if len(subdomain) == 0 {
		return nil, errors.New("missing subdomain")
	}
	baseUrl := fmt.Sprintf("https://%s.cybozu.com/g/", subdomain)
	return NewClientWithBaseUrl(baseUrl, user, password)
}

func NewClientWithBaseUrl(baseUrl, user, password string) (*Client, error) {
	if len(baseUrl) == 0 {
		return nil, errors.New("missing baseUrl")
	}

	if len(user) == 0 {
		return nil, errors.New("missing user")
	}

	if len(password) == 0 {
		return nil, errors.New("missing password")
	}
	return &Client{
		HttpClient: &http.Client{},
		ApiBase:    baseUrl,
		User:       user,
		Password:   password,
	}, nil
}

func (c *Client) fetchResource(method, path string, data interface{}, out interface{}) error {
	var body io.Reader
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}

		body = bytes.NewReader(b)
	}

	req, err := c.newRequest(method, path, body)
	if err != nil {
		return err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	if res.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		decoder.Decode(&errorResponse)
		return fmt.Errorf(
			"errorCode: %s, message: %s, cause: %s",
			errorResponse.Error.ErrorCode,
			errorResponse.Error.Message,
			errorResponse.Error.Cause,
		)
	}

	if out != nil {
		decoder.Decode(out)
	}

	return nil
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api/v1/%s", c.ApiBase, path))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	basicAuth := fmt.Sprintf("%s:%s", c.User, c.Password)
	encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(basicAuth))

	req.SetBasicAuth(c.User, c.Password)
	req.Header.Set("Host", fmt.Sprintf("%s:443", c.ApiBase))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Cybozu-Authorization", encodedBasicAuth)

	return req, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}
