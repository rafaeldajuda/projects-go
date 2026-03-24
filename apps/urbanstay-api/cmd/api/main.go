package main

import (
	"log"
	"urbanstay-api/internal/repository/sql"
)

func init() {

}

func main() {
	// app := fiber.New()

	// repo := memory.NewMemoryRepository()
	repo, err := sql.NewSqlProperty("sqlite3", "./data/teste.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repo.Conn.Close()

	repo.Conn.MustExec("CREATE TABLE person (FirstName TEXT NOT NULL);")
	// _, err = repo.Conn.Exec(`CREATE TABLE teste_sqlite`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("fim")

	// uc := usecase.NewPropertyUseCase(repo)
	// handler := http.NewPropertyHandler(uc)

	// app.Post("/properties", handler.CreateProperty)
	// app.Get("/properties", handler.ListProperties)

	// if err := app.Listen(":8000"); err != nil {
	// 	log.Fatal("server error:", err.Error())
	// }
}
