```txt
/app-agendamento
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada (Bootstrapping e DI)
├── internal/
│   ├── entity/              # Camada de Domínio (Modelos e Regras de Ouro)
│   │   └── appointment.go
│   ├── usecase/             # Camada de Aplicação (Lógica de Negócio)
│   │   ├── appointment_uc.go
│   │   └── interfaces.go    # Definição do que o Repository deve fazer
│   ├── infra/               # Camada de Adaptadores (Mundo Externo)
│   │   ├── database/        # Implementação SQL (Gorm, SQLC ou purista)
│   │   │   └── appointment_repository.go
│   │   └── web/             # Handlers HTTP e Rotas
│   │       ├── handlers/
│   │       └── middlewares/
├── pkg/                     # Utilitários globais (Logger, Validação)
├── api/                     # Documentação (Swagger/OpenAPI)
├── deployments/             # Configs de Docker e Nginx
├── sql/
│   └── migrations/          # Versionamento do banco (.sql)
├── go.mod
├── go.sum
└── docker-compose.yml
```