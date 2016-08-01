Build the image
```
docker build -t "revel_basic" .
```

Run MongoDB - 
```
docker run -d -p 27017:27017 -p 28017:28017 --name mongoContainer mongo:latest
```


Run revel -
```
docker run -d -p 9000:9000 --link mongoContainer:mongo revel_basic
```


While testing out the devo-
```
revel run github.com/sadhanandh/revelbasic/application dev 9000
```
