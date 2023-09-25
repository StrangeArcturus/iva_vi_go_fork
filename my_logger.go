package ivavigofork


import (
	"os"
	"fmt"
)


type myLogger struct {
	LogFile string
}


func (self myLogger) out(text string, lvl string) string {
	result := fmt.Sprintf("[%s]", lvl) + text
	if len(self.LogFile) > 0 {
		f, err := os.OpenFile(self.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err := f.WriteString(text); err != nil {
			panic(err)
		}
	}
	return result
}


func (self myLogger) Warn (text string) string {
	return self.out(text, "WARNING")
}


func (self myLogger) Error (text string) string {
	return self.out(text, "ERROR")
}


func (self myLogger) Log (text string) string {
	return self.out(text, "LOG")
}


var MyLogger = myLogger{Config.LogFile}
