## Коляда Даниил – ЭФМО-01-25

### Веб приложение для оценки финансовой надежности компании
[tspo-app.com](http://79.164.78.98:5137)

---

### Реализовано

Получение финансовых данных с сервиса Alpha Vantage

|Краткий обзор финансового положения компании|Ежеквартальные отчеты о выручке|
|-|-|
|![Screenshot](./screenshots/Screenshot_1.png)|![Screenshot](./screenshots/Screenshot_2.png)|
|Ежеквартальные отчеты о балансе компании|Ежеквартальные отчеты денежных потоков|
|![Screenshot](./screenshots/Screenshot_3.png)|![Screenshot](./screenshots/Screenshot_4.png)|

---

Подключение к Postgres и запись в базу данных

|![Screenshot](./screenshots/Screenshot_5.png)|![Screenshot](./screenshots/Screenshot_6.png)|
|-|-|
|![Screenshot](./screenshots/Screenshot_7.png)|![Screenshot](./screenshots/Screenshot_8.png)|

---

4 gRPC эндпоинта

- Overview
- IncomeStatement
- BalanceSheet
- CashFlow

|![Screenshot](./screenshots/Screenshot_9.png)|![Screenshot](./screenshots/Screenshot_10.png)|
|-|-|
|![Screenshot](./screenshots/Screenshot_11.png)|![Screenshot](./screenshots/Screenshot_12.png)|

---

Клиент на Vue

|![Screenshot](./screenshots/Screenshot_13.png)|![Screenshot](./screenshots/Screenshot_14.png)|
|-|-|
|![Screenshot](./screenshots/Screenshot_15.png)|![Screenshot](./screenshots/Screenshot_16.png)|

---

Сборка 3-x Docker контейнеров
- Go-server
- Vue-client
- Postgres

Написан [docker-compose](./docker-compose.yml)

![Screenshot](./screenshots/Screenshot_17.png)

---

Настроен [CI/CD Pipeline](./.github/workflows/)  

![Screenshot](./screenshots/Screenshot_18.png)

---

Приложение развернуто на удаленном сервере

![Screenshot](./screenshots/Screenshot_19.png)

---

### Дерево проекта
```
├── .github
│   └── workflows
│       ├── build-and-push-docker.yml
│       ├── publish-oci-artifact.yml
│       ├── deploy-to-server.yml
│       └── main.yml
├── .vscode
│   └── launch.json
├── go-server
│   ├── cmd
│   │   └── main.go
│   ├── db
│   │   ├── db.go
│   │   └── interfaces.go
│   ├── dtos
│   │   ├── balanceSheet.go
│   │   ├── cashFlow.go
│   │   ├── incomeStatement.go
│   │   └── overview.go
│   ├── handlers
│   │   ├── grpc.go
│   │   └── rest.go
│   ├── proto
│   │   ├── gen
│   │   │   ├── data.pb.go
│   │   │   └── data_grpc.pb.go
│   │   ├── rtos
│   │   │   ├── companyRequest.go
│   │   │   └── financialRequest.go
│   │   └── data.proto
│   ├── .dockerignore
│   ├── .env.ex
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── postgres
│   ├── .dockerignore
│   ├── .env.ex
│   ├── Dockerfile
│   └── init.sql
├── screenshots
│   └── ...
├── vue-client
│   ├── .vscode
│   │   └── extensions.json
│   ├── public
│   │   ├── favicon.svg
│   │   └── icons.svg
│   ├── src
│   │   ├── gen
│   │   │   ├── data.client.ts
│   │   │   └── data.ts
│   │   ├── App.vue
│   │   ├── Balance.vue
│   │   ├── CashFlow.vue
│   │   ├── Income.vue
│   │   ├── main.ts
│   │   └── styles.css
│   ├── .dockerignore
│   ├── .env.ex
│   ├── .gitignore
│   ├── Dockerfile
│   ├── README.md
│   ├── index.html
│   ├── nginx.conf.template
│   ├── package-lock.json
│   ├── package.json
│   ├── tsconfig.app.json
│   ├── tsconfig.json
│   ├── tsconfig.node.json
│   └── vite.config.ts
├── .gitignore
├── README.md
├── docker-compose.prof.yml
└── docker-compose.yml

19 directories, 75 files
```