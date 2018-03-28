package pulser

import (
	"fmt"
	"log"
	"net"
	"pulse/pulser/mtproto"

	"github.com/davecgh/go-spew/spew"
)

func HandleConnection(conn net.Conn) error {
	fmt.Println("Handling new connection...")

	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	conCacheData := &mtproto.CacheData{}

	for {
		data, err := mtproto.ReadData(conn, conCacheData)

		if err != nil {
			panic(err)
		}

		switch data.(type) {
		case mtproto.TL_req_pq:
			go handlerReqPQ(data, conn, conCacheData)
		case mtproto.TL_req_DH_params:
			go handlerReqDHParams(data, conn, conCacheData)
		case mtproto.TL_set_client_DH_params:
			go handlerSetClientDHParams(data, conn, conCacheData)
		case mtproto.TL_invokeWithLayer:
			go handlerinvokeWithLayer(data, conn, conCacheData)
		case mtproto.TL_msgs_ack:
			go handlerMsgsAck(data, conn, conCacheData)

		default:
			spew.Dump(data)
			log.Println("handler not found")
		}

	}
}
