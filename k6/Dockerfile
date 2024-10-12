FROM grafana/k6

COPY ./tests/script.js /tests/script.js

ENTRYPOINT [ "k6", "run", "/tests/script.js" ]