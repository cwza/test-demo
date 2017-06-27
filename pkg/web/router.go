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
	Router.Get("/seq/getValueByStep/{step}", seqGetValueByStep)
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

func seqGetValueByStep(w http.ResponseWriter, req *http.Request) {
	step, err := strconv.Atoi(chi.URLParam(req, "step"))
	if err != nil {
		fmt.Fprint(w, "step should be an integer")
		return
	}
	seqService := service.NewSeqService(nil)
	value := seqService.GetValueByStep(step)
	fmt.Fprintf(w, strconv.Itoa(value))
}
