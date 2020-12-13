package rest

import (
	"context"
	"fmt"
	"go_id/services/usecases"
	"net/http"
	"time"
)

type HttpServer interface {
	Run(ctx context.Context) error
}

type httpServer struct {
	ctx     context.Context
	httpSrv *http.Server
	service usecases.Service
}

func NewHttpServer(service usecases.Service) HttpServer {
	var h httpServer
	h.service = service
	return &h
}

func (h *httpServer) Run(ctx context.Context) error {
	ctx, canceller := context.WithCancel(ctx)
	h.ctx = ctx
	go func() {
		<-ctx.Done()
		shutdownCtx, _ := context.WithTimeout(ctx, 5*time.Second)
		if err := h.httpSrv.Shutdown(shutdownCtx); err != nil {
			fmt.Println("failed to shutdown server")
		}
	}()
	defer func() {
		canceller()
	}()
	endpoints := NewEndpoint(h.service)
	httpHandler := newHTTPHandler(endpoints)
	errChan := make(chan error)
	go func() {
		fmt.Sprint("http listener at address, ", 80)
		h.httpSrv = &http.Server{
			Addr:              ":3000",
			Handler:           httpHandler,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
		}
		errChan <- h.httpSrv.ListenAndServe()
	}()
	err := <-errChan
	return err
}
