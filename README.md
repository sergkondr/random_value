random_value
=====

Example plugin for [telegraf](https://github.com/influxdata/telegraf) that returns random value between min and max value.

### Telegraf config example
```toml
[[inputs.random_value]]
  min = 0
  max = 100
```

### Result in grafana
![screen](https://github.com/sergkondr/stuff/blob/master/telegraf_random_value.png)
