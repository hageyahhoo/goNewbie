package ut

import (
	"testing"
)

func TestConfig(t *testing.T) {
	result := Config()
	if result != "kubeconfig" {
		t.Errorf("Failure: kubeconfig is not %q", result)
	}
	t.Log("result", result)
}
