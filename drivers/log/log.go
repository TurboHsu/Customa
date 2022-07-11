package log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	c "github.com/TurboHsu/Customa/drivers/config"
	"github.com/TurboHsu/Customa/structs"
)

//Logging function, exporting to console and log file.
//Need to divide by levels, error, info, and others.

func Error(err string) {
	errMsg := fmt.Sprintf("%s [E] %s", time.Now().Format("2006-01-02 15:04:05"), err)
	if c.Config.Log.ConsoleOutput {
		fmt.Println(errMsg)
	}
	if c.Config.Log.WriteError {
		e := writeLog(errMsg)
		if e.IsErr {
			fmt.Printf("! [E] Cannot write log caz :%s", e.Msg)
		}
	}
}

func Info(info string) {
	infoMsg := fmt.Sprintf("%s [I] %s", time.Now().Format("2006-01-02 15:04:05"), info)
	if c.Config.Log.ConsoleOutput {
		fmt.Println(infoMsg)
	}
	if c.Config.Log.WriteInfo {
		e := writeLog(infoMsg)
		if e.IsErr {
			fmt.Printf("! [E] Cannot write log caz : %s", e.Msg)
		}
	}
}

//Writes the log to file
func writeLog(log string) structs.ErrMsg {
	file, err := os.OpenFile(c.Config.Log.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	var error structs.ErrMsg = structs.ErrMsg{IsErr: false, Msg: ""}
	if err != nil {
		error.IsErr = true
		error.Msg = err.Error()
		return error
	}
	defer file.Close()

	//Define a buffer
	w := bufio.NewWriter(file)
	w.WriteString(fmt.Sprintf("%s\n", log))
	//Write
	w.Flush()
	return error
}
