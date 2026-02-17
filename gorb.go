package gorb

import (
	"fmt"
	"net/http"
	"log"
	"github.com/Kaurikallaste/gorb/utils"
)

type Registerer func()

func Bootstrap(registerers []Registerer) {

	for _, registerer := range registerers {
		registerer()
	}

	port := utils.GetEnv().PORT

	if len(port) == 0 {
		port = "8080"
	}

	log.Print(fmt.Sprint("Starting server on port: ", port))
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
