package log

import "fmt"

type LogProcess struct {
	Rp chan string
	Pw chan string
	ReaderObj Reader
	WriterOjb Writer
}

type Reader interface {
	Read(rp chan string)
}

type Writer interface {
	Write(pw chan string)
}

type ReadFromFile struct {

}

type WriteToFile struct {

}

func (r *ReadFromFile) Read(rp chan string){
	rp <- "read log "
}

func (w *WriteToFile) Write(pw chan string){
	fmt.Println(<-pw + " write log")
}

func (l *LogProcess) ProcessLog(){
	str := <-l.Rp
	str += " process log"
	l.Pw <- str
}
