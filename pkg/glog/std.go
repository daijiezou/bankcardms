package glog

var std = newDefaultLogger()

var (
	WithCtx  = std.WithCtx
	SetLevel = std.SetLevel

	Debug     = std.Debug
	Debugw    = std.Debugw
	Debugf    = std.Debugf
	CtxDebug  = std.CtxDebug
	CtxDebugw = std.CtxDebugw
	CtxDebugf = std.CtxDebugf

	Info     = std.Info
	Infow    = std.Infow
	Infof    = std.Infof
	CtxInfo  = std.CtxInfo
	CtxInfow = std.CtxInfow
	CtxInfof = std.CtxInfof

	Warn     = std.Warn
	Warnw    = std.Warnw
	Warnf    = std.Warnf
	CtxWarn  = std.CtxWarn
	CtxWarnw = std.CtxWarnw
	CtxWarnf = std.CtxWarnf

	Error     = std.Error
	Errorw    = std.Errorw
	Errorf    = std.Errorf
	CtxError  = std.CtxError
	CtxErrorw = std.CtxErrorw
	CtxErrorf = std.CtxErrorf

	Fatal  = std.Fatal
	Fatalw = std.Fatalw
	Fatalf = std.Fatalf

	InfoWarnw  = std.InfoWarnw
	InfoErrorw = std.InfoErrorw
	InfoFatalw = std.InfoFatalw

	DebugWarnw  = std.DebugWarnw
	DebugErrorw = std.DebugErrorw
	DebugFatalw = std.DebugFatalw
)
