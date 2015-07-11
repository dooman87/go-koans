package go_koans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
    "time"
    "path/filepath"
    "log"
)

const (
	__string__       string  = "impossibly lame value"
	__int__          int     = -1
	__positive_int__ int     = 42
	__byte__         byte    = 255
	__bool__         bool    = false // ugh
	__boolean__      bool    = true  // oh well
	__float32__      float32 = -1.0
	__delete_me__    bool    = false
)

var __runner__ runner = nil
var failed bool = false
var lastModified time.Time = time.Now()
type koan func()

func TestKoans(t *testing.T) {
    log := log.New(os.Stdout, "logger:", log.Lshortfile)

    //currDirName := filepath.Dir(os.Args[0])
    currDirName := "."
    currDir, err := os.Open(currDirName)
    if err != nil {
        log.Fatal(err)
    }
    files,err := currDir.Readdir(0)
    if err != nil {
        log.Fatal(err)
    }

    var absFilesForWatch []string
    for _,v := range files {
        if strings.HasSuffix(v.Name(), ".go") {
            absFile,_ := filepath.Abs(v.Name())
            absFilesForWatch = append(absFilesForWatch, absFile)
            fmt.Printf("FILE [%s] last modified [%s]\n", absFile, v.ModTime().String())
        }
    }

    waitUntilModified := func() {
        var changed string
        for {
            //files,_ := currDir.Readdir(0)
            for _,v := range absFilesForWatch {
                //absFile,_ := filepath.Abs(v.Name())
                stat,_ := os.Stat(v)
                if stat.ModTime().Unix() > lastModified.Unix() {
                    lastModified = stat.ModTime()
                    changed = v
                }
            }
            if len(changed) > 0 {
                fmt.Printf("Changed %s\n", changed)
                return
            }
            time.Sleep(1 * time.Second)
            fmt.Printf("Waiting\n")
        }
    }

    waitUntilModified()

    tests := []koan{aboutBasics, aboutStrings}
    for i,v := range tests {
        //fmt.Printf("Run %s\n", v.name)
        if i == 0 {
            failed = false
        }
        v()
    }

    //for {
	//    aboutBasics()
	//    aboutStrings()
	//    aboutArrays()
	//    aboutSlices()
	//    aboutTypes()
	//    aboutControlFlow()
    //    aboutEnumeration()
	//    aboutAnonymousFunctions()
	//    aboutVariadicFunctions()
	//    aboutFiles()
	//    aboutInterfaces()
	//    aboutCommonInterfaces()
	//    aboutMaps()
	//    aboutPointers()
	//    aboutStructs()
	//    aboutAllocation()
	//    aboutChannels()
	//    aboutConcurrency()
	//    aboutPanics()
         //time.Sleep(1 * time.Second)
    //}

    if !failed {
	    fmt.Printf("\n%c[32;1mYou won life. Good job.%c[0m\n\n", 27, 27)
    }
    os.Exit(0)
}

func assert(o bool) {
	if !failed && !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
        failed = true
		//os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
