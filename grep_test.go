package main

import "testing"

func TestGrep(t *testing.T) {
   got := Grep("pranav", "some_text.txt")
   want := []string{"pranavwadekar", "Ppranavisokokatwork"}
   if got != want {
       t.Errorf("method produced wrong result. expected: %d, got: %d", want, got)
  }
}