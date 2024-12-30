package config

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigService(t *testing.T) {
	tests := []struct {
		name            string
		httpClient      *http.Client
		pubxId          string
		endpoint        string
		refreshInterval string
		wantErr         bool
	}{
		{
			name:            "valid configuration",
			httpClient:      &http.Client{},
			pubxId:         "testPublisher",
			endpoint:       "http://example.com",
			refreshInterval: "1m",
			wantErr:        false,
		},
		{
			name:            "invalid duration",
			httpClient:      &http.Client{},
			pubxId:         "testPublisher",
			endpoint:       "http://example.com",
			refreshInterval: "invalid",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configService, err := NewConfigService(tt.httpClient, tt.pubxId, tt.endpoint, tt.refreshInterval)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, configService)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, configService)
			}
		})
	}
}

func TestFetchConfig(t *testing.T) {
	tests := []struct {
		name       string
		serverResp http.HandlerFunc // Changed type to http.HandlerFunc
		wantConfig *Configuration
		wantErr    bool
	}{
		{
			name: "successful fetch",
			serverResp: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Explicit cast
				config := &Configuration{
					PublisherId:        "testPublisher",
					BufferInterval:     "30s",
					BufferSize:         "100",
					SamplingPercentage: 50,
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(config)
			}),
			wantConfig: &Configuration{
				PublisherId:        "testPublisher",
				BufferInterval:     "30s",
				BufferSize:         "100",
				SamplingPercentage: 50,
			},
			wantErr: false,
		},
		{
			name: "server error",
			serverResp: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // Explicit cast
				w.WriteHeader(http.StatusInternalServerError)
			}),
			wantConfig: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.serverResp) // No need for HandlerFunc here
			defer server.Close()

			endpointUrl, _ := url.Parse(server.URL)
			config, err := fetchConfig(server.Client(), endpointUrl)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, config)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantConfig, config)
			}
		})
	}
}

func TestConfigServiceImpl_Start(t *testing.T) {
	tests := []struct {
		name            string
		refreshInterval string
	}{
		{
			name:            "start and stop service",
			refreshInterval: "1m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := &http.Client{}
			pubxId := "testPublisher"
			endpoint := "http://example.com"

			configService, err := NewConfigService(httpClient, pubxId, endpoint, tt.refreshInterval)
			assert.NoError(t, err)

			configImpl := configService.(*ConfigServiceImpl)
			stop := make(chan struct{})
			done := make(chan struct{})

			configChan := configImpl.Start(stop)
			assert.NotNil(t, configChan)

			// Signal completion after a brief interval
			go func() {
				time.Sleep(10 * time.Millisecond)
				close(stop)
				done <- struct{}{}
			}()

			select {
			case <-done:
				// Success: service started and stopped
			case <-time.After(time.Second):
				t.Fatal("timeout waiting for service to stop")
			}
		})
	}
}

func TestConfigServiceImpl_IsSameAs(t *testing.T) {
	tests := []struct {
		name     string
		config1  *Configuration
		config2  *Configuration
		expected bool
	}{
		{
			name: "identical configs",
			config1: &Configuration{
				PublisherId:        "testPublisher",
				BufferInterval:     "30s",
				BufferSize:         "100",
				SamplingPercentage: 50,
			},
			config2: &Configuration{
				PublisherId:        "testPublisher",
				BufferInterval:     "30s",
				BufferSize:         "100",
				SamplingPercentage: 50,
			},
			expected: true,
		},
		{
			name: "different configs",
			config1: &Configuration{
				PublisherId:        "testPublisher",
				BufferInterval:     "30s",
				BufferSize:         "100",
				SamplingPercentage: 50,
			},
			config2: &Configuration{
				PublisherId:        "differentPublisher",
				BufferInterval:     "30s",
				BufferSize:         "100",
				SamplingPercentage: 50,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configService := &ConfigServiceImpl{}
			result := configService.IsSameAs(tt.config1, tt.config2)
			assert.Equal(t, tt.expected, result)
		})
	}
}