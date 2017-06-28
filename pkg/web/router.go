package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cwza/test-demo/pkg/service"
	"github.com/go-chi/chi"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()
	Router.Get("/health", health)
	Router.Get("/seq/getValue", seqGetValue)
	Router.Get("/seq/reset", seqReset)
	Router.Get("/seq/getNextByStep/{step:^[1-9]\\d*$}", seqGetNextByStep)
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func seqGetValue(w http.ResponseWriter, req *http.Request) {
	seqService := service.NewSeqService(nil)
	fmt.Fprintf(w, strconv.Itoa(seqService.GetValue()))
}

func seqReset(w http.ResponseWriter, req *http.Request) {
	seqService := service.NewSeqService(nil)
	seqService.Reset()
	return
}

func seqGetNextByStep(w http.ResponseWriter, req *http.Request) {
	step, _ := strconv.Atoi(chi.URLParam(req, "step"))
	seqService := service.NewSeqService(nil)
	value := seqService.GetNextByStep(step)
	fmt.Fprintf(w, strconv.Itoa(value))
}
