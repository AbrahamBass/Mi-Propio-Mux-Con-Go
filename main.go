package main

import (
	"fmt"
	"net/http"
)

type customeHandler func(http.ResponseWriter, *http.Request)

type MuxFacilito struct {
	rutasFacilitas map[string]customeHandler
}

func (this *MuxFacilito) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := this.rutasFacilitas[r.URL.Path]
	fn(w, r)
}

func (this *MuxFacilito) AddMux(ruta string, fn customeHandler) {
	this.rutasFacilitas[ruta] = fn
}

func main() {

	newMap := make(map[string]customeHandler)
	mux := &MuxFacilito{rutasFacilitas: newMap}

	mux.AddMux("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola como mi propio mux")
	})

	http.ListenAndServe("localhost:3000", mux)
}
