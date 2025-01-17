package main

import (
	"fmt"
	"os"
)

func main() {
	// export KUBERNETES_SERVICE_HOST=kubernetes.default.svc
	// export KUBERNETES_SERVICE_PORT=443
	fmt.Println(os.Getenv("KUBERNETES_SERVICE_PORT"))
}
