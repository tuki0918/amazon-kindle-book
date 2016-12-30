# amazon-kindle-book

### Usage

```
$ go run book.go
 ⇛ http server started on :8080
```

### Request

```
$ curl -s http://localhost:8080/\?item\=https://www.amazon.co.jp/gp/product/B008YOHMCS | jq  
{
  "name": "Think Simple　―アップルを生みだす熱狂的哲学",
  "image": "https://images-fe.ssl-images-amazon.com/images/I/41sZ0NeG5jL._SY346_.jpg",
  "link": "https://www.amazon.co.jp/gp/product/B008YOHMCS",
  "free": true,
  "description": "",
  "comment": ""
}
```
