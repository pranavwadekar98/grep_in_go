package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strings"
    "path/filepath"
)

func Grep(searched_string string, filename string, c chan []string){
    var output_lines [] string
    var file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), searched_string) {
            output_lines = append(output_lines, scanner.Text())
        }
    }
    c <- output_lines
//     fmt.Println(output_lines, <-c)
}

func main(){
    var files []string
    var matches []string

    c := make(chan []string, 3)

    root := "/home/pranavwadekar/jana/grep_in_go/"
    string_to_be_searched := "pranav"
    _, err1 := os.Stat(root)

    if os.IsNotExist(err1) {
        log.Fatal("File does not exist.")
    }
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if filepath.Ext(path) != ".txt" {
            return nil
        }
        files = append(files, path)
        return nil
        })
    if err != nil {
        fmt.Println(err)
    }
    for _, file := range files {
        go Grep(string_to_be_searched, file, c)
    }
    for i := range c {
		matches = append(matches, i...)
		fmt.Println(i)
	}
}