package logger

import (
	"context"
	"fmt"
	"theitem/shared/gogen"
	"time"
)

func NewSimpleJSONLogger(appData gogen.ApplicationData) Logger {
	return &simpleJSONLoggerImpl{AppData: appData}
}

type jsonLogModel struct {
	AppName   string `json:"appName"`
	AppInstID string `json:"appInstID"`
	Start     string `json:"start"`
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Location  string `json:"location"`
	Time      string `json:"time"`
}

func newJSONLogModel(lg *simpleJSONLoggerImpl, flag, loc string, msg, trid any) string {

	if flag == "ERROR" {
		return toJsonString(jsonLogModel{
			AppName:   lg.AppData.AppName,
			AppInstID: lg.AppData.AppInstanceID,
			Start:     lg.AppData.StartTime,
			Severity:  flag,
			Message:   fmt.Sprintf("%v %v %v", trid, loc, msg),
			Location:  loc,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return toJsonString(jsonLogModel{
		AppName:   lg.AppData.AppName,
		AppInstID: lg.AppData.AppInstanceID,
		Start:     lg.AppData.StartTime,
		Severity:  flag,
		Message:   fmt.Sprintf("%v %v", trid, msg),
		Location:  loc,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	})
}

type simpleJSONLoggerImpl struct {
	AppData gogen.ApplicationData
}

func (l simpleJSONLoggerImpl) Info(ctx context.Context, message string, args ...any) {
	messageWithArgs := fmt.Sprintf(message, args...)
	l.printLog(ctx, "INFO", messageWithArgs)
}

func (l simpleJSONLoggerImpl) Error(ctx context.Context, message string, args ...any) {
	messageWithArgs := fmt.Sprintf(message, args...)
	l.printLog(ctx, "ERROR", messageWithArgs)
}

func (l simpleJSONLoggerImpl) printLog(ctx context.Context, flag string, data any) {
	traceID := GetTraceID(ctx)
	fmt.Printf("%-5s %s %-60v %s\n", flag, traceID, data, getFileLocationInfo(3))
	// fmt.Println(newJSONLogModel(&l, flag, getFileLocationInfo(3), data, traceID))
}
