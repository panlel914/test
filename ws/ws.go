package ws

import (
	"fmt"
	"net/http"
	"test/impl"
	"time"
)
import "github.com/gorilla/websocket"

var(
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request){
	var(
		wsConn *websocket.Conn
		err error
		conn *impl.Connetction
		data []byte
	)

	if wsConn,err = upgrader.Upgrade(w,r, nil); err != nil{
		return
	}

	if conn,err = impl.InitConnection(wsConn); err != nil{
		goto ERR
	}

	go func(){
		for{
			if err = conn.WriteMessage([]byte("heartbeat"));err != nil{
				return
			}

			time.Sleep(1*time.Second)
		}
	}()

	for{
		if data,err = conn.ReadMessage();err != nil{
			goto ERR
		}

		if err = conn.WriteMessage(data);err != nil{
			goto ERR
		}
	}


	//for{
	//	if _,data,err = conn.ReadMessage();err != nil{
	//		goto ERR
	//	}
	//	fmt.Println(data)
	//	if err = conn.WriteMessage(websocket.TextMessage,data);err != nil{
	//		goto ERR
	//	}
	//}

	ERR:
		conn.Close()
		fmt.Println("error")
		//关闭连接

}

func StartWS(){
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}