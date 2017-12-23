package xingo

import (
	"fmt"
	"github.com/jmesyan/xingo/cluster"
	"github.com/jmesyan/xingo/clusterserver"
	_ "github.com/jmesyan/xingo/fnet"
	"github.com/jmesyan/xingo/fserver"
	"github.com/jmesyan/xingo/iface"
	"github.com/jmesyan/xingo/logger"
	"github.com/jmesyan/xingo/sys_rpc"
	"github.com/jmesyan/xingo/telnetcmd"
	_ "github.com/jmesyan/xingo/timer"
	"github.com/jmesyan/xingo/utils"
)

func NewXingoTcpServer() iface.Iserver {
	//do something
	//debugport 是否开放
	if utils.GlobalObject.DebugPort > 0 {
		if utils.GlobalObject.Host != "" {
			fserver.NewTcpServer("telnet_server", "tcp4", utils.GlobalObject.Host,
				utils.GlobalObject.DebugPort, 100, cluster.NewTelnetProtocol()).Start()
		} else {
			fserver.NewTcpServer("telnet_server", "tcp4", "127.0.0.1",
				utils.GlobalObject.DebugPort, 100, cluster.NewTelnetProtocol()).Start()
		}
		logger.Debug(fmt.Sprintf("telnet tool start: %s:%d.", utils.GlobalObject.Host, utils.GlobalObject.DebugPort))

	}

	//add command
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
	}

	s := fserver.NewServer()
	return s
}

func NewXingoMater(cfg string) *clusterserver.Master {
	s := clusterserver.NewMaster(cfg)
	//add rpc
	s.AddRpcRouter(&sys_rpc.MasterRpc{})
	//add command
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewCloseServerCommand())
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewReloadCfgCommand())
	}
	return s
}

func NewXingoCluterServer(nodename, cfg string) *clusterserver.ClusterServer {
	s := clusterserver.NewClusterServer(nodename, cfg)
	//add rpc
	s.AddRpcRouter(&sys_rpc.ChildRpc{})
	s.AddRpcRouter(&sys_rpc.RootRpc{})
	//add cmd
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
	}
	return s
}
