FROM migrate/migrate

COPY ./ /migrations

CMD ["-path", "/migrations", "-database",  "mysql://root:password@tcp(mysql:3306)/amalgam-db", "up"]
