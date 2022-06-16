# Эмулятор платежного сервиса (Golang)

### Запуск

docker-compose up --build

URL сервера: http://localhost:3200/
### Эндпоинты
- #### /api/transactions GET (Получение всех транзакций)
Пример запроса: <br>
http://localhost:3200/api/transactions <br>
Пример ответа: <br>
`[{
  "id": 1,
  "user_id": 1,
  "user_email": "email2@email.com",
  "total": "324442",
  "currency_id": 1,
  "status_id": 4,
  "created_at": "2022-06-16T03:44:18.233577Z",
  "updated_at": "2022-06-16T03:44:18.233577Z"
  },
]`
- #### /api/transactions/status/{id} GET (Проверка статуса платежа)
Пример запроса: <br>
http://localhost:3200/api/transactions/status/1 <br>
Пример ответа: <br>
`  {
  "transaction_id": 1,
  "title": "ERROR"
  }`
- #### /api/transactions/{id} GET (Получение транзакции)
Пример запроса: <br>
http://localhost:3200/api/transactions/1 <br>
Пример ответа: <br>
`{
"id": 1,
"user_id": 1,
"user_email": "email2@email.com",
"total": "324442",
"currency_id": 1,
"status_id": 4,
"created_at": "2022-06-16T03:44:18.233577Z",
"updated_at": "2022-06-16T03:44:18.233577Z"
}`
- #### /api/transactions/user-id/{id} GET (Получение транзакций пользователя по ID)
Пример запроса: <br>
http://localhost:3200/api/transactions/user-id/1 <br>
Пример ответа: <br>
`[
{
"id": 1,
"user_id": 1,
"user_email": "email2@email.com",
"total": "324442",
"currency_id": 1,
"status_id": 4,
"created_at": "2022-06-16T03:44:18.233577Z",
"updated_at": "2022-06-16T03:44:18.233577Z"
},]`
- #### /api/transactions/user-email/{email} GET (Получение транзакций пользователя по email)
Пример запроса:<br>
http://localhost:3200/api/transactions/user-email/wrong-email <br>
Пример ответа: <br>
"Not found"
- #### /api/transactions/close/{id} PUT (Отмена транзакции)
Пример запроса:<br>
http://localhost:3200/api/transactions/close/1 <br>
Пример ответа: <br>
`{
"id": 1,
"user_id": 1,
"user_email": "email2@email.com",
"total": "324442",
"currency_id": 1,
"status_id": 5, 
"created_at": "2022-06-16T03:44:18.233577Z",
"updated_at": "2022-06-16T03:44:18.233577Z"
}`
- #### /api/transactions/new POST (Создать транзакцию)
Пример запроса:<br>
{<br>
"user_id":123,<br>
"user_email":"email223@email.com",<br>
"total":324412312342,<br>
"currency_id":2<br>
}<br>
Пример ответа: <br>
"success"
### Изменение статуса платежной системой
Для изменения статуса, необходимо авторизоваться: 
- #### /api/login POST (Авторизация)
Для получения токена авторизации, выполните запрос:<br>
{ <br>
"for_token":1<br>
}<br>
В ответе, придет токен (пример токена): <br>
`"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTUzNjQzNjksInVzZXIiOiJVc2VyIn0.3lpBdnu3uUXOFk12sLTdsPBt_QcdFKBGCrLg-Te9JSE"`
- #### /api/transactions/change-status PUT (Изменение статуса платежной системой)
1. Добавьте токен в заголовки запроса: <br>
Key - Token, Value - полученный токен <br>
2. Пример запроса:<br>
{<br>
"transaction_id":1,<br>
"status_id":3<br>
}<br>
 Пример ответа: <br>
`{
"id": 1,
"user_id": 1,
"user_email": "email2@email.com",
"total": "324442",
"currency_id": 1,
"status_id": 3,
"created_at": "2022-06-16T03:44:18.233577Z",
"updated_at": "2022-06-16T04:12:51.104281Z"
}`
### Тесты
Тесты написаны для GET запросов, файл main_test.go. <br>
запуск тестов: <br> cd app <br> go test
