package gerr

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

var ErrDataNotFound = errors.New("data not found")

type GeminiError struct {
	Code      ResCode
	Msg       string
	CauseCode int
	CauseMsg  string
	Metadata  map[string]string
	Err       error
}

func (g *GeminiError) Unwrap() error {
	return g.Err
}

func (g *GeminiError) Error() string {
	if g.Msg != "" {
		return g.Msg
	}
	return g.Code.Msg()
}

// LogString
// not recommend
func (g *GeminiError) LogString(c *gin.Context, err error, msg string) string {
	return fmt.Sprintf("traceId:%s，"+
		"%s,"+
		"code:%d,msg:%s,causeCode:%d,causeMsg:%s\nstack:%+v",
		trace.SpanContextFromContext(c.Request.Context()).TraceID().String(), msg, g.Code, g.Msg, g.CauseCode, g.CauseMsg, err)
}

func (g *GeminiError) ErrorWithMsg(c *gin.Context, err error, msg string) string {
	return fmt.Sprintf("traceId:%s,extraMsg:%s,code:%d,msg:%s,causeCode:%d,causeMsg:%s\nstack:%+v",
		trace.SpanContextFromContext(c.Request.Context()).TraceID().String(), msg, g.Code, g.Msg, g.CauseCode, g.CauseMsg, err)
}

func (g *GeminiError) Warning(c *gin.Context, err error) string {
	return fmt.Sprintf("traceId:%s，err:%s,code:%d,msg:%s,causeCode:%d,causeMsg:%s",
		trace.SpanContextFromContext(c.Request.Context()).TraceID().String(), err.Error(), g.Code, g.Msg, g.CauseCode, g.CauseMsg)
}

func (g *GeminiError) WarningWithMsg(c *gin.Context, err error, msg string) string {
	return fmt.Sprintf("traceId:%s，extraMsg:%s,err:%s,code:%d,msg:%s,causeCode:%d,causeMsg:%s",
		trace.SpanContextFromContext(c.Request.Context()).TraceID().String(), msg, err.Error(), g.Code, g.Msg, g.CauseCode, g.CauseMsg)
}

func Info(c *gin.Context, msg string, md interface{}) string {
	return fmt.Sprintf("traceId:%s,msg:%s,md:%+v", trace.SpanContextFromContext(c.Request.Context()).TraceID().String(), msg, md)
}

func New(code ResCode, causeCode int, msg, causeMsg string) *GeminiError {
	return &GeminiError{
		Code:      code,
		Msg:       msg,
		CauseCode: causeCode,
		CauseMsg:  causeMsg,
	}
}

func FromError(err error) *GeminiError {
	if err == nil {
		return nil
	}
	var gErr = new(GeminiError)
	if errors.As(err, &gErr) {
		return gErr
	}
	return &GeminiError{
		Code:     ErrCodeServerBusy,
		CauseMsg: err.Error(),
	}
}

// TransErrorToGErr
// not recommend
func TransErrorToGErr(err error) *GeminiError {
	if err == nil {
		return nil
	}
	var gErr = new(GeminiError)
	if errors.As(err, &gErr) {
		return gErr
	}
	return &GeminiError{
		Code:     ErrCodeServerBusy,
		CauseMsg: err.Error(),
	}
}
