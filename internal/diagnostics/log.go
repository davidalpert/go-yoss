package diagnostics

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/go-yoss/internal/diagnostics/plaintext"
	"github.com/davidalpert/go-yoss/internal/env"
	"github.com/davidalpert/go-yoss/internal/version"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var Log *log.Entry

const (
	ENVKEY_LOG_LEVEL  = "yoss_LOG_LEVEL"
	ENVKEY_LOG_FORMAT = "yoss_LOG_FORMAT"
	ENVKEY_LOG_FILE   = "yoss_LOG_FILE"
)

func init() {
	Log = log.WithFields(log.Fields{
		"app":         version.Detail.AppName,
		"app_version": version.Detail.Version,
	})

}

func ConfigureLogger(streams printers.IOStreams) (cleanupFn func()) {
	// default cleanup: nothing to do
	cleanupFn = func() {}

	// configure logging
	logLevel := env.GetValueOrDefaultLogLevel(ENVKEY_LOG_LEVEL, log.FatalLevel)
	log.SetLevel(logLevel)
	// log sink
	var sink io.Writer
	var logFile = env.GetValueOrDefaultString(ENVKEY_LOG_FILE, "")
	var logDestination = "stdout"
	if logFile == "" {
		sink = streams.Out
	} else {
		fullPath, err := filepath.Abs(logFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		logFile, err := os.Create(fullPath)
		if err != nil {
			log.Fatal(err.Error())
		}
		cleanupFn = func() { logFile.Close() }
		logDestination = fullPath
		sink = logFile
	}

	if strings.EqualFold(env.GetValueOrDefaultString(ENVKEY_LOG_FORMAT, "text"), "json") {
		log.SetHandler(json.New(sink))
	} else {
		log.SetHandler(plaintext.New(sink))
	}

	Log.WithField("destination", logDestination).Debug("logging initialized")

	return
}
