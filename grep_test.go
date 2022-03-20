package main

import (
        "fmt"
        "testing"
        "reflect"
        )

func TestGrep(t *testing.T) {
   got := findOptimizedMatches([]string{"/home/pranavwadekar/jana/grep_in_go/some_text.txt"}, "pranav")
   want := []string{"pranavwadekar", "Ppranavisotwork", "ppranvisgoodatchess randompranav ispranavispranav"}
   if reflect.DeepEqual(got, want) {
       fmt.Println(want, got)
  } else {
       fmt.Printf("Test case failed\n want: %#v\n got:%#v\n", want, got)
       t.Errorf("Test Case Failed")
  }
}

func TestGetAllFiles(t *testing.T) {
   got, err := getFilesFromAllSubDirectories("/home/pranavwadekar/jana/grep_in_go/")
   want := []string{"/home/pranavwadekar/jana/grep_in_go/sample_files/sample.txt",
                    "/home/pranavwadekar/jana/grep_in_go/sample_files/sample1.txt",
                    "/home/pranavwadekar/jana/grep_in_go/some_text.txt"}
   if err != nil{
        panic(err)
   } else if reflect.DeepEqual(got, want) {
       fmt.Println(want, got)
  } else {
       fmt.Printf("Test case failed\n want: %#v\n got:%#v\n", want, got)
       t.Errorf("Test Case Failed")
  }
}

