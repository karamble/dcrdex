// Package mexctypes provides type definitions for the MEXC exchange API.
package mexctypes

// ListenKeyResponse represents the response from the /api/v3/userDataStream endpoint.
type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
	Code      int    `json:"code,omitempty"`    // Error code if request fails
	Message   string `json:"message,omitempty"` // Error message if request fails
}
