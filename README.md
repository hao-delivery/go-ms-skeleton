# Скелет приложения на go

# Шаблон архитектуры

**cmd** — содержит корневой файл `root.go` который является точкой подключения всех доступных команд и N кол-во файлов реализации команд;

**config** — содержит файл для указания локальных параметров. Например ip адреса сервиса конфигурации;

**docs** — общая документация по микросервису;

**internal** — все внутренние файлы реализации;

**pkg** — создается на корневом уровне для импортирования в другие приложения;

```
.
├── cmd
│   ├── root.go
│   └── run.go
├── config
│   └── config.toml
├── docs
│   └── about.md
├── internal
│   ├── core
│   │   ├── errors
│   │   ├── types
│   │   └── interfaces
│   ├── config
│   │   └── config.go
│   ├── mocks
│   ├── models
│   ├── services
│   └── utils
├── Makefile
└── README.md
```

## Запуск приложения 

Шаг 1

```
sudo docker-compose build
```

Шаг 2

```
sudo docker-compose up
```