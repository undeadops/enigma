package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	libhoney "github.com/honeycombio/libhoney-go"
)

func init() {
	hcConfig := libhoney.Config{
		WriteKey: os.Getenv("HONEY_KEY"),
		Dataset:  "myData",
	}

	// This will ensure that our libhoney events get printed to the
	// console. This allows for easier iterating and debugging of
	// instrumentation.
	if os.Getenv("ENV") != "prod" {
		hcConfig.Output = &libhoney.WriterOutput{}
	}

	if err := libhoney.Init(hcConfig); err != nil {
		log.Print(err)
		os.Exit(1)
	}

	if hnyTeam, err := libhoney.VerifyWriteKey(hcConfig); err != nil {
		log.Print(err)
		log.Print("Please make sure the HONEYCOMB_WRITEKEY environment variable is set.")
		os.Exit(1)
	} else {
		log.Print(fmt.Sprintf("Sending Honeycomb events to the %q dataset on %q team", hnyDatasetName, hnyTeam))
	}

	// Initialize fields that every sent event will have.

	// Getting hostname on every event can be very useful if, e.g., only a
	// particular host or set of hosts are the source of an issue.
	if hostname, err := os.Hostname(); err == nil {
		libhoney.AddField("system.hostname", hostname)
	}
	libhoney.AddDynamicField("runtime.num_goroutines", func() interface{} {
		return runtime.NumGoroutine()
	})
	libhoney.AddDynamicField("runtime.memory_inuse", func() interface{} {
		// This will ensure that every event includes information about
		// the memory usage of the process at the time the event was
		// sent.
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return mem.Alloc
	})
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// // Logger return log message
// func Logger() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(time.Now(), r.Method, r.URL)
// 		router.ServeHTTP(w, r) // dispatch the request
// 	})
// }
