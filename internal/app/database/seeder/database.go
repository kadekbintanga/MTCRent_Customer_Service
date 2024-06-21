package seeder

type DatabaseSeeder interface {
	Seed()
}

type data struct {
	DatabaseSeeder
}

func Run() {
	seeders := []data{}

	for _, seeder := range seeders {
		seeder.Seed()
	}
}
