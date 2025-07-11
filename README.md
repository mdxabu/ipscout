# IPScout
[![GoDoc](https://godoc.org/github.com/mxabu/ipscout?status.svg)](https://godoc.org/github.com/mdxabu/ipscout)


A network monitoring and packet analysis tool for real-time IP traffic inspection.


## Prerequisites

### Windows Users - Npcap Installation Required

IPScout requires **Npcap** (Nmap's packet capture library) to capture network packets on Windows.

1. **Download Npcap:**
   - Visit: https://npcap.com/#download
   - Download the latest stable version
   - **Important:** During installation, make sure to check "Install Npcap in WinPcap API-compatible Mode"

2. **Installation Steps:**
   ```
   1. Run the Npcap installer as Administrator
   2. Accept the license agreement
   3. ✅ Check "Install Npcap in WinPcap API-compatible Mode"
   4. ✅ Check "Support raw 802.11 traffic" (optional)
   5. Complete the installation
   6. Restart your computer
   ```

### Linux/macOS Users
- Most distributions come with libpcap pre-installed
- If needed: `sudo apt-get install libpcap-dev` (Ubuntu/Debian) or `brew install libpcap` (macOS)

## Installation

1. **Download IPScout:**
   - Download the latest release from the releases page
   - Or build from source (requires Go 1.19+)

2. **Download and Run:**
   ```bash
   go install github.com/mdxabu/ipscout
   
   ipscout start --monitor --<ipv4/ipv6>
   ```





## Usage

### Basic Usage

```bash

# Sniffing IPv4
ipscout start --monitor --ipv4

# Sniffing IPv6 
ipscout start --monitor --ipv6
```



## Available Commands

- `ipscout start` - Start network monitoring
- `ipscout version` - Show version information
- `ipscout help` - Show help information

For detailed command options, see [commands.md](docs/commands.md)

## Troubleshooting

### Common Issues

**"No suitable device found" error:**
- Ensure Npcap is properly installed (Windows)
- Run as Administrator/sudo
- Check if network interfaces are available

**Permission denied:**
- Run as Administrator (Windows) or with sudo (Linux/macOS)
- Ensure your user has permission to capture packets

**High CPU usage:**
- Use more specific packet filters
- Reduce the number of monitored interfaces



## Building from Source

```bash
git clone https://github.com/mdxabu/ipscout.git
cd ipscout
go mod download
go build -o ipscout
go install .
```


## Contributing

[CONTRIBUTING.md]
