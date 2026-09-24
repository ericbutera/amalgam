# K6 Tests

These k6 tests exercise the public GraphQL API. They provide a small smoke
test for the feed and article read paths; the larger GraphQL load test lives in
[`k6/load-test-graph`](../load-test-graph).

Read the [Grafana K6 docs](https://grafana.com/docs/k6/latest/) for more information.

## Quick start

```sh
mise run k6:test
```

## TODO

- [fake data generation](https://github.com/grafana/k6-example-data-generation/blob/main/src/index.js)
- [load test](https://k6.io/docs/testing-guides/running-large-tests/)

## Expected Output

The tilt environment has k6 wired up to run tests. The output should look like this:

```txt
         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/

     execution: local
        script: /tests/script.js
        output: -

     scenarios: (100.00%) 1 scenario, 1 max VUs, 1 iteration
              * default: 1 iteration for 1 VU


     █ GraphQL feed and article queries

       ✓ OK

       ✓ OK

     checks.........................: 100.00% 8 out of 8
     data_received..................: 2.0 kB  447 kB/s
     data_sent......................: 331 B   75 kB/s
     group_duration.................: avg=1.05ms   min=763.91µs med=1.01ms   max=1.44ms   p(90)=1.34ms   p(95)=1.39ms
     http_req_blocked...............: avg=156.56µs min=1.7µs    med=2.06µs   max=620.41µs p(90)=435µs    p(95)=527.71µs
     http_req_connecting............: avg=28.5µs   min=0s       med=0s       max=114µs    p(90)=79.8µs   p(95)=96.9µs
     http_req_duration..............: avg=766.05µs min=660.16µs med=775.25µs max=853.54µs p(90)=849.6µs  p(95)=851.57µs
       { expected_response:true }...: avg=766.05µs min=660.16µs med=775.25µs max=853.54µs p(90)=849.6µs  p(95)=851.57µs
     http_req_failed................: 0.00%   0 out of 4
     http_req_receiving.............: avg=109.72µs min=33.45µs  med=46.5µs   max=312.45µs p(90)=235.64µs p(95)=274.05µs
     http_req_sending...............: avg=19.54µs  min=4.33µs   med=9.12µs   max=55.58µs  p(90)=42.6µs   p(95)=49.09µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=636.78µs min=292.12µs med=731.75µs max=791.5µs  p(90)=791.41µs p(95)=791.45µs
     http_reqs......................: 4       910.461775/s
     iteration_duration.............: avg=4.31ms   min=4.31ms   med=4.31ms   max=4.31ms   p(90)=4.31ms   p(95)=4.31ms
     iterations.....................: 1       227.615444/s


running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [ 100% ] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU
```
