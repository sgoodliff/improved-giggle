package hello

import "testing"
// check that hello() returns expected
func TestHello(t *testing.T){
  expected := "Hello Go!"
  actual := hello()
  if actual != expected {
    t.Error("Test failed")
  }
}
