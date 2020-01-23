package sms

import "testing"

func TestSmSender(t *testing.T) {

	expected := "Sending Message SMS"
	if got := SendSMS(); expected != got {
		t.Errorf("sampleHello() = %q, want %q", got, expected)
	}

}
