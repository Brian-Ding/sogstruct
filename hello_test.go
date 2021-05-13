package gostruct

import "testing"

func TestHello(t *testing.T){
	want := "Hello: Brian"
	if got := Hello("Brian"); got != want{
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}