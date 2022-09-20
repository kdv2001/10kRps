# 10kRps

# Первый запуск 
Есть несколько вариантов запусков:
1) сбилдить сервер самому. Для этого необходимо задать несколько переменных сред 
    а) LISTEN_PORT - порт на котором поднимется сервер в формате :[port], если не задан, то сервр поднимется на 8010 порту 
    б) REDIS_CONTAINER - ip адрес и порт редис в формате [ip]:[port], если адрес не задан будет использован localhost:6379
2) запустить docker compose:
для этого вополните следующие команды:
```
sudo docker compose build 
sduo docker compose up
```
В этом случае редис прокинется на 6380 порт, а сервер на 8010 порт localhost
Для проверки перейдем в браузере по урлу 127.0.0.1:8010. Елси пришла ошибка 404(Not found) нам необходимо заполнить базу

# Заполнение базы 
Если вы использовали запуск через docker compose, вам необходимо заполнить redis данными иначе. Для этого запустите docker
compose выполнив (если он у вас не запщен):
```
docker sompose up
```
Подключитесь к redis, используя:
```
redis-cli -p 6380
```
Далее будем вносить последовательно записи, используя команду:
```
 zadd hackers [score] [name] 
```
