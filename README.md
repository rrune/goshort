# qnd.be
Just another simple url shortener to use with curl. Kinda like [0x0.st](https://0x0.st). Just way more shitty. It was written in one day, so don't expext anything great.

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

That's all it can do. Here, auth is probably on all the way atm.

## host
If that repo I linked is still private, well, tough luck. Contact me if you really want it, but maybe just use something you can actually trust to not explode because of bad error handling.

- clone [github.com/rrune/goshort](https://github.com/rrune/goshort)
- copy ``/data/config.yml.sample`` to ``/data/config.yml``
- configure the config
- change port in ``docker-compose.yml``
- ``docker-compose up``

SQL table should be called ``shortLinks``, 3 colums:

- short (type text)
- url (type text)
- timestamp (type timestamp)

## stuff used
This runs on Go. That's basically it. Uses mariadb to store stuff. If you want to use some other database, you have to write that code yourself.

## abuse/contact
Should this actually be without auth, either because I decided to make it public or because I fucked up, please only link stuff compliant with german law. Thanks!

[rune@ruune.de](mailto:rune@ruune.de)