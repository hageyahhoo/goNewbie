package ut

import (
	"testing"
)

func TestConfig(t *testing.T) {
	result := Config()
	if result != "kubeconfig" {
		t.Errorf("Failure: kubeconfig is not %s", result)
	}
	t.Log("result", result)
}
