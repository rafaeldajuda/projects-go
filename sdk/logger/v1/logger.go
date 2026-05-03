package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	lg    *zap.Logger
	sugar *zap.SugaredLogger
	std   = &Std{}
)

func init() {
	// Inicialização padrão para evitar nil pointer se alguém chamar o log antes do StartZapLog
	lg = zap.NewNop()
	sugar = lg.Sugar()
}

func SetStd(orgName, servName, servID string) {
	std.orgName = orgName
	std.servName = servName
	std.servID = servID
}

func StartZapLog() {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	fmt.Println("LOG_LEVEL:", logLevel)

	config := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,
		},
	}

	switch logLevel {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	var err error
	// lg, err = config.Build(zap.AddCallerSkip(1)) // AddCallerSkip faz o log mostrar onde VOCÊ chamou, não esta função wrapper
	lg, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("Erro ao inicializar logger: %v", err))
	}

	sugar = lg.Sugar()
}

// Sync limpa o buffer de logs. Deve ser chamado via 'defer logger.Sync()' no main.
func Sync() error {
	return lg.Sync()
}

// Template de log padrão

func log(messageID string, s string) string {
	return fmt.Sprintf(" | %s | %s | %s | %s | %s | %s",
		std.orgName, std.servName, std.servID, timeNow(), messageID, s)
}

// EBI | PACCINI | 0929 | abc123 | time |

// Funções para Logs Estruturados (Performance Alta)

func Debug(msg string, s string) {
	lg.Debug(log(msg, s))
}

func Info(msg string, s string) {
	lg.Info(log(msg, s))
}

func Warn(msg string, s string) {
	lg.Warn(log(msg, s))
}

func Error(msg string, s string) {
	lg.Error(log(msg, s))
}

func Fatal(msg string, s string) {
	lg.Fatal(log(msg, s))
}

// Funções para Logs Formatados

func Infof(messageID, template string, args ...interface{}) {
	Info(messageID, fmt.Sprintf(template, args...))
}

func Errorf(messageID, template string, args ...interface{}) {
	Error(messageID, fmt.Sprintf(template, args...))
}

func Debugf(messageID, template string, args ...interface{}) {
	Debug(messageID, fmt.Sprintf(template, args...))
}

// Funções para logs http

func Req(messageID, name, method, url string, timeout time.Duration, headers []byte, body []byte) {
	Infof(messageID, "%s | httpRequestConfig | TIMEOUT[%v] | METHOD[%s] | URL[%s]", name, timeout, method, url)
	Debugf(messageID, "%s | requestHeader | %s", name, Clean(headers))

	if len(body) == 0 {
		Debugf(messageID, "%s | requestPayload | %s", name, "empty")
	} else {
		Debugf(messageID, "%s | requestPayload | %s", name, Clean(body))
	}

	Infof(messageID, "%s | requestHeadersSize %dB | requestBodySize %dB | requestTotalSize %dB", name, len(headers), len(body), len(headers)+len(body))
	Infof(messageID, "%s | request starting", name)
}

func Resp(messageID, name string, headers []byte, body []byte, code int) {
	Debugf(messageID, "%s | responseHeader | %s", name, Clean(headers))

	if len(body) == 0 {
		Infof(messageID, "%s | responsePayload | %s", name, "empty")
	} else {
		Infof(messageID, "%s | responsePayload | %s", name, Clean(body))
	}

	Infof(messageID, "%s | responseContentLength | %dB", name, len(body))
	Infof(messageID, "%s | responseHttpStatus | %d", name, code)
}

// Utilitários

func Clean(payload []byte) string {
	p := bytes.ReplaceAll(payload, []byte("\t"), []byte(""))
	p = bytes.ReplaceAll(payload, []byte("\n"), []byte(""))
	return string(p)
}

func JsonToString(obj interface{}) string {
	s, err := json.Marshal(obj)
	if err != nil {
		return "{}"
	}
	return string(Clean(s))
}

func Elapsed(messageId, name string, init time.Time) {
	end := time.Now()
	duration := end.Sub(init)
	durationMs := int64(duration / time.Millisecond)
	durationString := strconv.FormatInt(durationMs, 10)
	Infof(messageId, "%s | totalTime | %s ms", name, durationString)
}

func timeNow() string {
	t := time.Now()
	timeNow := t.Format("20060102150405")
	return timeNow
}

func GenMessageid() (string, error) {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return newUuid.String(), nil
}
