package core

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
	"unicode/utf8"

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

// Configure a more terminal-like display with wider padding
func getTableWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 6, ' ', 0)
}

// Terminal UI colors
var (
	headerStyle   = color.New(color.Bold, color.FgHiWhite).SprintFunc()
	senderColor   = color.New(color.FgCyan)
	receiverColor = color.New(color.FgGreen)
	protoColors   = map[string]*color.Color{
		"ICMP": color.New(color.FgMagenta, color.Bold),
		"TCP":  color.New(color.FgBlue, color.Bold),
		"UDP":  color.New(color.FgYellow, color.Bold),
	}
)

func wrapText(text string, limit int) string {
	if utf8.RuneCountInString(text) <= limit {
		return text
	}
	runes := []rune(text)
	var lines []string
	for i := 0; i < len(runes); i += limit {
		end := i + limit
		if end > len(runes) {
			end = len(runes)
		}
		lines = append(lines, string(runes[i:end]))
	}
	return strings.Join(lines, "\n")
}

func LogNetworkEventIPv4(
	senderName, senderIP, senderLocation string,
	receiverName, receiverIP, receiverLocation string,
	protocol string,
) {
	protoColor, ok := protoColors[protocol]
	if !ok {
		protoColor = color.New(color.FgWhite)
	}

	w := getTableWriter()

	senderNameWrapped := wrapText(senderName, 30)
	receiverNameWrapped := wrapText(receiverName, 30)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		senderColor.Sprintf("%-30s", senderNameWrapped),
		senderColor.Sprintf("%-20s", senderIP),
		receiverColor.Sprintf("%-30s", receiverNameWrapped),
		receiverColor.Sprintf("%-20s", receiverIP),
		protoColor.Sprintf("%-10s", protocol),
	)
	w.Flush()
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

	w := getTableWriter()
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		senderColor.Sprintf("%-30s", senderName),
		senderColor.Sprintf("%-45s", senderIP),
		receiverColor.Sprintf("%-30s", receiverName),
		receiverColor.Sprintf("%-45s", receiverIP),
		protoColor.Sprintf("%-10s", protocol),
	)
	w.Flush()
}


func LogNetworkEventHeaderIPv4() {
	w := getTableWriter()
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		headerStyle(fmt.Sprintf("%-30s", "SENDER_NAME")),
		headerStyle(fmt.Sprintf("%-45s", "SENDER_IP")),
		headerStyle(fmt.Sprintf("%-30s", "RECEIVER_NAME")),
		headerStyle(fmt.Sprintf("%-45s", "RECEIVER_IP")),
		headerStyle(fmt.Sprintf("%-10s", "PROTOCOL")),
	)
	w.Flush()
}


func LogNetworkEventHeaderIPv6() {
	w := getTableWriter()
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		headerStyle(fmt.Sprintf("%-30s", "SENDER_NAME")),
		headerStyle(fmt.Sprintf("%-45s", "SENDER_IP")),
		headerStyle(fmt.Sprintf("%-30s", "RECEIVER_NAME")),
		headerStyle(fmt.Sprintf("%-45s", "RECEIVER_IP")),
		headerStyle(fmt.Sprintf("%-10s", "PROTOCOL")),
	)
	w.Flush()
}

