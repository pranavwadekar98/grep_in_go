package main

import (
    "fmt"
    "os"
    "log"
    "sync"
    "bufio"
    "strings"
    "path/filepath"
)

func Grep(searched_string string, filename string, c chan []string){
    var output_lines [] string
    var file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755) //create if not exists
    if err != nil { //if any error (permission, user)
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file) //scanning the contents
    for scanner.Scan() { //reading file line by line
        if strings.Contains(scanner.Text(), searched_string) {
            output_lines = append(output_lines, scanner.Text()) //storing if pattern matches.
        }
    }
    c <- output_lines
//     fmt.Println(output_lines, <-c)
}

func getFilesFromAllSubDirectories(path string)([]string, error){
    var files []string
    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if filepath.Ext(path) != ".txt" {
            return nil
        }
        files = append(files, path)
        return nil
        }) //getting all the .text files from all the sub folders
    return files, err
}

func findOptimizedMatches(files []string, string_to_be_searched string) ([]string){
    var matches []string
    var wg sync.WaitGroup
    c := make(chan []string)

    for _, file := range files {
        wg.Add(1) //adding count of goroutines for waitgroup
        go Grep(string_to_be_searched, file, c)
    }
    go func() {
		for i := range c {
			matches = append(matches, i...)
			wg.Done() // decrement the counter as
		}
	}()
	wg.Wait() //wait till counter goes to 0
	return matches
}

func main(){
    var matches []string

    root := "/home/pranavwadekar/jana/grep_in_go/"
    string_to_be_searched := "pranav"

    _, err1 := os.Stat(root)

    if os.IsNotExist(err1) {
        log.Fatal("File/Directory does not exist.")
    }
    files, err := getFilesFromAllSubDirectories(root)

    if err != nil {
        fmt.Println(err)
    }
    matches = findOptimizedMatches(files, string_to_be_searched)
	fmt.Println(matches)

}