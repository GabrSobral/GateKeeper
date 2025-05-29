package main

import (
	"log/slog"
	"net/http"

	_ "github.com/gate-keeper/cmd/server/docs"
	"github.com/gate-keeper/internal/infra/database"
	"github.com/gate-keeper/internal/presentation/http/routing"
	"github.com/joho/godotenv"
)

//	@title			GateKeeper API
//	@version		1
//	@description	This is the GateKeeper API documentation.
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//
// @host		localhost:8080
// @BasePath	/v1
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	pool, err := database.NewConnectionPool()

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	router := routing.SetHttpRoutes(pool)

	slog.Info("✅ Server is running on port 8080")

	http.ListenAndServe(":8080", router)
}
