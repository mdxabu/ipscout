# Available Commands

## ipscout start

Start network packet monitoring and analysis.

### Usage
```bash
ipscout start [flags]
```

### Flags
- `--monitor` - Monitor the network (required to start packet sniffing)
- `--ipv4` - Output the IPv4 data (mutually exclusive with --ipv6)
- `--ipv6` - Output the IPv6 data (mutually exclusive with --ipv4)
- `--srcip <ip>` - Filter packets by source IP address

### Examples
```bash
# Start monitoring IPv4 traffic
ipscout start --monitor --ipv4

# Start monitoring IPv6 traffic
ipscout start --monitor --ipv6

# Monitor with source IP filter
ipscout start --monitor --ipv4 --srcip "192.168.1.1"

```



## ipscout init

Initialize the ipscout configuration file with Wi-Fi IP detection.

### Usage
```bash
ipscout init
```

### Description
- Creates `ipscoutconfig.yaml` configuration file
- Automatically detects Wi-Fi interface
- Extracts IPv4 and IPv6 addresses
- Will not overwrite existing configuration files

## ipscout help

Show help information for commands.

```bash
ipscout help
ipscout help start
ipscout help version
ipscout help init
```


