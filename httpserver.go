package httpServer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/chrislws/homework/httpserver"
)

func HttpServer() {
	// create a new *ServMux object
	mux := http.NewServeMux()
	// register routings
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/healthz", handler.Healthz)
	mux.HandleFunc("/metrics",promhttp.Handler())
	serv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	// start http server with gorouting
	go func() {
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start http server failed: %s\n", err)
		}
	}()
	// grace shutdown
	quit := make(chan os.Signal, 1)
	// receive system signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block
	// service will be shut down in 5 seconds, wait for the request to be processed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown server failed: %s", err)
	}
	log.Println("server shutdown successfully")
}
func randInt(min int, max int) int {
        rand.Seed(time.Now().UTC().UnixNano())
        return min + rand.Intn(max-min)
}

func rootHandler(w http.ResponseWriter, r *http.Request)
    timer := metrics.NewTimer()
    defer timer.ObserveTotal()
    user  := r.URL.Query().Get("user")
    delay := randInt("10","2000)
    time.Sleep(time.Millisecond*time.Duration(delay))
