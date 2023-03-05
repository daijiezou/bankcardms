package glog

import (
	"io"
)

type Option struct {
	Output    io.Writer
	Level     Level
	Formatter string
}

type OptionFn func(*Option)

var levelMap = map[string]Level{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"fatal": FatalLevel,
}

func getLevel(level string) Level {
	lv := InfoLevel
	if lve, ok := levelMap[level]; ok {
		lv = lve
	}

	return lv
}

func SetLevelOpt(lv string) OptionFn {
	return func(option *Option) {
		option.Level = getLevel(lv)
	}
}

func SetFormatterOpt(formatter string) OptionFn {
	return func(option *Option) {
		if formatter == "" {
			formatter = TextFormatter
		}
		option.Formatter = formatter
	}
}
