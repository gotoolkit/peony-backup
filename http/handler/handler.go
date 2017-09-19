package handler

// Handler is a collection of all the service handlers.
type Handler struct {
	DockerHandler *DockerHandler
	UserHandler   *UserHandler
}
