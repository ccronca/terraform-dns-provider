package client

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (client *Client) doRequest(query string) error {
	req, err := client.createRequest(query)
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

func (client *Client) createRequest(query string) (*http.Request, error) {

	apiUrlStr, err := client.uriForAPI(query)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", apiUrlStr, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(client.apiUser, client.apiPassword))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (client *Client) uriForAPI(query string) (string, error) {
	apiBase, err := url.Parse(client.baseUrl + "?" + query)
	if err != nil {
		return "", err
	}
	return apiBase.String(), nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
