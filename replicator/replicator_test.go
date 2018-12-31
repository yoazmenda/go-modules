package replicator

import "testing"

func TestHelloWorld(t *testing.T) {
	if replicate() != "replicate!!" {
		t.Fatal("Test fail")
	}
}
