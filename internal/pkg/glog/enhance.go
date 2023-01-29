package glog

import "go.uber.org/zap"

func (l *Logger) DebugWarnw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Warnw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Debugw(msg, keysAndValues...)
	}
}

func (l *Logger) DebugErrorw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Errorw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Debugw(msg, keysAndValues...)
	}
}

func (l *Logger) DebugFatalw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Fatalw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Debugw(msg, keysAndValues...)
	}
}

func (l *Logger) InfoWarnw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Warnw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Infow(msg, keysAndValues...)
	}
}

func (l *Logger) InfoErrorw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Errorw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Infow(msg, keysAndValues...)
	}
}

func (l *Logger) InfoFatalw(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		l.sugarLogger.With(zap.Error(err)).Fatalw(msg, keysAndValues...)
	} else {
		l.sugarLogger.Infow(msg, keysAndValues...)
	}
}
