FROM grafana/k6

COPY ./script.js /tests/script.js

ENTRYPOINT [ "k6", "run", "/tests/script.js" ]
