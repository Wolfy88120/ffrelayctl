package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	domainAddressesPath = APIBasePath + "domainaddresses/"
)

func (c *Client) ListDomainAddresses() ([]DomainAddress, error) {
	resp, err := c.Get(domainAddressesPath)
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

	var addresses []DomainAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (c *Client) GetDomainAddress(id int) (*DomainAddress, error) {
	path := fmt.Sprintf("%s%d/", domainAddressesPath, id)
	resp, err := c.Get(path)
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

	var address DomainAddress
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (c *Client) CreateDomainAddress(req CreateDomainAddressRequest) (*DomainAddress, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Post(domainAddressesPath, strings.NewReader(string(jsonBody)))
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

	var address DomainAddress
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (c *Client) UpdateDomainAddress(id int, req UpdateDomainAddressRequest) (*DomainAddress, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	path := fmt.Sprintf("%s%d/", domainAddressesPath, id)
	resp, err := c.Patch(path, strings.NewReader(string(jsonBody)))
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

	var address DomainAddress
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (c *Client) DeleteDomainAddress(id int) error {
	path := fmt.Sprintf("%s%d/", domainAddressesPath, id)
	resp, err := c.Delete(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	return nil
}
