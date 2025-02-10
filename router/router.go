package router

import (
	"fmt"
	"github.com/JonecoBoy/netCheck/net"
	ninaRouter "github.com/JonecoBoy/nina/router"
	"net/http"
)

func NewRouter() *ninaRouter.ServeMux {
	roteador := ninaRouter.NewRouter()

	roteador.GET("/hello/{id}/{abc}", helloHandler, nil)
	roteador.GET("/ping", pingHandler, nil)

	adminGroup := roteador.GROUP("/admin", nil, nil)
	adminGroup.GET("/", func(w http.ResponseWriter, r *ninaRouter.NinaRequest) { fmt.Println("Request to / from group") }, nil)
	adminGroup.GET("/hello", helloHandler, nil)
	adminGroup.POST("/hello", func(w http.ResponseWriter, r *ninaRouter.NinaRequest) {
		fmt.Fprint(w, "Hello from post")
	}, []ninaRouter.Middleware{})

	return roteador
}

func helloHandler(w http.ResponseWriter, r *ninaRouter.NinaRequest) {
	fmt.Println("Request to /hello")
	jonas := r.PathValue("id")
	fmt.Fprint(w, "Hello from get")
	fmt.Printf("Hello from /hello/%s", jonas)
}

func pingHandler(w http.ResponseWriter, r *ninaRouter.NinaRequest) {
	url := r.PathValue("url")
	if url == "" {
		url = r.URL.Query().Get("url")
	}
	if net.Ping(url) {
		fmt.Fprintf(w, "The webpage %s is online.\n", url)
	} else {
		fmt.Fprintf(w, "The webpage %s is offline.\n", url)
	}
}
