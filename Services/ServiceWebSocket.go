package Services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Ping(c *gin.Context)  {
	ws, err :=upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(message))
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)


		if err != nil {
			break
		}
	}
}