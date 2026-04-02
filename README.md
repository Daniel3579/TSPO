## Коляда Даниил – ЭФМО-01-25

### Веб приложение для оценки финансовой надежности компании

---

### Реализовано

- Получение финансовых данных с сервиса Alpha Vantage

    А именно:
    - Краткий обзор финансового положения компании
    ![Screenshot](./screenshots/Screenshot_1.png)

    - Ежеквартальные отчеты о выручке
    ![Screenshot](./screenshots/Screenshot_2.png)

    - Ежеквартальные отчеты о балансе компании
    ![Screenshot](./screenshots/Screenshot_3.png)

    - Ежеквартальные отчеты денежных потоков
    ![Screenshot](./screenshots/Screenshot_4.png)

---

- Подключение к Postgres и запись в базу данных
    ![Screenshot](./screenshots/Screenshot_5.png)
    |![Screenshot](./screenshots/Screenshot_6.png)|![Screenshot](./screenshots/Screenshot_7.png)|![Screenshot](./screenshots/Screenshot_8.png)|
    |-|-|-|

---

### В разрботке

gRPC API

---

### Дерево проекта
```
├── .gitignore
├── .env_example
├── README.md
├── cmd
│   └── main.go
├── db
│   ├── db.go
│   ├── interfaces.go
├── dtos
│   ├── balanceSheet.go
│   ├── cashFlow.go
│   ├── incomeStatement.go
│   └── overview.go
├── go.mod
├── go.sum
├── handlers
│   └── grpc.go
├── proto
│   ├── data.proto
│   ├── dtos
│   │   └── companyRequest.go
│   └── gen
│       ├── data.pb.go
│       └── data_grpc.pb.go
└── screenshots
    ├── ...

8 directories, 25 files
```