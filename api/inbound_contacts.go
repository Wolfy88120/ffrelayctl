package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	inboundContactsPath = APIBasePath + "inboundcontact/"
)

func (c *Client) ListInboundContacts() ([]InboundContact, error) {
	resp, err := c.Get(inboundContactsPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	var contacts []InboundContact
	if err := json.Unmarshal(body, &contacts); err != nil {
		return nil, err
	}

	return contacts, nil
}
