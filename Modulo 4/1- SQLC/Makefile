# Ajuda a executar comandos no projeto
createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

# Ignora se tiver algum arquivo com o nome passado e roda só o comando a cima ao inves do arquivo
.PHONY: migrate migratedown createmigration

# para rodar o comando
	# make createmigration
	# make migrate
	# make migratedown
