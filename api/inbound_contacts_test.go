package api

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestClient_ListInboundContacts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name           string
		mockResponse   string
		mockStatusCode int
		wantErr        bool
		wantCount      int
	}{
		{
			name: "successful list inbound contacts",
			mockResponse: `[
				{
					"id": 1,
					"relay_number": 1,
					"inbound_number": "+18001234567",
					"last_inbound_date": "2026-01-01T19:20:37.224184Z",
					"last_inbound_type": "call",
					"num_calls": 1,
					"num_calls_blocked": 0,
					"last_call_date": "2026-01-02T19:20:37.224184Z",
					"num_texts": 0,
					"num_texts_blocked": 0,
					"last_text_date": null,
					"blocked": false
				},
				{
					"id": 2,
					"relay_number": 1,
					"inbound_number": "+18001234569",
					"last_inbound_date": "2026-01-02T17:12:50.215331Z",
					"last_inbound_type": "text",
					"num_calls": 0,
					"num_calls_blocked": 0,
					"last_call_date": null,
					"num_texts": 1,
					"num_texts_blocked": 0,
					"last_text_date": "2026-01-02T17:12:50.215331Z",
					"blocked": false
				}
			]`,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      2,
		},
		{
			name:           "empty list",
			mockResponse:   `[]`,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      0,
		},
		{
			name: "null last_call_date and last_text_date fields",
			mockResponse: `[
				{
					"id": 1,
					"relay_number": 1,
					"inbound_number": "+18001234560",
					"last_inbound_date": "2026-01-07T19:20:37.224184Z",
					"last_inbound_type": "call",
					"num_calls": 1,
					"num_calls_blocked": 0,
					"last_call_date": null,
					"num_texts": 0,
					"num_texts_blocked": 0,
					"last_text_date": null,
					"blocked": false
				}
			]`,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
			wantCount:      1,
		},
		{
			name:           "unauthorized",
			mockResponse:   `{"detail": "Invalid token"}`,
			mockStatusCode: http.StatusUnauthorized,
			wantErr:        true,
		},
		{
			name:           "not found - no relay number",
			mockResponse:   `{"detail": "No RelayNumber matches the given query."}`,
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
		{
			name:           "forbidden - no phone subscription",
			mockResponse:   `{"detail": "Phone subscription required"}`,
			mockStatusCode: http.StatusForbidden,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Reset()

			client := NewClient("test")

			httpmock.RegisterResponder("GET", DefaultBaseURL+"/api/v1/inboundcontact/",
				httpmock.NewStringResponder(tt.mockStatusCode, tt.mockResponse))

			contacts, err := client.ListInboundContacts()

			if (err != nil) != tt.wantErr {
				t.Errorf("ListInboundContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(contacts) != tt.wantCount {
					t.Errorf("ListInboundContacts() got %d contacts, want %d", len(contacts), tt.wantCount)
				}

				if tt.wantCount > 0 {
					if contacts[0].InboundNumber == "" {
						t.Error("ListInboundContacts() first contact has empty InboundNumber field")
					}
					if contacts[0].LastInboundType == "" {
						t.Error("ListInboundContacts() first contact has empty LastInboundType field")
					}
				}

				if tt.name == "null last_call_date and last_text_date fields" && tt.wantCount > 0 {
					if contacts[0].LastCallDate != nil {
						t.Errorf("ListInboundContacts() expected nil LastCallDate, got %v", *contacts[0].LastCallDate)
					}
					if contacts[0].LastTextDate != nil {
						t.Errorf("ListInboundContacts() expected nil LastTextDate, got %v", *contacts[0].LastTextDate)
					}
				}
			}
		})
	}
}

func TestClient_ListInboundContacts_InvalidJSON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client := NewClient("test")

	httpmock.RegisterResponder("GET", DefaultBaseURL+"/api/v1/inboundcontact/",
		httpmock.NewStringResponder(http.StatusOK, `invalid json`))

	_, err := client.ListInboundContacts()
	if err == nil {
		t.Error("ListInboundContacts() expected error for invalid JSON, got nil")
	}
}
