package client

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (client *Client) doRequest(dns interface{}) error {
	req, err := client.createRequest(dns)
	if err != nil {
		return err
	}

	// Perform the request
	var resp *http.Response
	resp, err = client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("API error %s: %s", resp.Status, body)
	}
	// Don't care about resp body
	return nil
}

func (client *Client) createRequest(dns) (*http.Request, error) {

	apiUrlStr, err := client.uriForAPI()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", apiUrlStr, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(client.api_user, client.api_password))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *Client) uriForAPI(host, ip_addr string) (string, error) {
	apiBase, err := url.Parse(client.baseUrl)
	if err != nil {
		return "", err
	}
	q := apiBase.Query()
	q.Add("host", host)
	q.Add("ip_addr", ip_addr)
	apiBase.RawQuery = q.Encode()
	return apiBase.String(), nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
