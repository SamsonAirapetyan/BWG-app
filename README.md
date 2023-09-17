# BWG-app
##Тестовое задание

Нужно реализовать транзакционную систему.

• Invoice -> человеку зачисляются средства по ручке " /invoice" с такими параметрами в теле, как код валюты ("USDT", "RUB", "EUR", etc.), количество средств (число с плавающей точкой), номер кошелька или карты.

• Withdraw -> человек выводит средства со своего баланса по валюте, которую он выбрал по ручке "/withdraw" с такими параметрами в теле, как код валюты, количество средств, номер кошелька или карты куда зачисляются средства.

• В транзакционной системе должны быть статусы транзакции ("Error",
"Success", "Created"). Статусы "Error" и "Success" должны быть финальными.

• Должна быть реализована ручка по получению актуального и замороженного баланса клиентов. Актуальный баланс это тот баланс, который можно вывести. Замороженный баланс - это тот баланс, который, находится в ожидании (со статусом "Created").

• База данных должны быть PostgreSQL.

• Баланс клиента не может уйти ниже нуля.

### Для запуска приложения:

```
make build && make run
```

## Пример работы проекта
#### Запрос через Postman 
![image](https://github.com/SamsonAirapetyan/BWG-app/blob/master/assets/postman_takeoff.png)

#### Пример хранения транзакций в Базе данных
Каждой транзакции присваивается стату (Success и Error являются финальными)

![image](https://github.com/SamsonAirapetyan/BWG-app/blob/master/assets/status.png)

#### Вывод балансов

С учетем ТЗ выводится только тот баланс последний статус которого не является Error

![image](https://github.com/SamsonAirapetyan/BWG-app/blob/master/assets/balance.png)

