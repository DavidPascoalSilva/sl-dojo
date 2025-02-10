package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	log.Println("Starting")

	// Conectar ao MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/sl_dojo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrar estrutura para o banco
	db.AutoMigrate(&Person{})


	// Criar uma instancia Fiber
	app := fiber.New()

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	// Criar uma nova pessoa via API
	app.Post("/people", func(c *fiber.Ctx) error {
		person := new(Person)
		if err := c.BodyParser(person); err != nil {
			return c.Status(400).SendString("Erro ao processar")
		}
		db.Create(person)
		return c.SendString(fmt.Sprintf("Nome %s", person.Name))
	})

	// person := Person{Name: "Zé das Couves", Age: 45}
	// db.Create(&person)

	// Iniciar o servidor
	log.Fatal(app.Listen(":3000"))
}