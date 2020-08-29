package questions

import (
	"net/http"
	"os"
	"strings"

	libhoney "github.com/honeycombio/libhoney-go"
)

func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		// Get token from the Authorization header
		// format: Authorization: Bearer
		bearer := r.Header.Get("authorization")
		if bearer != "" {
			token = strings.TrimPrefix(bearer, "Bearer ")
		}

		// If the token is empty...
		if token == "" {
			// If we get here, the required token is missing
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if token != "" {
			authToken := os.Getenv("AUTH_TOKEN")
			if token != authToken {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func hnyEventFromRequest(r *http.Request) *libhoney.Event {
	ev, ok := r.Context().Value("questionsHandlers").(*libhoney.Event)
	if !ok {
		// We control the way this is being put on context anyway.
		panic("Couldn't get libhoney event from request context")
	}

	return ev
}

func addFinalErr(err *error, ev *libhoney.Event) {
	if *err != nil {
		ev.AddField("error", (*err).Error())
	}
}
