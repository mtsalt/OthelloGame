package main

import (
	"fmt"
	"log"
	"net/http"

	"BoardGame/othello"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://news.mynavi.jp/techplus/article/gogogo-17/

func main() {

	o := othello.NewMatch()
	o.UpdateBoard(1, 2)

	// fmt.Println("Start server")
	// var server = InitRouter()
	// err := server.Run("0.0.0.0:8089")
	// if err != nil {
	// 	log.Fatal("Server error")
	// }
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("static", "static")

	// Get request list
	router.GET("/", indexPage)
	router.GET("/ws", WsEndpoint)

	return router
}

// Get request func
func indexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WebSocketsからの返却用データの構造体
type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

func WsEndpoint(ctx *gin.Context) {
	// HTTPサーバーコネクションをWebSocketsプロトコルにアップグレード
	ws, err := upgradeConnection.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		log.Println("#####################", string(message))
		if err != nil {
			fmt.Println(err)
			break
		}

		if string(message) == "ping" {
			message = []byte("pong")
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	// log.Println("OK Client Connecting")

	// var response WsJsonResponse
	// response.Message = `Connected server with websocket`

	// err = ws.WriteJSON(response)
	// if err != nil {
	// 	log.Println(err)
	// }
}
