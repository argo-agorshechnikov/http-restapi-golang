# http-restapi-golang

go mod init github.com/argo-agorshechnikov/http-restapi-golang - инициализация go модулей, чтобы управлять зависимостями 

main.go - точка входа в программу, 

Makefile - используется, чтобы не запоминать все возможные команды для программы

http-restapi-golang/internal/app/apiserver - пакет, содержащий apiserver, функция New() возвращает сконфигурированный инстанс структуры APIServer, функция Start() - запускает http сервер, подключает к БД и тд


Чтобы не хардкодить различные данные(порты для сервака, url ...), нужно научить apiserver конфигурироваться для этого можно юзать toml файлы. Для этого создан файл internal/app/apiserver/config.go, структура конфига содержит параметры, которые будут сконфигурены, функция NewConfig() возвращает инициализированный конфиг с дефолтными параметрами

apiserver.toml - файл, содержащий сам конфиг

---

# Зависимости отображаются также в go.mod(require)
1) BurntSushi/toml - для парсинга конфиг файла (go get github.com/BurntSushi/toml)

---
