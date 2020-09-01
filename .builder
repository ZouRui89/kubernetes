build: true
go_version: 1.14.7
build_cmd: export SYM_K8S_VERSION=${tag} && make clean && KUBE_BUILD_PLATFORMS=linux/amd64 make all GOFLAGS=-v GOGCFLAGS="-N -l"
output_dir: _output/bin
upload_bin: kube-apiserver kube-controller-manager kube-scheduler kubectl kubelet kubeadm