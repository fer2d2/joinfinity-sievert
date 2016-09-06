package cmd

// Available commands
const (
	Root      = "sievert"
	New       = "new"       // generate new sievert project [--override]
	Bootstrap = "bootstrap" // build conf files from sievert.yml specification
	Run       = "run"       // docker-compose up -d [build before]
	Build     = "build"     // [--regenerate (docker-compose kill && docker-compose rm --force && clean-dockers.sh && docker-compose up -d)]
	Delete    = "delete"    // docker-compose rm [service] [--purge (remove IMAGES and containers) --clean-project (purge and remove .sievert directory)]
	Stop      = "stop"      // docker-compose stop [service]
	Kill      = "kill"      // docker-compose kill [service]
	Restart   = "restart"   // docker-compose restart [service] [--hard (docker-compose kill && docker-compose rm && docker-compose up -d)]
	Logs      = "logs"      // docker-compose logs [docker-compose logs --follow]
)
