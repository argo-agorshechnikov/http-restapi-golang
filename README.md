# http-restapi-golang

go mod init github.com/argo-agorshechnikov/http-restapi-golang - инициализация go модулей, чтобы управлять зависимостями 

main.go - точка входа в программу, 

Makefile - используется, чтобы не запоминать все возможные команды для программы

http-restapi-golang/internal/app/apiserver - пакет, содержащий apiserver, функция New() возвращает сконфигурированный инстанс структуры APIServer, функция Start() - запускает http сервер, подключает к БД и тд


