package ut

func Config() string {
	return "kubeconfig"
}

func Server(key string) string {
	if key == "kube-apiserver" {
		return "kube-apiserver"
	}
	if key == "kube-controller-manager" {
		return "kube-controller-manager"
	}
	return "kube-scheduler"
}
