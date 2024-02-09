# GroupAssist
### Описание
Бекенд переписанный на языке Go когда-то реализованного мною RESTFul-приложения [GroupAssistant](https://github.com/psevdocoder/sipi_backend).

Приложение для автоматизации создания очередей для сдачи работ, фиксирования посещаемости студентов, создания опросов внутри группы.

### Особенности
- Структура проекта следует чистой архитектуре.
- Используется SQLX + Gin + [InMemoryCache](https://github.com/psevdocoder/InMemoryCacheTTL).
- InMemoryCache имплементирован как Gin-Middleware для кэширования GET запросов.

InMemoryCache - собственная реализация кэша в памяти для хранения записей в паре ключ-значение с TTL. После проведения нагрузочного тестирования выяснилось, что кэш позволяет повысить RPS с 200 до 1000-1100 при GET запросах на один и тот же эндпоинт. 

### TODO
- [ ] Добавить пользователей, JWT-токены, разделение пользователей на роли админ, модератор, пользователь.
- [ ] Rate limiting
- [ ] Опросы
- [ ] Посещаемость, журнал
