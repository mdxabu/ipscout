package core

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	debug bool
}

var log = Logger{debug: false}

func SetDebug(enabled bool) {
	log.debug = enabled
}

func logMessage(level string, levelColor *color.Color, format string, a ...interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(w, "%s\t%s\t%s\n", levelColor.Sprintf("[%s]", level), timestamp, msg)
	w.Flush()
}

func Info(format string, a ...interface{}) {
	logMessage("INFO", color.New(color.FgGreen), format, a...)
}

func Warn(format string, a ...interface{}) {
	logMessage("WARN", color.New(color.FgYellow), format, a...)
}

func Error(format string, a ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "%s\t%s\t%s\n", color.New(color.FgRed).Sprintf("[ERROR]"), timestamp, msg)
}

func Debug(format string, a ...interface{}) {
	if !log.debug {
		return
	}
	logMessage("DEBUG", color.New(color.FgCyan), format, a...)
}

var (
	senderColor   = color.New(color.FgCyan)
	receiverColor = color.New(color.FgGreen)
	protoColors   = map[string]*color.Color{
		"ICMP": color.New(color.FgMagenta),
		"TCP":  color.New(color.FgBlue),
		"UDP":  color.New(color.FgYellow),
	}
)

func LogNetworkEventIPv4(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	protoColor, ok := protoColors[protocol]
	if !ok {
		protoColor = color.New(color.FgWhite)
	}

	fmt.Printf("| %s | %s | %s | %s | %s | %s | %s |\n",
		senderColor.Sprintf("%-22s", senderName),
		senderColor.Sprintf("%-20s", senderIP),
		senderColor.Sprintf("%-38s", senderLocation),
		receiverColor.Sprintf("%-22s", receiverName),
		receiverColor.Sprintf("%-20s", receiverIP),
		receiverColor.Sprintf("%-38s", receiverLocation),
		protoColor.Sprintf("%-10s", protocol),
	)
}

func LogNetworkEventIPv6(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	protoColor, ok := protoColors[protocol]
	if !ok {
		protoColor = color.New(color.FgWhite)
	}

	fmt.Printf("| %s | %s | %s | %s | %s | %s | %s |\n",
		senderColor.Sprintf("%-22s", senderName),
		senderColor.Sprintf("%-46s", senderIP),
		senderColor.Sprintf("%-42s", senderLocation),
		receiverColor.Sprintf("%-22s", receiverName),
		receiverColor.Sprintf("%-46s", receiverIP),
		receiverColor.Sprintf("%-42s", receiverLocation),
		protoColor.Sprintf("%-10s", protocol),
	)
}

func LogNetworkEventHeaderIPv4() {
	headerStyle := color.New(color.Bold, color.FgHiWhite).SprintFunc()

	header := "| Sender Name            | Sender IP              | Sender Location                            | Receiver Name          | Receiver IP            | Receiver Location                          | Protocol   |"
	separator := strings.Repeat("-", len(header))

	fmt.Println(headerStyle(header))
	fmt.Println(separator)
}

func LogNetworkEventHeaderIPv6() {
	headerStyle := color.New(color.Bold, color.FgHiWhite).SprintFunc()

	header := "| Sender Name            | Sender IP                                      | Sender Location                              | Receiver Name          | Receiver IP                                    | Receiver Location                            | Protocol   |"
	separator := strings.Repeat("-", len(header))

	fmt.Println(headerStyle(header))
	fmt.Println(separator)
}
