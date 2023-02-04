# qnd.be
Just another simple url shortener to use with curl. Kinda like [0x0.st](https://0x0.st). Just way more shitty. It was written in one night, so don't expext anything great.

## usage
there is built in authentication, but that can be turned off for adding links

without auth:
```
curl -X POST \
-d https://example.com \
https://qnd.be
```

with auth:
```
curl -X POST \
-H "Authorization: Bearer <token>" \
-d https://example.com \
https://qnd.be
```

Deleting only works with auth:
```
curl -X DELETE \
-H "Authorization: Bearer <token>" \
-d <short> \
https://qnd.be
```

That's all it can do. Very simple GUIs can also by found at ``/add/`` and ``/delete/`` (second slash is important).

## host
Calling this stable would be wrong. It was written in one night, it's mostly untested and lacks features. If you still wanna use it:

- clone [github.com/rrune/goshort](https://github.com/rrune/goshort)
- create the SQL table using `db.sql`
- copy ``/data/config.yml.sample`` to ``/data/config.yml``
- configure the config
- change port in ``docker-compose.yml``
- ``docker-compose up``

## stuff used
This is written in Go. That's basically it. Uses mariadb/mysql to store stuff. If you want to use some other database, you have to write that code yourself (has to implement the Database interface).

## abuse/contact
Please only link stuff compliant with german law. Thanks! If you want something removed for whatever reason, send me a mail. Responses might take a day.

[rune@ruune.de](mailto:rune@ruune.de)