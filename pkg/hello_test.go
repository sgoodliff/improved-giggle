package hello

import "testing"

func TestHello(t *testing.T){
  expected := "Hello Go!"
  actual := hello("Perl")
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}
