# Скелет приложения на go

# Шаблон архитектуры

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