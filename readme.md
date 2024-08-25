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

Simply run the following command with your proxy details:

```bash
curlmate <proxy-host>:<proxy-port>:<username>:<password>
```

It will then form the `curl` command and execute it for you. Example:
    
```bash
curl --proxy http://<username>:<password>@<proxy-host>:<proxy-port> https://ipinfo.io
```

## License

This project is open-source and available under the MIT License. Feel free to use, modify, and share!