package main

import (
	"log"
	"os"
	"project-root/config"
	"project-root/db/seeders"
)

func main() {
	config.InitEnv()

	if len(os.Args) < 2 {
		log.Fatal("Command required (example: seed)")
	}

	command := os.Args[1]

	switch command {
	case "seed":
		runSeed()
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func runSeed() {
	db := config.InitDB()

	list := []seeders.Seeder{
		seeders.BookSeeder{},
	}

	for _, s := range list {
		log.Printf("🌱 Running %T\n", s)

		if err := s.Run(db); err != nil {
			log.Fatalf("Seeder failed: %v", err)
		}
	}

	log.Println("✅ Seeding completed")
}
