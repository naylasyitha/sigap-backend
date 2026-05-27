package bootstrap

import (
	"fmt"
	"log"
	"os"

	"sigap-backend/app/interface/rest"
	"sigap-backend/app/repository"
	"sigap-backend/app/usecase"
	"sigap-backend/domain/entity"
	"sigap-backend/infra/env"
	"sigap-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() error {
	config, err := env.New()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.MpasiMenu{},
		&entity.MpasiIngredient{},
		&entity.MpasiStep{},
		&entity.MpasiNutrition{},
		&entity.Child{},
		&entity.Schedule{},
		&entity.GrowthRecord{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()
	app.Static("/assets", "/app/assets")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	v1 := app.Group("/api/v1")

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	rest.NewUserHandler(v1, userUsecase)

	menuRepo := repository.NewMenuRepository(db)
	menuUsecase := usecase.NewMenuUsecase(menuRepo)
	rest.NewMenuHandler(v1, menuUsecase)

	childRepo := repository.NewChildRepository(db)
	childUsecase := usecase.NewChildUsecase(childRepo)
	rest.NewChildHandler(v1, childUsecase, middleware.AuthMiddleware)

	scheduleRepo := repository.NewScheduleRepository(db)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepo, childRepo)
	rest.NewScheduleHandler(v1, scheduleUsecase, middleware.AuthMiddleware)

	growthRepo := repository.NewGrowthRecordRepository(db)
	growthUsecase := usecase.NewGrowthRecordUsecase(growthRepo, childRepo)
	rest.NewGrowthRecordHandler(v1, growthUsecase, middleware.AuthMiddleware)

	port := env.GetEnv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", config.AppPort)
	}

	return app.Listen("0.0.0.0:" + port)
}
