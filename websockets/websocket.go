package websockets

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshishKothariii/gotalkapi/services"
	"github.com/AshishKothariii/gotalkapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Replace this with your logic to allow specific origins
	
		// For example, allow requests from any origin for testing
		return true
	},
}

type Client struct {
	username string
	userId primitive.ObjectID
	conn *websocket.Conn
	send chan []byte
}
type WebSocketController struct {
        service services.UserService
		hub *Hub
}

// NewUserController creates a new UserController
func NewWebSocketController(s services.UserService,h *Hub) *WebSocketController {
        return &WebSocketController{service: s,hub: h}

}

func (c *Client) write() {
	defer c.conn.Close()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case msg := <-h.broadcast:
			for client := range h.clients {
				client.send <- msg
			}
		}
	}
}

func  (wc *WebSocketController)HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	token,err :=c.Cookie("token")
     if err!=nil{
		log.Print(err)
	 }
	 ans,err:=utils.ParseJWT(token)
	 if err!=nil{
		fmt.Print(err)
	 }
	
	userid,err :=primitive.ObjectIDFromHex(ans)
	 if err!=nil{
		fmt.Print(err)
	 }
	user,err:= wc.service.GetUserByID(c,userid)
	  if err!=nil{
		fmt.Print(err)
	 }
	client := &Client{conn: conn, send: make(chan []byte)}
	client.username=user.Email
	client.userId=user.ID
	wc.hub.register <- client
    
	go client.write()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			wc.hub.unregister <- client
			break
		}
		wc.hub.broadcast <- msg
	}
}
