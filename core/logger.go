package core

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	debug bool
}

var log = Logger{
	debug: false,
}

func SetDebug(enabled bool) {
	log.debug = enabled
}

func Info(format string, a ...interface{}) {
	prefix := color.New(color.FgGreen).Sprint("[INFO]")
	fmt.Printf("%s %s %s\n", prefix, time.Now().Format(time.RFC3339), fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	prefix := color.New(color.FgYellow).Sprint("[WARN]")
	fmt.Printf("%s %s %s\n", prefix, time.Now().Format(time.RFC3339), fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	prefix := color.New(color.FgRed).Sprint("[ERROR]")
	fmt.Fprintf(os.Stderr, "%s %s %s\n", prefix, time.Now().Format(time.RFC3339), fmt.Sprintf(format, a...))
}

func Debug(format string, a ...interface{}) {
	if !log.debug {
		return
	}
	prefix := color.New(color.FgCyan).Sprint("[DEBUG]")
	fmt.Printf("%s %s %s\n", prefix, time.Now().Format(time.RFC3339), fmt.Sprintf(format, a...))
}

// LogNetworkEvent prints a table row for a network event.
func LogNetworkEvent(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	var protoColored string
	switch protocol {
	case "ICMP":
		protoColored = color.New(color.FgCyan).Sprint(protocol)
	case "TCP":
		protoColored = color.New(color.FgGreen).Sprint(protocol)
	case "UDP":
		protoColored = color.New(color.FgYellow).Sprint(protocol)
	default:
		protoColored = color.New(color.FgMagenta).Sprint(protocol)
	}

	fmt.Printf(
		"| %-15s | %-15s | %-20s | %-15s | %-15s | %-20s | %-8s |\n",
		senderName, senderIP, senderLocation,
		receiverName, receiverIP, receiverLocation,
		protoColored,
	)
}

// LogNetworkEventHeader prints the table header for network events.
func LogNetworkEventHeader() {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Println(bold(
			"| Sender Name       | Sender IP                                               | Sender Location                          | Receiver Name      | Receiver IP                                               | Receiver Location                         | Protocol |",
))
fmt.Println("|-------------------|---------------------------------------------------------|------------------------------------------|--------------------|-----------------------------------------------------------|-------------------------------------------|----------|")
}

// LogNetworkEventIPv4 prints a table row for an IPv4 network event.
func LogNetworkEventIPv4(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	var protoColored string
	switch protocol {
	case "ICMP":
		protoColored = color.New(color.FgCyan).Sprint(protocol)
	case "TCP":
		protoColored = color.New(color.FgGreen).Sprint(protocol)
	case "UDP":
		protoColored = color.New(color.FgYellow).Sprint(protocol)
	default:
		protoColored = color.New(color.FgMagenta).Sprint(protocol)
	}

	fmt.Printf(
		"| %-18s | %-20s | %-36s | %-18s | %-20s | %-36s | %-8s |\n",
		senderName, senderIP, senderLocation,
		receiverName, receiverIP, receiverLocation,
		protoColored,
	)
}

// LogNetworkEventIPv6 prints a table row for an IPv6 network event.
func LogNetworkEventIPv6(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	var protoColored string
	switch protocol {
	case "ICMP":
		protoColored = color.New(color.FgCyan).Sprint(protocol)
	case "TCP":
		protoColored = color.New(color.FgGreen).Sprint(protocol)
	case "UDP":
		protoColored = color.New(color.FgYellow).Sprint(protocol)
	default:
		protoColored = color.New(color.FgMagenta).Sprint(protocol)
	}

	fmt.Printf(
		"| %-18s | %-42s | %-40s | %-18s | %-42s | %-40s | %-8s |\n",
		senderName, senderIP, senderLocation,
		receiverName, receiverIP, receiverLocation,
		protoColored,
	)
}

// LogNetworkEventHeaderIPv4 prints the table header for IPv4 network events.
func LogNetworkEventHeaderIPv4() {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Println(bold(
			"| Sender Name       | Sender IP                                               | Sender Location                          | Receiver Name      | Receiver IP                                               | Receiver Location                         | Protocol |",
))
fmt.Println("|-------------------|---------------------------------------------------------|------------------------------------------|--------------------|-----------------------------------------------------------|-------------------------------------------|----------|")
}


// LogNetworkEventHeaderIPv6 prints the table header for IPv6 network events.
func LogNetworkEventHeaderIPv6() {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Println(bold(
				"| Sender Name       | Sender IP                                               | Sender Location                          | Receiver Name      | Receiver IP                                               | Receiver Location                         | Protocol |",
	))
	fmt.Println("|-------------------|---------------------------------------------------------|------------------------------------------|--------------------|-----------------------------------------------------------|-------------------------------------------|----------|")
}
