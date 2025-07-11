package core

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)



type Logger struct{ debug bool }

var log = Logger{debug: false}

func SetDebug(enabled bool) { log.debug = enabled }

func logMessage(level string, levelColor *color.Color, format string, a ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stdout, "%s [%s] %s\n", levelColor.Sprintf("[%s]", level), timestamp, msg)
}

func Info(format string, a ...interface{})  { logMessage("INFO", color.New(color.FgGreen), format, a...) }
func Warn(format string, a ...interface{})  { logMessage("WARN", color.New(color.FgYellow), format, a...) }
func Error(format string, a ...interface{}) { logMessage("ERROR", color.New(color.FgRed), format, a...) }
func Debug(format string, a ...interface{}) {
	if log.debug {
		logMessage("DEBUG", color.New(color.FgCyan), format, a...)
	}
}



type tblCfg struct {
	nameW int 
	ipW   int 
}

var (
	ipv4Cfg = tblCfg{nameW: 18, ipW: 15} 
	ipv6Cfg = tblCfg{nameW: 18, ipW: 42} 
)

var (
	senderColor   = color.New(color.FgCyan)
	receiverColor = color.New(color.FgGreen)
	protoColors   = map[string]*color.Color{
		"TCP":  color.New(color.FgBlue, color.Bold),
		"UDP":  color.New(color.FgYellow, color.Bold),
		"ICMP": color.New(color.FgMagenta, color.Bold),
	}
	headerStyle = color.New(color.Bold, color.FgHiWhite).SprintFunc()
)



var (
	headerPrintedIPv4 bool
	headerPrintedIPv6 bool
	headerMu          sync.Mutex
)

func printHeader(isIPv6 bool) {
	headerMu.Lock()
	defer headerMu.Unlock()

	switch {
	case !isIPv6 && headerPrintedIPv4:
		return
	case isIPv6 && headerPrintedIPv6:
		return
	}

	cfg := ipv4Cfg
	if isIPv6 {
		cfg = ipv6Cfg
	}

	fmt.Printf("%-*s %-*s %-*s %-*s %-8s\n",
		cfg.nameW, headerStyle("SENDER_NAME"),
		cfg.ipW, headerStyle("SENDER_IP"),
		cfg.nameW, headerStyle("RECEIVER_NAME"),
		cfg.ipW, headerStyle("RECEIVER_IP"),
		headerStyle("PROTOCOL"),
	)
	fmt.Println(strings.Repeat("-", cfg.nameW*2+cfg.ipW*2+8+4)) // simple underline

	if isIPv6 {
		headerPrintedIPv6 = true
	} else {
		headerPrintedIPv4 = true
	}
}


func LogNetworkEventHeaderIPv4() { printHeader(false) }
func LogNetworkEventHeaderIPv6() { printHeader(true) }



func printRow(senderName, senderIP, receiverName, receiverIP, proto string, isIPv6 bool) {
	cfg := ipv4Cfg
	if isIPv6 {
		cfg = ipv6Cfg
	}

	// choose protocol colour
	pc, ok := protoColors[proto]
	if !ok {
		pc = color.New(color.FgWhite)
	}

	fmt.Printf("%-*s %-*s %-*s %-*s %-8s\n",
		cfg.nameW, senderColor.Sprintf(senderName),
		cfg.ipW, senderColor.Sprintf(senderIP),
		cfg.nameW, receiverColor.Sprintf(receiverName),
		cfg.ipW, receiverColor.Sprintf(receiverIP),
		pc.Sprintf(proto),
	)
}


func LogNetworkEventIPv4(senderName, senderIP, _ string,
	receiverName, receiverIP, _ string,
	proto string,
) {
	printRow(senderName, senderIP, receiverName, receiverIP, proto, false)
}

func LogNetworkEventIPv6(senderName, senderIP, _ string,
	receiverName, receiverIP, _ string,
	proto string,
) {
	printRow(senderName, senderIP, receiverName, receiverIP, proto, true)
}
