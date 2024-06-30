package main

import (
	"os"

	"github.com/JamesCCoder/resume_backend_go/config"
	"github.com/JamesCCoder/resume_backend_go/routes"
)

func main() {

	
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    config.LoadConfig(env)
    config.ConnectMongoDB()

    router := routes.SetupRouter()

    router.SetTrustedProxies(nil)

    router.Run(":8085")
}
