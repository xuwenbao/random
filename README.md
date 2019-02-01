# Simple

Prometheus监控测试应用

> Docker Image地址：registry.cn-hangzhou.aliyuncs.com/xuwenbao/simple:latest

## 测试URL

* “http://<docker ip>:8080/”：测试请求地址
* “http://<docker ip>:8080/metrics”：Metrics获取地址

## Prometheus Query

按版本监控每分钟请求数（过滤Prometheus的抓取Reqeust）。

```sql
sum(negroni_requests_total{app="simple", path != "/metrics"} - negroni_requests_total{app="simple", path != "/metrics"} offset 1m) by (version)
```
