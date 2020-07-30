package websocket

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
	"kGoChat/datamodel"
)

var ws *websocket.Server

var websocketControllers = map[string]*WebsocketController{}

func ConfigureMVC(m *mvc.Application) {
	ws = websocket.New(websocket.Config{})
	// /websocket/iris-ws.js
	m.Router.Any("/iris-ws.js", websocket.ClientHandler())

	// This will bind the result of ws.Upgrade which is a websocket.Connection
	// to the controller(s) served by the `m.Handle`.
	m.Register(ws.Upgrade)
	m.Handle(new(WebsocketController))
}

type WebsocketController struct {
	// Note that you could use an anonymous field as well, it doesn't matter, binder will find it.
	//
	// This is the current websocket connection, each client has its own instance of the *WebsocketController.
	Conn websocket.Connection

	user datamodel.UserInfo
}

func (c *WebsocketController) onLeave(roomName string) {
	// This will call the "visit" event on all clients, except the current one,
	// (it can't because it's left but for any case use this type of design)
	delete(websocketControllers, c.Conn.ID())
}

func (c *WebsocketController) onJoin() {
	// This will call the "visit" event on all clients, including the current
	// with the 'newCount' variable.
	//_ = c.Conn.To(websocket.All).Emit("Join", c.Conn.ID())
}

func (c *WebsocketController) Get( /* websocket.Connection could be lived here as well, it doesn't matter */ ) {
	websocketControllers[c.Conn.ID()] = c

	c.user.UUID = c.Conn.ID()
	c.user.WebSocketID = c.Conn.ID()

	c.Conn.OnLeave(c.onLeave)
	c.Conn.On("Chat", c.onChat)
	c.Conn.On("User", c.onUser)
	c.Conn.On("WebRTC", c.onWebRTC)

	c.onJoin()
	// call it after all event callbacks registration.
	c.Conn.Wait()
}
