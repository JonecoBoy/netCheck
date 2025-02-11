package nina

import ninaRouter "github.com/jonecoboy/nina/router"

func NewRouter() *ninaRouter.ServeMux {
	return ninaRouter.NewRouter()
}
