package neonlog

import (
    "log"
    "os"
    "time"
)

//Logger - main struct for logging
type Logger struct{

    pathToLogs string
    filePrefix string
    warningLogger *log.Logger
    infoLogger    *log.Logger
    errorLogger   *log.Logger
    debugLogger   *log.Logger
    initialized bool
    lastDate string
    debug bool
    file *os.File

}

//Init - initialization for logger struct
func (l *Logger) Init(path string, prefix string, debug bool) error{
    
    var logFileName string
    curTime:=time.Now()
    logFileName = path+"/"+prefix+"-"+curTime.Format("2006-01-02")+".log"

    file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    l.initialized=true
    l.file=file
    l.pathToLogs=path
    l.filePrefix=prefix
    l.debug=debug

    l.infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
    l.warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
    l.errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    l.debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime)

    return nil
}

func (l *Logger) checkDate() error{

    curTime:=time.Now()
    curDate:=curTime.Format("2006-01-02")

    if curDate!=l.lastDate{
        l.lastDate=curDate
        l.file.Close()

        logFileName := l.pathToLogs+"/"+l.filePrefix+"-"+curDate+".log"

        file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
            return err
        }
        l.file=file

        

        l.infoLogger.SetOutput(file)
        l.infoLogger.SetPrefix("INFO: ")
        l.infoLogger.SetFlags(log.Ldate|log.Ltime)

        l.warningLogger.SetOutput(file)
        l.warningLogger.SetPrefix("WARNING: ")
        l.warningLogger.SetFlags(log.Ldate|log.Ltime)

        l.errorLogger.SetOutput(file)
        l.errorLogger.SetPrefix("ERROR: ")
        l.errorLogger.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

        l.debugLogger.SetOutput(file)
        l.debugLogger.SetPrefix("DEBUG: ")
        l.debugLogger.SetFlags(log.Ldate|log.Ltime)
    }

    return nil
}

//Info - log msg with Info level
func (l *Logger) Info(str string) {

    err:=l.checkDate()
    if err!=nil{
        //TODO: set error bit in reg
    }else if l.initialized{
        l.infoLogger.Println(str)
    }
}

//Debug - log msg with Debug level
func (l *Logger) Debug(str string) {

    err:=l.checkDate()
    if err!=nil{
        //TODO: set error bit in reg
    }else if l.initialized && l.debug{
        l.debugLogger.Println(str)
    }
}

//Warning - log msg with Warning level
func (l *Logger) Warning(str string) {

    err:=l.checkDate()
    if err!=nil{
        //TODO: set error bit in reg
    }else if l.initialized{
        l.warningLogger.Println(str)
    }
}

//Error - log msg with Error level
func (l *Logger) Error(str string) {

    err:=l.checkDate()
    if err!=nil{
        //TODO: set error bit in reg
    }else if l.initialized{
        l.errorLogger.Println(str)
    }
}