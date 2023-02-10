package glog

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	JsonFormatter = "json"
	TextFormatter = "text"
)

type loggerKey struct{}

type Level = zapcore.Level

const (
	DebugLevel = zap.DebugLevel
	InfoLevel  = zap.InfoLevel
	WarnLevel  = zap.WarnLevel
	ErrorLevel = zap.ErrorLevel
	FatalLevel = zap.FatalLevel
)

var (
	port = ":9090"
)

type Logger struct {
	logger      *zap.Logger
	level       zap.AtomicLevel
	sugarLogger *zap.SugaredLogger
	opt         *Option
	zapConfig   zap.Config
	enableColor bool
}

var defaultOption = &Option{
	Output:    os.Stdout,
	Level:     DebugLevel,
	Formatter: TextFormatter,
}

func (l *Logger) defaultEncoderConfig() zapcore.EncoderConfig {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	encoderConf := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	if l.enableColor {
		encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	return encoderConf
}

// 创建默认日志记录器
func newDefaultLogger() *Logger {
	return NewLogger(SetLevelOpt("info"), SetFormatterOpt(TextFormatter))
}

func NewLogger(optionFn ...OptionFn) *Logger {
	l := &Logger{}

	// default
	l.opt = &Option{
		Level:     zapcore.InfoLevel,
		Formatter: TextFormatter,
		Output:    os.Stdout,
	}

	for _, o := range optionFn {
		o(l.opt)
	}

	l.init()

	return l
}

func (l *Logger) init() {
	var (
		encoderCfg = l.defaultEncoderConfig()
		logEncoder = zapcore.NewConsoleEncoder(encoderCfg)
	)

	if l.opt.Formatter == JsonFormatter {
		logEncoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	l.level = zap.NewAtomicLevelAt(l.opt.Level)

	//http.HandleFunc("/debug/setLevel", l.level.ServeHTTP)
	//go func() {
	//	if err := http.ListenAndServe(port, nil); err != nil {
	//		panic(err)
	//	}
	//}()

	core := zapcore.NewCore(
		logEncoder,
		zapcore.AddSync(l.opt.Output),
		l.level,
	)

	// 初始化 zap Logger
	zapLogger := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	)

	l.logger = zapLogger
	l.sugarLogger = zapLogger.Sugar()
	defer l.logger.Sync()
}

func WithContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *Logger {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return std
	}
	return logger.(*Logger)
}

func (l *Logger) WithCtx(ctx context.Context) *Logger {
	traceId := ""
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		traceId = span.TraceID().String()
	}

	return l.With(String("traceId", traceId))
}

func (l *Logger) clone() *Logger {
	copy := *l
	return &copy
}

func (l *Logger) SetLevel(level string) {
	lv := getLevel(level)
	l.level.SetLevel(lv)
}

func (l *Logger) SetFormatter(formatter string) {
	l.opt.Formatter = formatter
}

func (l *Logger) GetFormatter() (formatter string) {
	return l.opt.Formatter
}

func (l *Logger) GetLevel() (lv string) {
	return l.level.String()
}

func sprintf(template string, args ...interface{}) string {
	msg := template
	if msg == "" && len(args) > 0 {
		msg = fmt.Sprint(args...)
	} else if msg != "" && len(args) > 0 {
		msg = fmt.Sprintf(template, args...)
	}
	return msg
}

func normalizeMessage(msg string) string {
	return fmt.Sprintf("%-32s", msg)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugw(sprintf(template, args...))
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Infow(msg, keysAndValues...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(sprintf(template, args...))
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Warnw(msg, keysAndValues...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(sprintf(template, args...))
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(sprintf(template, args...))
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(sprintf(template, args...))
}

func panicDetail(msg string, fields ...Field) {
	enc := zapcore.NewMapObjectEncoder()
	for _, field := range fields {
		field.AddTo(enc)
	}

	// 控制台输出
	fmt.Printf("panic: \n    msg: %s\n", msg)
	for key, val := range enc.Fields {
		fmt.Printf("    %s: %s\n", key, fmt.Sprintf("%+v", val))
	}
}

func (l *Logger) With(fields ...Field) *Logger {
	logger := l.logger.With(fields...)
	return &Logger{
		logger:      logger,
		sugarLogger: logger.Sugar(),
	}
}

func (l *Logger) CtxDebug(ctx context.Context, msg string, fields ...Field) {
	l.logger.Debug(msg, logFields(ctx, fields)...)
}

func (l *Logger) CtxDebugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Debugw(msg, logKVs(ctx, keysAndValues)...)
}

func (l *Logger) CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	l.WithCtx(ctx).sugarLogger.Debugf(template, args...)
}

func (l *Logger) CtxInfo(ctx context.Context, msg string, fields ...Field) {
	l.logger.Info(msg, logFields(ctx, fields)...)
}

func (l *Logger) CtxInfow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Infow(msg, logKVs(ctx, keysAndValues)...)
}

func (l *Logger) CtxInfof(ctx context.Context, template string, args ...interface{}) {
	l.WithCtx(ctx).sugarLogger.Infof(template, args...)
}

func (l *Logger) CtxWarn(ctx context.Context, msg string, fields ...Field) {
	l.logger.Warn(msg, logFields(ctx, fields)...)
}

func (l *Logger) CtxWarnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Warnw(msg, logKVs(ctx, keysAndValues)...)
}

func (l *Logger) CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	l.WithCtx(ctx).sugarLogger.Warnf(template, args...)
}

func (l *Logger) CtxError(ctx context.Context, msg string, fields ...Field) {
	l.logger.Error(msg, logFields(ctx, fields)...)
}

func (l *Logger) CtxErrorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugarLogger.Errorw(msg, logKVs(ctx, keysAndValues)...)
}

func (l *Logger) CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	l.WithCtx(ctx).sugarLogger.Errorf(template, args...)
}

func logKVs(ctx context.Context, kvs []interface{}) []interface{} {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.HasTraceID() {
		return kvs
	}
	traceID := spanCtx.TraceID().String()
	kvs = append(kvs, "traceId", traceID)
	return kvs
}

func logFields(ctx context.Context, fields []Field) []Field {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.HasTraceID() {
		return fields
	}
	traceID := spanCtx.TraceID().String()
	fields = append(fields, zap.String("traceId", traceID))
	return fields
}
