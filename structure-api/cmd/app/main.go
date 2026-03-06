package main

import (
	"log"
	"net/http"
	"org-structure-api/internal/db"
	"org-structure-api/internal/handler"
	"org-structure-api/internal/repository"
	"org-structure-api/internal/router"
	"org-structure-api/internal/service"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func main() {

	database := db.NewDB()

	repo := repository.NewDepartmentRepository(database)
	service := service.NewDepartmentService(repo)
	handler := handler.NewDepartmentHandler(service)

	r := router.NewRouter(handler)

	log.Println("server started at :8080")

	http.ListenAndServe(":8080", r)
}
