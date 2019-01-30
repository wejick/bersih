# Bersih : Golang Clean Architecture App Generator

```shell
$bersih create-app autocomplete
git repo url : github.com/wejick/sample
.. Creating application

$bersih create-service service-name-here
.. Creating service service-name-here

$bersih create-http-adapter adapter-name-here
.. Creating adapter adapter-name-here
```

App structure 

/cmd
 /autocomplete
  main.go
/pkg
 /autocomplete
  /model
  /repo
  /service
  /endpoint
   /http
   /grpc
 /adapter-name-here
  client.go


way to provide model from json
https://github.com/mholt/json-to-go

pass to go-fmt