# DockerCompile

```sh
docker build -t go-app .
docker run -p 8080:8080 go-app --name servergo

curl http://172.17.0.2:8080/color

curl -X POST http://172.17.0.2:8080/sumar -H "Content-Type: application/json" -d '{"a": 10, "b": 5}'
curl -X POST http://172.17.0.2:8080/restar -H "Content-Type: application/json" -d '{"a": 10, "b": 3}'
curl http://172.17.0.2:8080/props
```


