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
	"github.com/gofiber/fiber/v2/middleware/cors"

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
		&entity.Child{},
		&entity.Schedule{},
		&entity.GrowthRecord{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()
	app.Static("/assets", "./assets")
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
	rest.NewChildHandler(v1, childUsecase)

	scheduleRepo := repository.NewScheduleRepository(db)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepo, childRepo)
	rest.NewScheduleHandler(v1, scheduleUsecase)

	growthRepo := repository.NewGrowthRecordRepository(db)
	growthUsecase := usecase.NewGrowthRecordUsecase(growthRepo, childRepo)
	rest.NewGrowthRecordHandler(v1, growthUsecase)

	port := env.GetEnv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", config.AppPort)
	}

	return app.Listen("0.0.0.0:" + port)
}
