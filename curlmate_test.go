package main

import (
	"testing"
)

func TestParseProxy(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ProxyInfo
		wantErr  bool
	}{
		{
			name:  "Simple host:port without auth",
			input: "proxy.example.com:8080",
			expected: ProxyInfo{
				Protocol: "http",
				Host:     "proxy.example.com",
				Port:     "8080",
				Username: "",
				Password: "",
			},
			wantErr: false,
		},
		{
			name:  "Format with @ symbol",
			input: "user:pass@proxy.example.com:8080",
			expected: ProxyInfo{
				Protocol: "http",
				Host:     "proxy.example.com",
				Port:     "8080",
				Username: "user",
				Password: "pass",
			},
			wantErr: false,
		},
		{
			name:  "Original format host:port:user:pass",
			input: "proxy.example.com:8080:user:pass",
			expected: ProxyInfo{
				Protocol: "http",
				Host:     "proxy.example.com",
				Port:     "8080",
				Username: "user",
				Password: "pass",
			},
			wantErr: false,
		},
		{
			name:  "HTTP protocol specified",
			input: "http://proxy.example.com:8080",
			expected: ProxyInfo{
				Protocol: "http",
				Host:     "proxy.example.com",
				Port:     "8080",
				Username: "",
				Password: "",
			},
			wantErr: false,
		},
		{
			name:  "HTTPS protocol with auth",
			input: "https://user:pass@proxy.example.com:8443",
			expected: ProxyInfo{
				Protocol: "https",
				Host:     "proxy.example.com",
				Port:     "8443",
				Username: "user",
				Password: "pass",
			},
			wantErr: false,
		},
		{
			name:  "SOCKS5 protocol",
			input: "socks5://user:pass@proxy.example.com:1080",
			expected: ProxyInfo{
				Protocol: "socks5",
				Host:     "proxy.example.com",
				Port:     "1080",
				Username: "user",
				Password: "pass",
			},
			wantErr: false,
		},
		{
			name:  "SOCKS4 without auth",
			input: "socks4://proxy.example.com:1080",
			expected: ProxyInfo{
				Protocol: "socks4",
				Host:     "proxy.example.com",
				Port:     "1080",
				Username: "",
				Password: "",
			},
			wantErr: false,
		},
		{
			name:    "Invalid port number",
			input:   "proxy.example.com:99999",
			wantErr: true,
		},
		{
			name:    "Invalid format too few components",
			input:   "proxy.example.com",
			wantErr: true,
		},
		{
			name:    "Invalid format too many components",
			input:   "a:b:c:d:e",
			wantErr: true,
		},
		{
			name:    "Unsupported protocol",
			input:   "ftp://proxy.example.com:8080",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseProxy(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("parseProxy() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("parseProxy() unexpected error: %v", err)
				return
			}

			if result.Protocol != tt.expected.Protocol {
				t.Errorf("Protocol = %v, want %v", result.Protocol, tt.expected.Protocol)
			}
			if result.Host != tt.expected.Host {
				t.Errorf("Host = %v, want %v", result.Host, tt.expected.Host)
			}
			if result.Port != tt.expected.Port {
				t.Errorf("Port = %v, want %v", result.Port, tt.expected.Port)
			}
			if result.Username != tt.expected.Username {
				t.Errorf("Username = %v, want %v", result.Username, tt.expected.Username)
			}
			if result.Password != tt.expected.Password {
				t.Errorf("Password = %v, want %v", result.Password, tt.expected.Password)
			}
		})
	}
}

func TestBuildCurlCommand(t *testing.T) {
	tests := []struct {
		name     string
		proxy    ProxyInfo
		expected string
	}{
		{
			name: "HTTP proxy with auth",
			proxy: ProxyInfo{
				Protocol: "http",
				Host:     "proxy.example.com",
				Port:     "8080",
				Username: "user",
				Password: "pass",
			},
			expected: "curl --proxy http://user:pass@proxy.example.com:8080 https://ipinfo.io",
		},
		{
			name: "SOCKS5 without auth",
			proxy: ProxyInfo{
				Protocol: "socks5",
				Host:     "proxy.example.com",
				Port:     "1080",
				Username: "",
				Password: "",
			},
			expected: "curl --proxy socks5://proxy.example.com:1080 https://ipinfo.io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildCurlCommand(&tt.proxy)
			if result != tt.expected {
				t.Errorf("buildCurlCommand() = %v, want %v", result, tt.expected)
			}
		})
	}
}
