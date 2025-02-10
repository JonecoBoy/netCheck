package nina

import ninaRouter "github.com/JonecoBoy/nina/router"

func NewRouter() *ninaRouter.ServeMux {
	return ninaRouter.NewRouter()
}
