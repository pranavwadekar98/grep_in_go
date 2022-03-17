package naive_main

import (
        "fmt"
        "testing"
        "reflect"
        )

func TestGrep(t *testing.T) {
   got := Grep("pranav", "some_text.txt")
   want := []string{"pranavwadekar", "Ppranavisokokatwork"}
   if reflect.DeepEqual(got, want) {
       fmt.Println(want, got)
  }
}