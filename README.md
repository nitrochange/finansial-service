## Financial-Service
Приложение Finansial-Service реализует тестовое задание на позицию стажера в Авито в 2021 году.
Смысл сервиса заключается в управлении балансом пользователей, в нем присутствуютв возможности 
пополнения, списания, получения баланса, перевода.Файл с SQL сущностями приложен, однако он использовался на ранних этапах написания. В настоящий момент
таблицы создаются в базе на старте при помощи ORM. Сервис написан с использованием
`gin`, в качестве драйвера к базе использовался `pq`. Использовалась база данных Postgres
, поднимаемая из докер контейнера. Порядок работы с приложением примерно следующий: 
в первую очередь после поднятия приложения следует удостовериться при помощи
`health-check`(Пример см ниже), что приложение действительно запущено. Следующим шагом
является добавления в базу пользователей при помощи `createUser`(Пример см также ниже).
Возвращаемые id при создании пользователей являются их идентификаторами, их следует
использовать в остальных запросах в полях id. Hope you will enjoy my app:)
### Запуск приложения

База данных Postgres поднимается из Docker контейнера на порте 5432
при помощи следующей команды.
```sh
docker-compose -f postgres.yml up
```
Приложение запускается командой из папки main, поднимается на порте 8080 (в будущем
планируется Dockerfile). Для того, чтобы не было проблем с неподченными библиотеками
лучше запускать из среды `Goland`
```shell
go run main.go
```

### Примеры запросов
Во время разработки и тестирования использовался Postman, поэтому примеры запросов
будут приведены для него:
#### Создание пользователя(POST)
```shell
localhost:8080/createUser
```
```shell
{
    	"FirstName":    "FirstName",
	"SecondName":   "SecondName",
	"Email":        "Email",
	"PhoneNumber":  "PhoneNumber",
	"Balance":      "Balance",
	"Address": 	"Address",

}
```
Успешный ответ:
```shell
{
    "userId": "c3eb22bc-4dd2-428e-91d3-efe658e84539"
}
```
#### Пополнение баланса(GET)
В приложении подразумевается что числа при пополнении и снятии передаются без знака!
```shell
http://localhost:8080/addition
```
```shell
{
    "id": "c3eb22bc-4dd2-428e-91d3-efe658e84539",
    "amount": "999"
}
```
Успешный ответ:
```shell
{
    "message": "Addition completed"
}
```
#### Снятие средств(GET)
```shell
http://localhost:8080/write-down
```
```shell
{
    "id": "85540221-db1c-4821-8022-66e0619df56b",
    "amount": "9"
}
```
Успешный ответ
```shell
{
    "message": "Write-down completed"
}
```
#### Запрос баланса(GET)
```shell
http://localhost:8080/getBalance
```
```shell
{
    "id":"85540221-db1c-4821-8022-66e0619df56b"
}
```
Успешный ответ:
```shell
{
    "balance": "582",
    "firstName": "Misha",
    "id": "85540221-db1c-4821-8022-66e0619df56b",
    "secondName": "Tihonov"
}
```
#### Выполнение перевода(POST)
```shell
http://localhost:8080/transact
```
```shell
{
    "sender_user_id": "dd53108b-93d7-40ef-8aa1-c70a73a89f79",
    "receiver_user_id": "85540221-db1c-4821-8022-66e0619df56b",
    "Amount":"100"
}
```
Успешный ответ:
```shell
{
    "message": "Addition completed"
}{
    "message": "Write-down completed"
}{
    "message": "Transaction completed"
}
```
#### Health check(GET)
Для проверки работоспособности приложения реализован endpoint health-check
```shell
http://localhost:8080/
```
Если приложение запущено:
```shell
{
    "message": "App is running successfully"
}
```

### Структура БД
База данных состоит их двух таблиц(USERS, TRANSACTIONS). Их структура описана 
в модуле с моделями.

