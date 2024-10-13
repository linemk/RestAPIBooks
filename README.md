Вот задание, которое поможет закрепить знания по Go и подготовиться к собеседованию:

### Задание: Разработка API для управления библиотекой книг

**Цель**: Создать RESTful API для управления библиотекой книг с помощью Go и базы данных SQLite. Тебе нужно будет реализовать операции создания, чтения, обновления и удаления (CRUD) для книг и авторов.

#### Функциональные требования:
1. **Сущности**:
    - **Book** (Книга) с полями:
        - `id` (целое число, первичный ключ, автоинкремент)
        - `title` (строка)
        - `author_id` (целое число, внешний ключ на таблицу `authors`)
        - `published_year` (целое число)
    - **Author** (Автор) с полями:
        - `id` (целое число, первичный ключ, автоинкремент)
        - `name` (строка)
        - `birth_year` (целое число)

2. **Методы API**:
    - **/books**:
        - `POST`: Добавление новой книги.
        - `GET`: Получение списка всех книг с фильтрацией по `author_id`.
    - **/books/{id}**:
        - `GET`: Получение информации о книге по её `id`.
        - `PUT`: Обновление данных книги.
        - `DELETE`: Удаление книги.
    - **/authors**:
        - `POST`: Добавление нового автора.
        - `GET`: Получение списка всех авторов.
    - **/authors/{id}**:
        - `GET`: Получение информации об авторе по его `id`.
        - `PUT`: Обновление данных автора.
        - `DELETE`: Удаление автора (если у автора есть книги, его удаление запрещено).

3. **Требования к коду**:
    - Валидация данных на уровне обработчиков (например, проверка корректности `published_year`).
    - Использование транзакций для операций удаления, чтобы проверить связи и при необходимости откатить изменения.
    - Обработка ошибок, возвращающая подробные сообщения клиенту (например, валидационные ошибки или ошибки БД).
    - Тесты для каждого эндпоинта API, с проверкой позитивных и негативных сценариев.

#### Советы по реализации:
- **База данных**: Создай схему SQLite с таблицами `books` и `authors` и настрой внешние ключи.
- **Организация кода**: Раздели код на слои (например, слой доступа к данным, слой бизнес-логики и слой обработчиков).
- **Модули и пакеты**: Создай отдельные пакеты для работы с БД и API. Это поможет в тестировании и улучшит читаемость.
- **Тесты**: Напиши тесты для эндпоинтов API, включая проверки на:
    - Успешное добавление, обновление и удаление данных.
    - Обработку ошибок, таких как попытка удаления автора с привязанными книгами.

Такое задание позволяет отработать навыки работы с Go и SQLite, а также расширить опыт создания RESTful API, включая тестирование и обработку ошибок. Если тебе нужно помочь с каким-то конкретным аспектом реализации, дай знать!