package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connetction struct {
	weConn *websocket.Conn
	inChan chan[]byte
	outChan chan[]byte
	closeChan chan byte
	isClose bool
	mutex sync.Mutex
}

func InitConnection(weConn *websocket.Conn)(conn *Connetction, err error){
	conn = &Connetction{
		weConn:weConn,
		inChan:make(chan []byte, 1000),
		outChan:make(chan []byte, 1000),
		closeChan:make(chan byte),
	}

	go conn.readLoop()
	go conn.writeLoop()
	return
}


func (conn *Connetction)ReadMessage()(data []byte, err error){
	select {
		case data =<- conn.inChan:
		case <- conn.closeChan:
			err = errors.New("connection is closed")
	}
	return
}

func (conn *Connetction)WriteMessage(data []byte)(err error){
	select {
		case conn.outChan <- data:
		case <- conn.closeChan:
			err = errors.New("connection is closed")
	}
	conn.outChan <- data

	return
}

func (conn *Connetction)Close(){
	conn.weConn.Close()
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
	}
	conn.mutex.Unlock()
}

func (conn *Connetction)readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.weConn.ReadMessage(); err != nil {
			conn.Close()
		}
		select {
		case conn.inChan <- data:
		case  <- conn.closeChan:
			goto ERR
		}

	}
	ERR:
		conn.Close()
}

func (conn *Connetction)writeLoop(){
	var(
		data []byte
		err error
	)
	for{
		select {
		case data =<- conn.outChan:
		case <- conn.closeChan:
			goto ERR
		}
		if err = conn.weConn.WriteMessage(websocket.TextMessage,data);err != nil{
			conn.Close()
		}
	}

	ERR:
		conn.Close()
}