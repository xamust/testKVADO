Согласно Т.З., было необходимо cпроектировать базу данных, в которой содержится авторы
книг и сами книги. 
Необходимо написать сервис который будет по
автору искать книги, а по книге искать её авторов.

Требования к сервису:
- Сервис должен принимать запрос по GRPC;
- Должна быть использована база данных MySQL;
- Код сервиса должен быть хорошо откомментирован;
- Код должен быть покрыт unit тестами;
- В сервисе должен лежать Dockerfile, для запуска базы данных с
тестовыми данными;
- Должна быть написана документация, как запустить сервис.
Плюсом будет если в документации будут указания на команды,
для запуска сервиса и его окружения, через Makefile;
- код должен быть выложен на github.


При запуске docker-compose разворачиваются 3 контейнера:
1. сам сервис (search),
2. "боевая" база (KVADOTest),
3. база "для тестирования" (KVADOTestForTest) (в принципе ее наличие в текущем задании не требуется, но решил ее реализовать, для наглядности тестов).

БД MySQL (боевая и тестовая) имеет две таблицы TestBooksTable и TestWritersTable, заполненные тестовыми данными (бд и наполнение создаются скриптом при развертывании).

В каталоге сервиса реализован Makefile c функционалом: init, build, run, compile, test.

Код в достаточной мере откомментирован. Реализованы unit тесты для тестировани запросов к бд и функционала gRPC, в рамках данного сервиса.

Для использования gRPC  реализована структура GRPCServer и функция Search, из которой производятся запросы к БД (FindByWriter и FindByBook, реализованные в пакете store).

Также реализован клиент (client) gRPC, для удобства и наглядности работы сервиса (запускается независимо от контейнеров).

В сервисе, в качестве логгера, используется logrus.

Конфигурация глобальных параметров сервиса осуществляется через конфигурационный файл server.toml. 
