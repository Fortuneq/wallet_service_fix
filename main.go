package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"wallet-service/domain/controller"
	"wallet-service/domain/repository"
	"wallet-service/domain/service"
)

var db *sqlx.DB

func main() {
	var err error
	dsn := "host=localhost user=wallet_user password=wallet_password dbname=wallet_db sslmode=disable"
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	fmt.Println("Успешное подключение к БД")

	// Инициализация репозиториев
	userRepo := repository.NewPostgresUserRepository(db)
	walletRepo := repository.NewWalletRepository(db)

	// Инициализация схемы
	if err := userRepo.InitializeSchema(); err != nil {
		log.Fatalf("Ошибка создания таблицы users: %v", err)
	}

	if err := walletRepo.InitializeSchema(); err != nil {
		log.Fatalf("Ошибка создания таблицы wallets: %v", err)
	}

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	walletService := service.NewWalletService(walletRepo)

	app := &controller.App{
		UserService:   *userService,
		WalletService: *walletService,
	}

	app.Run()
}
