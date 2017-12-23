package sys_rpc

import (
	"github.com/jmesyan/xingo/cluster"

	"github.com/jmesyan/xingo/clusterserver"
	"github.com/jmesyan/xingo/logger"
	"github.com/jmesyan/xingo/utils"
)

type RootRpc struct {
}

/*
子节点连上来的通知
*/
func (this *RootRpc) TakeProxy(request *cluster.RpcRequest) {
	name := request.Rpcdata.Args[0].(string)
	logger.Info("child node " + name + " connected to " + utils.GlobalObject.Name)
	//加到childs并且绑定链接connetion对象
	clusterserver.GlobalClusterServer.AddChild(name, request.Fconn)
}
