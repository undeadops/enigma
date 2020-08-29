package config

import (
	"context"
	"log"
	"net/http"
	"time"

	libhoney "github.com/honeycombio/libhoney-go"
)

var (
	hostname       string
	hnyDatasetName = "undead.questions"
	hnyContextKey  = "enigmaEvent"
)

type HoneyResponseWriter struct {
	*libhoney.Event
	http.ResponseWriter
	StatusCode int
}

func (hrw *HoneyResponseWriter) WriteHeader(status int) {
	// Mark this down for adding to the libhoney event later.
	hrw.StatusCode = status
	hrw.ResponseWriter.WriteHeader(status)
}

func addRequestProps(req *http.Request, ev *libhoney.Event) {
	// Add a variety of details about the HTTP request, such as user agent
	// and method, to any created libhoney event.
	ev.AddField("request.method", req.Method)
	ev.AddField("request.path", req.URL.Path)
	ev.AddField("request.host", req.URL.Host)
	ev.AddField("request.proto", req.Proto)
	ev.AddField("request.content_length", req.ContentLength)
	ev.AddField("request.remote_addr", req.RemoteAddr)
	ev.AddField("request.user_agent", req.UserAgent())
}

// HoneycombMiddleware will wrap our HTTP handle funcs to automatically
// generate an event-per-request and set properties on them.
func HoneycombMiddleware(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// We'll time each HTTP request and add that as a property to
		// the sent Honeycomb event, so start the timer for that.
		startHandler := time.Now()
		ev := libhoney.NewEvent()

		defer func() {
			if err := ev.Send(); err != nil {
				log.Print("Error sending libhoney event: ", err)
			}
		}()

		addRequestProps(r, ev)

		// Create a context where we will store the libhoney event. We
		// will add default values to this event for every HTTP
		// request, and the user can access it to add their own
		// (powerful, custom) fields.
		ctx := context.WithValue(r.Context(), hnyContextKey, ev)
		reqWithContext := r.WithContext(ctx)

		honeyResponseWriter := &HoneyResponseWriter{
			Event:          ev,
			ResponseWriter: w,
			StatusCode:     200,
		}

		fn(honeyResponseWriter, reqWithContext)

		ev.AddField("response.status_code", honeyResponseWriter.StatusCode)
		handlerDuration := time.Since(startHandler)
		ev.AddField("timers.total_time_ms", handlerDuration/time.Millisecond)
	}
}
