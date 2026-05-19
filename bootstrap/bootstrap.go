package bootstrap

import (
	"fmt"
	"log"

	"sigap-backend/app/interface/rest"
	"sigap-backend/app/repository"
	"sigap-backend/app/usecase"
	"sigap-backend/domain/entity"
	"sigap-backend/infra/env"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() error {
	config, err := env.New()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.MpasiMenu{},
		&entity.MpasiIngredient{},
		&entity.MpasiStep{},
		&entity.MpasiNutrition{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()
	// app.Use(logger.New())
	// app.Use(cors.New())

	v1 := app.Group("/api/v1")

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	rest.NewUserHandler(v1, userUsecase)

	menuRepo := repository.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepo)
	rest.NewMenuHandler(v1, menuUsecase)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
