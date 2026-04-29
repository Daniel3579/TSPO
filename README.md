## Коляда Даниил – ЭФМО-01-25

### Веб приложение для оценки финансовой надежности компании

---

### Реализовано

Получение финансовых данных с сервиса Alpha Vantage

- Краткий обзор финансового положения компании
![Screenshot](./screenshots/Screenshot_1.png)

- Ежеквартальные отчеты о выручке
![Screenshot](./screenshots/Screenshot_2.png)

- Ежеквартальные отчеты о балансе компании
![Screenshot](./screenshots/Screenshot_3.png)

- Ежеквартальные отчеты денежных потоков
![Screenshot](./screenshots/Screenshot_4.png)

---

Подключение к Postgres и запись в базу данных
![Screenshot](./screenshots/Screenshot_5.png)
|![Screenshot](./screenshots/Screenshot_6.png)|![Screenshot](./screenshots/Screenshot_7.png)|![Screenshot](./screenshots/Screenshot_8.png)|
|-|-|-|

---

4 gRPC эндпоинта

- Overview
- IncomeStatement
- BalanceSheet
- CashFlow

|![Screenshot](./screenshots/Screenshot_9.png)|![Screenshot](./screenshots/Screenshot_10.png)|
|-|-|
|![Screenshot](./screenshots/Screenshot_11.png)|![Screenshot](./screenshots/Screenshot_12.png)|

### В разрботке

Beб интерфейс

---

### Дерево проекта
```
├── .vscode
│   └── launch.json
├── cmd
│   └── main.go
├── db
│   ├── db.go
│   ├── interfaces.go
│   └── tables.sql
├── dtos
│   ├── balanceSheet.go
│   ├── cashFlow.go
│   ├── incomeStatement.go
│   └── overview.go
├── handlers
│   └── grpc.go
├── proto
│   ├── gen
│   │   ├── data.pb.go
│   │   └── data_grpc.pb.go
│   ├── rtos
│   │   ├── companyRequest.go
│   │   └── financialRequest.go
│   └── data.proto
├── screenshots
│   └── ...
├── .env_example
├── .gitignore
├── README.md
├── go.mod
└── go.sum

10 directories, 32 files
```