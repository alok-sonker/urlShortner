package main

func execute() error {
	// Initialise gin
	server := initServer()

	// Load HTML
	loadViews(server)

	// Initialise routes
	initRoutes(server)

	// Start Server
	if err := run(server); err != nil {
		return err
	}
	return nil
}
