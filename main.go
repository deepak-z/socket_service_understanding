package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Home Page");
}

func reader(conn *websocket.Conn){
	for {
		//2. here that message gets read
		messageType,p,err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return 
		}

		//3. Printed on our console
		log.Println(string(p))

		//5. Sent back to the client
		if err := conn.WriteMessage(messageType,p); err != nil{
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter,r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool {return true};
	ws,err := upgrader.Upgrade(w,r,nil)
	if(err != nil){
		log.Println("err")
	}
	log.Println("Client Successfully Connected....")
	reader(ws)
}
func setupRoutes(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/ws",wsEndpoint)
}
func main(){
	fmt.Println("Go WebSockets ")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":7080",nil))
}