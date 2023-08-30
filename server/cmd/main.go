package main

import (
	"github.com/joomcold/go-next-docker/internal/configs"
	"github.com/joomcold/go-next-docker/internal/database"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func init() {
	initializers.Environment()
	initializers.Postgresql()
	database.Migrations()
}

func main() {
	configs.Routes()
}
