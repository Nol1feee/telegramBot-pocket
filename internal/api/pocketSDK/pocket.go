package pocketSDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	host = "https://getpocket.com/"

	endpointToken        = "v3/oauth/request"
	endpointAutorization = "auth/authorize?request_token=%s&redirect_uri=%s"
	endpointAccess       = "v3/oauth/authorize?consumer_key=%s&code=%s"
	endpointAdd          = "v3/add"

	defaultTimeout = time.Second * 5
)

type AutheticationRequest struct {
	Access_token string
	Username     string
}

type Client struct {
	Client      *http.Client
	ConsumerKey string
	RedirectUri string
	AccessToken string
}

func NewClient(consumerkey string, redirectUri string) (*Client, error) {
	if consumerkey == "" {
		return nil, errors.New("consumer key can't be nothing")
	}
	return &Client{
		Client: &http.Client{
			Timeout: defaultTimeout,
		},
		ConsumerKey: consumerkey,
		RedirectUri: redirectUri,
	}, nil
}

func (c *Client) doPostHtpp(jsonData []byte, urlArg string) (url.Values, error) {
	req, err := http.NewRequest(http.MethodPost, host+urlArg, bytes.NewBuffer(jsonData))
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "request generation error")
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF8")

	resp, err := c.Client.Do(req)
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "error sending generated request")
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("status code - %d", resp.StatusCode)
		return url.Values{}, errors.New(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return url.Values{}, err
	}

	val, err := url.ParseQuery(string(body))
	if err != nil {
		return url.Values{}, err
	}

	return val, err
}

func (c *Client) GetRequestToken() (string, error) {
	jsonData, err := json.Marshal(map[string]string{
		"consumer_key": c.ConsumerKey,
		"redirect_uri": c.RedirectUri,
	})

	if err != nil {
		return "", errors.WithMessage(err, "marshalling error")
	}

	body, err := c.doPostHtpp(jsonData, endpointToken)
	if err != nil {
		return "", err
	}

	res := body.Get("code")

	return res, nil
}

func (c *Client) GetAutorizationUrl(endpointToken string, redirectUri string) (string, error) {
	if endpointToken == "" || c.RedirectUri == "" {
		return "", errors.New("request token and redirect uri can't be nothing")
	}

	return fmt.Sprintf(host+endpointAutorization, endpointToken, redirectUri), nil
}

func (c *Client) Authetication(requestToken string) (*AutheticationRequest, error) {
	jsonData, err := json.Marshal(map[string]string{
		"consumer_key": c.ConsumerKey,
		"code":         requestToken,
	})
	if err != nil {
		return &AutheticationRequest{}, err
	}

	res, err := c.doPostHtpp(jsonData, endpointAccess)
	if err != nil {
		return &AutheticationRequest{}, err
	}

	accessToken, username := res.Get("access_token"), res.Get("username")
	if accessToken == "" || username == "" {
		return &AutheticationRequest{}, errors.New("access token && username can't be nothing")
	}

	return &AutheticationRequest{
		Access_token: accessToken,
		Username:     username,
	}, nil
}

func (c *Client) Add(urlToAdd string) error {
	if urlToAdd == "" {
		return errors.New("url can't be nothing")
	}

	jsonData, err := json.Marshal(map[string]string{
		"url":          urlToAdd,
		"consumer_key": c.ConsumerKey,
		"access_token": c.AccessToken,
	})
	if err != nil {
		return err
	}

	_, err = c.doPostHtpp(jsonData, endpointAdd)
	if err != nil {
		return err
	}
	return nil
}
