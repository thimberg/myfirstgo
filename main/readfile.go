package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "io"
    "os"
    "strings"
)

func readAll (fileName string) {
    b, err := ioutil.ReadFile(fileName) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
}

func readLine (fileName string) []string {
    var rtnArr []string

    fp, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    reader := bufio.NewReaderSize(fp, 4096)

    for {
        line, _, err := reader.ReadLine()
//        fmt.Println(string(line))
        rtnArr = append(rtnArr, string(line))
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
    }

    return rtnArr
}

type TargetCSV struct {
	url string
	urlSelector string
	titleSelector string
	contentSelector string
	encoding string
}

func readCsv (fileName string) []TargetCSV {
	var rtnArr []TargetCSV
	
	for _, value := range readLine(fileName) {
		
		items := strings.Split(value, ",")
		if len(items) == 5 {
			url := strings.Trim(strings.Trim(items[0], ""), "\"")
			urlSelector := strings.Trim(strings.Trim(items[1], ""), "\"")
			titleSelector := strings.Trim(strings.Trim(items[2], ""), "\"")
			contentSelector := strings.Trim(strings.Trim(items[3], ""), "\"")
			encoding := strings.Trim(strings.Trim(items[4], ""), "\"")
			rtnArr = append(rtnArr, TargetCSV {url, urlSelector, titleSelector, contentSelector, encoding} )
		}
	}
	
	return rtnArr
}

///func main() {
///
///    readAll("input.txt") 
///    for _, value := range readLine("input.txt") {
///        fmt.Println(value)
///    }
///}
