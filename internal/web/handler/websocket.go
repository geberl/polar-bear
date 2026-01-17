package handler

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"polar-bear/internal/core"
	"polar-bear/internal/event"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/pod"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Websocket(
	ed event.Distribution,
	store store.Store,
) http.Handler {
	logger := slog.With("component", "handler-websocket")

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			nsName, err := url.QueryUnescape(r.PathValue("ns"))
			if err != nil {
				logger.Error(
					"error decoding namespace name",
					"error", err,
				)
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}

			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				logger.Error("websocket upgrade failed", "err", err)
				return
			}

			logger.Debug("websocket connection established", "ns", nsName)

			updates := make(chan string)
			ed.Register(updates)

			defer func() {
				ed.Unregister(updates)
				conn.Close()
				close(updates)
			}()

			go writer(logger, conn, updates, nsName, store)
			reader(conn)
		},
	)
}

func writer(
	logger *slog.Logger,
	ws *websocket.Conn,
	updates chan string,
	nsName string,
	store store.Store,
) {
	pingTicker := time.NewTicker(pingPeriod)
	var buf bytes.Buffer

	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		case update := <-updates:
			if update == "" {
				return
			}
			logger.Info("got pod update", "key", update)

			podInfos := core.GetPods(store, nsName)
			tc := pod.PodList("", podInfos, "outerHTML")

			buf.Reset()
			err := tc.Render(context.Background(), &buf)
			if err != nil {
				logger.Error("unable to render template", "err", err)
				return
			}

			err = ws.WriteMessage(websocket.TextMessage, buf.Bytes())
			if err != nil {
				logger.Error("unable to write buffer to websocket", "err", err)
				return
			}
		case <-pingTicker.C:
			logger.Debug("send ping")
			_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				logger.Error("unable to write ping to websocket", "err", err)
				return
			}
		}
	}
}

func reader(ws *websocket.Conn) {
	defer ws.Close()

	ws.SetReadLimit(512)
	_ = ws.SetReadDeadline(time.Now().Add(pongWait))

	ws.SetPongHandler(func(string) error {
		_ = ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
