package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type ProxyInfo struct {
	Protocol string
	Host     string
	Port     string
	Username string
	Password string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: curlmate <proxy-string>")
		fmt.Println("\nSupported formats:")
		fmt.Println("  host:port")
		fmt.Println("  host:port:username:password")
		fmt.Println("  username:password@host:port")
		fmt.Println("  protocol://host:port")
		fmt.Println("  protocol://username:password@host:port")
		fmt.Println("\nSupported protocols: http, https, socks4, socks5")
		os.Exit(1)
	}

	proxyString := os.Args[1]
	proxy, err := parseProxy(proxyString)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	curlCmd := buildCurlCommand(proxy)
	fmt.Println(curlCmd)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", curlCmd)
	} else {
		cmd = exec.Command("sh", "-c", curlCmd)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing curl command:", err)
		os.Exit(1)
	}
}

func parseProxy(proxyString string) (*ProxyInfo, error) {
	proxy := &ProxyInfo{Protocol: "http"} // default protocol

	// Check if protocol is specified (e.g., http://, socks5://)
	if strings.Contains(proxyString, "://") {
		parts := strings.SplitN(proxyString, "://", 2)
		proxy.Protocol = strings.ToLower(parts[0])
		proxyString = parts[1]

		// Validate protocol
		validProtocols := map[string]bool{"http": true, "https": true, "socks4": true, "socks5": true, "socks4a": true}
		if !validProtocols[proxy.Protocol] {
			return nil, fmt.Errorf("unsupported protocol: %s (supported: http, https, socks4, socks5)", proxy.Protocol)
		}
	}

	// Check if credentials are included with @ symbol (username:password@host:port)
	if strings.Contains(proxyString, "@") {
		parts := strings.SplitN(proxyString, "@", 2)
		credentials := parts[0]
		hostPort := parts[1]

		// Parse credentials
		credParts := strings.SplitN(credentials, ":", 2)
		if len(credParts) == 2 {
			proxy.Username = credParts[0]
			proxy.Password = credParts[1]
		} else {
			return nil, fmt.Errorf("invalid credentials format in: %s", credentials)
		}

		// Parse host:port
		if err := parseHostPort(hostPort, proxy); err != nil {
			return nil, err
		}
	} else {
		// Check format without @ symbol
		components := strings.Split(proxyString, ":")

		switch len(components) {
		case 2:
			// Format: host:port (no authentication)
			proxy.Host = components[0]
			proxy.Port = components[1]
		case 4:
			// Format: host:port:username:password (original format)
			proxy.Host = components[0]
			proxy.Port = components[1]
			proxy.Username = components[2]
			proxy.Password = components[3]
		default:
			return nil, fmt.Errorf("invalid proxy format. See usage for supported formats")
		}
	}

	// Validate host
	if proxy.Host == "" {
		return nil, fmt.Errorf("proxy host cannot be empty")
	}

	// Validate port
	if proxy.Port == "" {
		return nil, fmt.Errorf("proxy port cannot be empty")
	}
	port, err := strconv.Atoi(proxy.Port)
	if err != nil || port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid port number: %s (must be 1-65535)", proxy.Port)
	}

	return proxy, nil
}

func parseHostPort(hostPort string, proxy *ProxyInfo) error {
	parts := strings.SplitN(hostPort, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid host:port format: %s", hostPort)
	}
	proxy.Host = parts[0]
	proxy.Port = parts[1]
	return nil
}

func buildCurlCommand(proxy *ProxyInfo) string {
	var proxyURL string

	// Build the proxy URL based on protocol and credentials
	if proxy.Username != "" && proxy.Password != "" {
		proxyURL = fmt.Sprintf("%s://%s:%s@%s:%s", proxy.Protocol, proxy.Username, proxy.Password, proxy.Host, proxy.Port)
	} else {
		proxyURL = fmt.Sprintf("%s://%s:%s", proxy.Protocol, proxy.Host, proxy.Port)
	}

	return fmt.Sprintf("curl --proxy %s https://ipinfo.io", proxyURL)
}
