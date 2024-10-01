package log

import "fmt"

type Logger struct {
	lastLog string
	logging bool
}

func (this *Logger) TurnLoggingOff() {
	this.logging = false
}

func (this *Logger) LogMessage(message string) {
	if !this.logging {
		return
	}

	this.lastLog = message
}

func (this *Logger) LogMessageComp(format string, data ...any) {
	if !this.logging {
		return
	}

	this.lastLog = fmt.Sprintf(format, data...)
}

func (this *Logger) LogMessageComp2(format string, data ...any) {
	if !this.logging {
		return
	}

	if len(data) == 0 {
		this.lastLog = format
		return
	}

	this.lastLog = fmt.Sprintf(format, data...)
}
