// Golang library for https://rapidapi.com/armangokka/api/translo
package translo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type API struct {
	rapidApiKey string
	client      *http.Client
}

func NewAPI(rapidApiKey string) API {
	return API{
		rapidApiKey: rapidApiKey,
		client:      http.DefaultClient,
	}
}

func NewAPIWithClient(rapidApiKey string, client *http.Client) API {
	if client == nil {
		client = http.DefaultClient
	}
	return API{
		rapidApiKey: rapidApiKey,
		client:      client,
	}
}

func (c API) Translate(ctx context.Context, from, to, text string) (Translation, error) {
	params := url.Values{}
	params.Set("from", from)
	params.Set("to", to)
	params.Set("text", text)
	req, err := http.NewRequestWithContext(ctx, "POST", apiHost+"api/v3/translate", bytes.NewBufferString(params.Encode()))
	if err != nil {
		return Translation{}, err
	}
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	req.Header["X-RapidAPI-Key"] = []string{c.rapidApiKey}
	req.Header["X-RapidAPI-Host"] = []string{"translo.p.rapidapi.com"}

	resp, err := c.client.Do(req)
	if err != nil {
		return Translation{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Translation{}, err
	}
	var out Translation
	if err = json.Unmarshal(body, &out); err != nil {
		return Translation{}, fmt.Errorf("Translate: couldn't decode API's answer [%d]: %s->%s %s", resp.StatusCode, from, to, string(body))
	}
	if !out.Ok {
		if out.Error != "" {
			return Translation{}, fmt.Errorf(out.Error)
		}
		if out.Message != "" {
			return Translation{}, fmt.Errorf(out.Message)
		}
		return Translation{}, fmt.Errorf(string(body))
	}
	return out, nil
}

func (c API) BatchTranslate(ctx context.Context, batches []Batch) ([]Batch, error) {
	if len(batches) == 0 {
		return batches, nil
	}
	data, err := json.Marshal(batches)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", apiHost+"api/v3/batch_translate", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["X-RapidAPI-Key"] = []string{c.rapidApiKey}
	req.Header["X-RapidAPI-Host"] = []string{"translo.p.rapidapi.com"}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out BatchTranslation
	if err = json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("BatchTranslate: couldn't decode API's answer [%d]: %s", resp.StatusCode, string(body))
	}
	if !out.Ok {
		if out.Error != "" {
			return nil, fmt.Errorf(out.Error)
		}
		if out.Message != "" {
			return nil, fmt.Errorf(out.Message)
		}
		return nil, fmt.Errorf(string(body))
	}
	return out.BatchTranslations, nil
}

func (c API) Detect(ctx context.Context, text string) (string, error) {
	if len(text) > 200 {
		text = text[:200]
	}
	req, err := http.NewRequestWithContext(ctx, "GET", apiHost+"api/v3/detect?text="+url.PathEscape(text), nil)
	if err != nil {
		return "", err
	}
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["X-RapidAPI-Key"] = []string{c.rapidApiKey}
	req.Header["X-RapidAPI-Host"] = []string{"translo.p.rapidapi.com"}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var out Detection
	if err = json.Unmarshal(body, &out); err != nil {
		return "", fmt.Errorf("Detect: couldn't decode API's answer [%d]: text: %s translation: %s", resp.StatusCode, text, string(body))
	}
	if !out.Ok {
		if out.Error != "" {
			return "", fmt.Errorf(out.Error)
		}
		if out.Message != "" {
			return "", fmt.Errorf(out.Message)
		}
		return "", fmt.Errorf(string(body))
	}
	return out.Lang, nil
}
