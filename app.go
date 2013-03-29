package goes

import (
	"fmt"
	"net/http"
	"reflect"
	"github.com/gorilla/mux"
)

type RouteHandler struct {
	Handler interface{}
}

func (h RouteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handlerType := reflect.TypeOf(h.Handler)
	handlerValue := reflect.ValueOf(h.Handler)
	handlerArgs := []reflect.Value{}

	if handlerType.NumIn() == 1 {
		handlerStruct := reflect.New(handlerType.In(0)).Interface()	
		if err := mapToStruct(handlerStruct, mux.Vars(req)); err != nil {
			panic(fmt.Sprintf("Error converting params to handler struct: %s", err.Error()))
		}
		handlerArgs = append(handlerArgs, reflect.ValueOf(handlerStruct).Elem())
	}

	handlerResponse := handlerValue.Call(handlerArgs)
	fmt.Fprint(w, handlerResponse[0])
}

type App struct {
	Router mux.Router
}

func (a *App) Route(pat string, h interface{}) {
	a.Router.Handle(pat, RouteHandler{Handler:h})
}

func (a *App) Run(bind string, port int) {
	bindTo := fmt.Sprintf("%s:%d", bind, port)
	http.Handle("/", &a.Router)
	http.ListenAndServe(bindTo, &a.Router)
}
