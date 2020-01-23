package main

import "testing"

func TestSampleHello(t *testing.T) {

	spec := "Hello Again!"
	if got := sampleHello(); got != spec {
		t.Errorf("sampleHello() = %q, want %q", got, spec)
	}

}
