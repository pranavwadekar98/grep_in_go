package sample_subroutine_main

import "fmt"
import (
    "os"
    "bufio"
    "strings"
)

func matchString(pattern string, arr []string, c chan []string){
    var output_lines [] string
    for _, v := range arr {
		if strings.Contains(v, pattern) {
            output_lines = append(output_lines, v)
        }
	}
	c <- output_lines
}

func Grep(searched_string string, filename string) ([]string){
    var file_out [] string
    var output_lines [] string
    var file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        panic(err)
        return output_lines
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
            file_out = append(file_out, scanner.Text())
        }
    c := make(chan []string)
    go matchString(searched_string, file_out[:len(file_out)/2], c)
	go matchString(searched_string, file_out[len(file_out)/2:], c)
    output_lines = append(<-c, <-c...)
    return output_lines
}

func main(){
    output := Grep("pranav", "some_text.txt")
    fmt.Println(output)
}