package testutil

import (
	aliyun "github.com/AliyunContainerService/terway/pkg/aliyun/client"
	register "github.com/AliyunContainerService/terway/pkg/controller"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func NewManager(rest *rest.Config, openAPI aliyun.OpenAPI, directClient client.Client) (manager.Manager, *register.ControllerCtx) {
	_ = "STUB: not implemented"
	return *new(manager.Manager), nil
}
