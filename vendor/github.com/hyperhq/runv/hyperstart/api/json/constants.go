package json

// when APIVERSION < 1000000, the version MUST be exactly matched on both sides
const VERSION = 4243

const (
	INIT_VERSION = iota // 0
	INIT_STARTPOD
	INIT_GETPOD_DEPRECATED
	INIT_STOPPOD_DEPRECATED
	INIT_DESTROYPOD
	INIT_RESTARTCONTAINER_DEPRECATED // 5
	INIT_EXECCMD
	INIT_FINISHCMD_DEPRECATED
	INIT_READY
	INIT_ACK
	INIT_ERROR // 10
	INIT_WINSIZE
	INIT_PING
	INIT_FINISHPOD_DEPRECATED
	INIT_NEXT
	INIT_WRITEFILE // 15
	INIT_READFILE
	INIT_NEWCONTAINER
	INIT_KILLCONTAINER
	INIT_ONLINECPUMEM
	INIT_SETUPINTERFACE // 20
	INIT_SETUPROUTE
	INIT_REMOVECONTAINER
	INIT_PROCESSASYNCEVENT
	INIT_SIGNALPROCESS
)

// "hyperstart" is the special container ID for adding processes.
const HYPERSTART_EXEC_CONTAINER = "hyperstart"

const HYPER_VSOCK_CTL_PORT = 2718
const HYPER_VSOCK_MSG_PORT = 2719
