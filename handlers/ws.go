package handlers

import (
	"chassis/dep"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func ws(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {
	// upgrade http handler to websocket server
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}
	defer conn.Close()

	// start server
	for {
		// JSON MESSAGES
		var question string
		err := conn.ReadJSON(&question)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}

		err = conn.WriteJSON("answer")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}

		// BYTE MESSAGES
		bytes, err := os.ReadFile("./speech.mp3")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}

		err = conn.WriteMessage(websocket.BinaryMessage, bytes)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			d.Log.Error(err)
			return
		}
	}
}
