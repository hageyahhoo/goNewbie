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

func TestServer(t *testing.T) {
	for key, expected := range map[string]string{
		"kube-apiserver":          "kube-apiserver",
		"kube-controller-manager": "kube-controller-manager",
		"xxx":                     "kube-scheduler",
	} {
		if actual := Server(key); expected != actual {
			t.Errorf("Failure: expected %q but returned %q", expected, actual)
		}
	}
}
