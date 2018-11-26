package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/creedasaurus/kalkulator/stacker"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	serverShutdownTimeout = 10 * time.Second
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/compute", handleCompute)

	serv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		err := serv.ListenAndServe()
		switch err {
		case nil:
		case http.ErrServerClosed:
			fmt.Println(err)
		}
	}()

	switch <-signalChannel {
	case syscall.SIGQUIT:
		fmt.Println("SIGQUIT - exiting")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM - exiting")
	case syscall.SIGINT:
		fmt.Println("SIGINT - exiting")
	case syscall.SIGHUP:
		fmt.Println("SIGHUP - exiting")
	default:
		fmt.Println("Signal received - exiting")
	}

	tShutdown := time.NewTimer(serverShutdownTimeout)
	shutdownChannel := make(chan error)

	go func() {
		shutdownChannel <- serv.Shutdown(ctx)
	}()

	select {
	case err := <-shutdownChannel:
		if err != nil {
			fmt.Println("shutdown finished with an error:", err)
		} else {
			fmt.Println("shutdown finished successfully.")
		}
	case <-tShutdown.C:
		fmt.Println("shutdown timed out")
	}

	cancelFunc()

	fmt.Println("exiting real nice")
}

type ComputeRequest struct {
	Expression string `json:"expression"`
}

type ComputeResponse struct {
	Answer float64 `json:"answer"`
}

func handleCompute(w http.ResponseWriter, r *http.Request) {
	algo := ComputeRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&algo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// timer to test out
	t := time.NewTimer(time.Second * 3)
	<-t.C

	ans, err := stacker.ProcessStatement(algo.Expression)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := ComputeResponse{Answer: ans}

	respMarshalled, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(respMarshalled)
}
