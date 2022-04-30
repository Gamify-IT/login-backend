package main

import (
	restapi2 "github.com/Gamify-IT/login-backend/src/gen/restapi"
	"github.com/Gamify-IT/login-backend/src/gen/restapi/operations"
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

//go:generate mkdir -p gen/
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate server --target ./gen --name Login --spec ./swagger/swagger.yml --principal interface{} --exclude-main

func main() {

	swaggerSpec, err := loads.Embedded(restapi2.SwaggerJSON, restapi2.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewLoginAPI(swaggerSpec)
	server := restapi2.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "A Todo list application"
	parser.LongDescription = "From the todo list tutorial on goswagger.io"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
