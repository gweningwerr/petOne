# Pet project on Golang

#### Библиотеки для работы с БД

* database/sql - стандартная
* sqlx - ползволяет писать сложные запросы
* gosql - потенциально тоже мощная надо изучить

#### Инструмент миграции golang-migrate/migrate
Подробнее об инструменте https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

От сюда берём нужную нам версию: https://github.com/golang-migrate/migrate/releases

Команда инициализации
```bash 
migrate create -seq -ext -digits sql -dir migrations pet_one
```

Применение миграции port=8932 user=postgres password=pass dbname=pet_one
```bash
# добавляет миграцию
migrate -path migrations -database "postgres://localhost:8932/pet_one?sslmode=disable&user=postgres&password=pass" up
```

```bash
# убирает миграцию
migrate -path migrations -database "postgres://localhost:8932/pet_one?sslmode=disable&user=postgres&password=pass" down
```

```sql
-- up
CREATE TABLE users
(
id            INTEGER     NOT NULL PRIMARY KEY,
name          VARCHAR     NOT NULL,
email         VARCHAR     NOT NULL UNIQUE,
age           INTEGER     NOT NULL,
password_hash VARCHAR(32) NOT NULL
);

--down
DROP TABLE users;
```

