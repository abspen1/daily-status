# daily-status
OpenFaaS Go function to get a hash from redis.

```bash
faas template pull https://github.com/openfaas-incubator/golang-http-template
faas-cli up --build-arg GO111MODULE=on -f daily-status.yml'
```
