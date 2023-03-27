// cls - custom log service

package cls

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	loggerMode     int = 1
	logFile        *os.File
	logFileErr     error
	logStoragePath string

	onceInitLogFile sync.Once
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func SetLogStoragePath(p *string) {
	logStoragePath = *p
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Logs are stored with group by date
// Open (if exist) / create (if not exist) log file
func getLogFile() {

	//	delete old log file

	//	create log filename
	filename := fmt.Sprintf("%s/%s.log", logStoragePath, time.Now().Format("20060102"))

	//	if file not exist, create it, else just open it
	logFile, logFileErr = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

}

func saveLog(msg string) {
	if logFileErr != nil {
		return
	}

	logFile.Write([]byte(msg))
}

func processLog(msg string) {

	switch loggerMode {
	case 1:
		{
			saveLog(msg)
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Show / store program log
func Fail(msg string) {

	onceInitLogFile.Do(
		getLogFile,
	)

	processLog(
		fmt.Sprintf(
			"FAIL -- %s -- %s \n",
			time.Now().Format("2006-01-02 15:04:05"),
			msg,
		),
	)

}

// ------------------------------------------------------------------------ //

// Show / store
// program 'Info' log
func Info(msg string) {

	onceInitLogFile.Do(
		getLogFile,
	)

	processLog(
		fmt.Sprintf(
			"FAIL -- %s -- %s \n",
			time.Now().Format("2006-01-02 15:04:05"),
			msg,
		),
	)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
