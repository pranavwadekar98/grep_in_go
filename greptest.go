package Testing

import "testing"
import "./main"

func TestGrep(t *testing.T) {
  if got, want := Grep("pranav", 2), []string{"pranavwadekar", "Ppranavisokokatwork"}; got != want {
    t.Errorf("method produced wrong result. expected: %d, got: %d", want, got)
  }
}