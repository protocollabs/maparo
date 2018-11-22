package server

import "fmt"
import "os"
import "github.com/monfron/mapago/controlPlane/ctrl/serverProtos"
import "github.com/monfron/mapago/controlPlane/ctrl/shared"

var CTRL_PORT = 64321
var DEF_BUFFER_SIZE = 8096 * 8
var ID string
var ARCH string
var OS string
var MODULES string

func RunServer(laddr string, port int, callSize int) {
	var repDataObj *shared.DataObj

	// construct apriori
	ID = shared.ConstructId()
	OS = shared.DetectOs()
	ARCH = shared.DetectArch()
	MODULES = shared.ConvMapToStr(supportedModules())

	ch := make(chan shared.ChResult)

	/* Disabled during udp mc dev*/


		tcpObj := serverProtos.NewTcpObj("TcpConn1", laddr, port, callSize)
		tcpObj.Start(ch)


	/*
		udpObj := serverProtos.NewUdpObj("UdpConn1", laddr, port, callSize)
		udpObj.Start(ch)
	*/

	/*
	udpMcObj := serverProtos.NewUdpMcObj("UdpMcConn1", laddr, port, callSize)
	udpMcObj.Start(ch)
	*/

	for {
		request := <-ch
		reqDataObj := shared.ConvJsonToDataStruct(request.Json)

		switch reqDataObj.Type {
		case shared.INFO_REQUEST:
			repDataObj = constructInfoReply(reqDataObj)
		case shared.INFO_REPLY:
			fmt.Println("I received an INFO_REPLY, I am ignoring this")
			continue
		case shared.MEASUREMENT_START_REQUESTS:
			fmt.Println("Construct MEASUREMENT_START_REP")

		case shared.MEASUREMENT_STOP_REQUEST:
			fmt.Println("Construct MEASUREMENT_STOP_REP")

		case shared.MEASUREMENT_INFO_REQUEST:
			fmt.Println("Construct MEASUREMENT_INFO_REP")

		case shared.TIME_DIFF_REQUEST:
			fmt.Println("Construct TIME_DIFF_REP")

		case shared.WARNING_ERR_MSG:
			fmt.Println("WARNING_ERR_MSG")

		default:
			fmt.Printf("Unknown type")
			os.Exit(1)
		}

		json := shared.ConvDataStructToJson(repDataObj)
		request.ConnObj.WriteAnswer(json)
		request.ConnObj.CloseConn()
		break
	}
}

func constructInfoReply(reqDataObj *shared.DataObj) *shared.DataObj {
	fmt.Println("\nConstructing INFO_REP")

	// construct ID
	repDataObj := new(shared.DataObj)
	repDataObj.Type = shared.INFO_REPLY
	repDataObj.Id = ID
	repDataObj.Seq = "0"
	repDataObj.Seq_rp = reqDataObj.Seq
	// maparo STD: timestamp replied untouched by server
	reqDataObj.Ts = reqDataObj.Ts
	repDataObj.Modules = MODULES
	repDataObj.Arch = ARCH
	repDataObj.Os = OS
	repDataObj.Info = "fancyInfo"
	return repDataObj
}

func supportedModules() map[string]string {
	fmt.Println("\nConstructing supported modules")
	supportedMods := make(map[string]string)
	supportedMods["udp-goodput"] = "no-support"
	supportedMods["tcp-goodput"] = "no-support"
	return supportedMods
}
