package git

import "fmt"

type Message struct {
	Name string
	Message string
}

func (gw *gitworker) message( message string) {
	gw.messages <- Message{gw.gr.Name, message}
}

func (gw *gitworker) messagef(message string, params ...interface{}){
	gw.messages <- Message{gw.gr.Name, fmt.Sprintf(message, params...)}
}