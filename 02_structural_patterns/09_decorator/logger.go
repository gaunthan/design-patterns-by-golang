package decorator

import (
	"fmt"
)

type Logger interface {
	Log(str string)
}

type ScreenLogger struct{}

func (p *ScreenLogger) Log(str string) {
	fmt.Print(str)
}

func NewScreenLogger() *ScreenLogger {
	return &ScreenLogger{}
}

type LoggerDecorator struct {
	logger Logger
}

func (p *LoggerDecorator) Log(str string) {
	p.logger.Log(str)
}

type DateLoggerDecorator struct {
	LoggerDecorator
}

func (p *DateLoggerDecorator) Log(str string) {
	p.logger.Log(fmt.Sprintf("[%v] %s", "June 18, 2021", str))
}

func NewDateLoggerDecorator(logger Logger) *DateLoggerDecorator {
	return &DateLoggerDecorator{
		LoggerDecorator: LoggerDecorator{
			logger: logger,
		},
	}
}

type LineLoggerDecorator struct {
	LoggerDecorator
}

func (p *LineLoggerDecorator) Log(str string) {
	p.logger.Log(fmt.Sprintf("%s\n", str))
}

func NewLineLoggerDecorator(logger Logger) *LineLoggerDecorator {
	return &LineLoggerDecorator{
		LoggerDecorator: LoggerDecorator{
			logger: logger,
		},
	}
}
