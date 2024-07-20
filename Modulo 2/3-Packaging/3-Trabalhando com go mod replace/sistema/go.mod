module github.com/ingrid-paulino/goexpert/sistema

go 1.21.1

//OBS: Essa forma com replace não é a melhor forma de importar um pacote que não esteja no repositorio publicado, pois o go get ou go mod init não irá funcionar. Ao enviar o projeto para o github, o replace deve ser removido e o pacote deve ser importado normalmente.
//é necessario fazer o replace para o pacote achar o caminho correto localmente.

//Isso é uma gambiarra para o pacote achar o caminho correto localmente. na aula 4 será ensinado a forma correta de importar um pacote localmente.
//Comando que gera esse replace: go mod edit -replace github.com/ingrid-paulino/goexpert/package/math=../math   - go mod tidy
//URL relativa para o pacote math --> isso suja o go.mod, use sempre workspaces para evitar isso. Aula 4.
replace github.com/ingrid-paulino/goexpert/package/math => ./../math

require github.com/ingrid-paulino/goexpert/package/math v0.0.0-00010101000000-000000000000
