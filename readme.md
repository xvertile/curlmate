# CurlMate

Tired of manually forming a proxy string every time you use `curl`? **CurlMate** simplifies the process by converting your proxy string into a `curl` command and executing it with a single command.

## why

I am so tired of manually forming a proxy string every time I use `curl` to make a request.

## How It Works

Instead of repeatedly crafting your proxy strings for `curl`, just pass your proxy details to CurlMate, and it does the rest!
## Installation

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri "https://d2qy48d1h0e7ws.cloudfront.net/curlmate/windows/curlmate.exe" -OutFile "$env:USERPROFILE\curlmate.exe"; $env:PATH += ";$env:USERPROFILE"; [System.Environment]::SetEnvironmentVariable("Path", "$env:PATH;$env:USERPROFILE", [System.EnvironmentVariableTarget]::User)
```

### macOS
```bash
sudo curl -o /usr/local/bin/curlmate "https://d2qy48d1h0e7ws.cloudfront.net/curlmate/macos/curlmate" && sudo chmod +x /usr/local/bin/curlmate && echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bash_profile && source ~/.bash_profile
```

### Linux
```bash
curl -o /usr/local/bin/curlmate "https://d2qy48d1h0e7ws.cloudfront.net/curlmate/linux/curlmate" && chmod +x /usr/local/bin/curlmate && echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc && source ~/.bashrc
```

## Usage

CurlMate supports multiple proxy formats and types, making it flexible for different use cases.

### Basic Usage

```bash
curlmate <proxy-string>
```

### Supported Formats

CurlMate automatically detects and parses various proxy string formats:

#### 1. Simple proxy without authentication
```bash
curlmate proxy.example.com:8080
```

#### 2. Proxy with authentication (original format)
```bash
curlmate proxy.example.com:8080:username:password
```

#### 3. Proxy with @ symbol (pre-formatted)
```bash
curlmate username:password@proxy.example.com:8080
```

#### 4. Proxy with protocol specified
```bash
curlmate http://proxy.example.com:8080
curlmate https://username:password@proxy.example.com:8443
```

#### 5. SOCKS proxies
```bash
curlmate socks4://proxy.example.com:1080
curlmate socks5://username:password@proxy.example.com:1080
```

### Supported Proxy Types

- **HTTP** (default)
- **HTTPS**
- **SOCKS4**
- **SOCKS5**

### Example Output

When you run CurlMate, it will form and execute the curl command:

```bash
curl --proxy http://username:password@proxy.example.com:8080 https://ipinfo.io
```

The command will display your IP information through the specified proxy.

## License

This project is open-source and available under the MIT License. Feel free to use, modify, and share!