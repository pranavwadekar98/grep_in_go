package naive_main

import "fmt"
import (
    "os"
    "bufio"
    "strings"
)

func Grep(searched_string string, filename string) ([]string){
    var output_lines [] string
    var file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        panic(err)
        return output_lines
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), searched_string) {
            output_lines = append(output_lines, scanner.Text())
        }
    }
    return output_lines
}

func main(){
    output := Grep("pranav", "some_text.txt")
    fmt.Println(output)
}