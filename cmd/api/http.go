package api

func SetupHTTP() {
	app, err := SetupHTTPApplication()
	if err != nil {
		app.config.log.Fatalf("failed to setup application (http) :%v", err)
	}

	api := app.Mount()
	if err := app.Run(api); err != nil {
		app.config.log.Fatalf("failed to start application (http) :%v", err)
	}
}
