## Коляда Даниил – ЭФМО-01-25

### Веб приложение для оценки финансовой надежности компании

### Реализовано

Получение финансовых данных с сервиса Alpha Vantage

А именно:
- Краткий обзор финансового положения компании

![Screenshot](./screenshots/Screenshot_1.png)

- Ежеквартальные отчеты о выручке

![Screenshot](./screenshots/Screenshot_2.png)

- Ежеквартальные отчеты о балансе компании

![Screenshot](./screenshots/Screenshot_3.png)

- Ежеквартальные отчеты денежных потоков

![Screenshot](./screenshots/Screenshot_4.png)

### В разрботке
Подключение к Postgres и запись в базу данных

### Дерево проекта
```
├── README.md
├── cmd
│   └── main.go
├── db
│   └── db.go
├── dtos
│   ├── dtos.go
│   ├── incomeStatement.go
│   ├── interfaces.go
│   ├── overview.go
├── go.mod
├── proto
│   ├── data.pb.go
│   ├── data.proto
│   └── data_grpc.pb.go
└── screenshots
    ├── ...

6 directories, 15 files
```