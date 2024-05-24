Warehouse-API:
Этот проект представляет собой RESTful API для управления складом. API позволяет резервировать, освобождать и получать информацию о товарах на складе.

База данных имеет таблицы c следующими атрибутами:
package model
```
Table Warehouse{
    id
    name
    is_available
}

Table Product{
    id
    name
    size
    code
    quantity
    warehouse_id
}
```

Запуск приложения
1. Выполните клонирование репозитория "git clone https://github.com/your_username/warehouse-api.git" 
2. В терминале введите "make up". Приложение забилдится и запустится на порту 8080. 
3. Приложение с записями в базе данных для работы - запущено и готово к использованию!

В файле packages.md можно найти аргументацию выбора пакетов.

Описание API методов с работающим запросом и ответом:

1. ReserveProducts POST
Метод ReserveProducts резервирует указанные продукты на складе. Для каждого продукта метод уменьшает количество на складе на 1 (если кол-во 0 - удаляет запись) и увеличивает количество зарезервированных продуктов на 1 (если товара с таким кодом нет - создает новую запись). Если склад недоступен (is_available = FALSE) операции недоступны.

Состояние до запроса:
```
stock
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":95,"warehouse_id":1},{"id":2,"name":"Product 2","size":"Medium","code":"P002","quantity":198,"warehouse_id":1}]
reserved
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":10,"warehouse_id":1},{"id":4,"name":"Product 2","size":"Medium","code":"P002","quantity":2,"warehouse_id":1}]
```
Запрос:
```
 curl -X POST -H "Content-Type: application/json" -d '{"codes":["P001", "P002"]}' http://localhost:8080/reserve
```
Результат:
```
stock
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":94,"warehouse_id":1},{"id":2,"name":"Product 2","size":"Medium","code":"P002","quantity":197,"warehouse_id":1}]
reserved
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":11,"warehouse_id":1},{"id":4,"name":"Product 2","size":"Medium","code":"P002","quantity":3,"warehouse_id":1}]

```

2. ReleaseProducts POST
Метод ReleaseProducts освобождает ранее зарезервированные продукты на складе (отмена резервации). Для каждого продукта метод уменьшает количество зарезервированных продуктов на 1 (если кол-во 0 - удаляет запись) и увеличивает количество продуктов на складе на 1 (если товара с таким кодом нет - создает новую запись). Если склад недоступен (is_available = FALSE) операции недоступны.

Состояние до запроса:
```
stock
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":94,"warehouse_id":1},{"id":2,"name":"Product 2","size":"Medium","code":"P002","quantity":197,"warehouse_id":1}]
reserved
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":11,"warehouse_id":1},{"id":4,"name":"Product 2","size":"Medium","code":"P002","quantity":3,"warehouse_id":1}]
```
Запрос:
```
 curl -X POST -H "Content-Type: application/json" -d '{"codes":["P001", "P002"]}' http://localhost:8080/release
```
Результат:
```
stock
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":95,"warehouse_id":1},{"id":2,"name":"Product 2","size":"Medium","code":"P002","quantity":198,"warehouse_id":1}]
reserved
[{"id":1,"name":"Product 1","size":"Small","code":"P001","quantity":10,"warehouse_id":1},{"id":4,"name":"Product 2","size":"Medium","code":"P002","quantity":2,"warehouse_id":1}]

```

3. GetStock GET
Метод GetStock возвращает список продуктов, находящихся на складе с указанным идентификатором. Метод используется для получения текущего остатка товаров на складе.

Запрос:
```
curl -X GET http://localhost:8080/warehouse/2/stock
```
Результат:
```
[{"id":3,"name":"Product 3","size":"Large","code":"P003","quantity":150,"warehouse_id":2},{"id":4,"name":"Product 4","size":"Small","code":"P004","quantity":250,"warehouse_id":2}]
```

4. GetReservedStock GET
Метод GetReservedStock возвращает список продуктов, которые были зарезервированы на указанном складе.

Запрос:
```
curl -X GET http://localhost:8080/warehouse/3/reserved
```
Результат:
```
[{"id":3,"name":"Product 5","size":"Medium","code":"P005","quantity":3,"warehouse_id":3}]
```
