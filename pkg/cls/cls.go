// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package cls

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type Log struct {
	Err error
	Msg string
	Add map[string]string
}

// ------------------------------------------------------------------------ //

func (log Log) formatAdd() string {
	if b, bErr := json.Marshal(log.Add); bErr == nil {
		return string(b)
	} else {
		panic(bErr)
	}
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func Info(log Log) {
	fmt.Printf(
		"Lvl:INFO;:;Msg:%s;:;Add:%s;:;Ts:%s;:; \n",
		remExtraSpace(log.Msg),
		log.formatAdd(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
}

// ------------------------------------------------------------------------ //

func Warn(log Log) {
	fmt.Printf(
		"Lvl:WARN;:;Msg:%s;:;Err:%s;:;Add:%s;:;Ts:%s;:; \n",
		remExtraSpace(log.Msg),
		log.Err.Error(),
		log.formatAdd(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
}

// ------------------------------------------------------------------------ //

func Fail(log Log) {
	fmt.Printf(
		"Lvl:FAIL;:;Msg:%s;:;Err:%s;:;Add:%s;:;Ts:%s;:; \n",
		remExtraSpace(log.Msg),
		log.Err.Error(),
		log.formatAdd(),
		time.Now().Format("2006-01-02 15:04:05"),
	)

	os.Exit(1)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func remExtraSpace(m string) string {
	return strings.Join(strings.Fields(m), " ")
}
