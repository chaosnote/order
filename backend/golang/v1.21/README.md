# Order Backend

``` bash
### 啟動
sudo docker-compose up -d 
### 測試
sudo docker-compose up -d mariadb redis
### 停止
sudo docker-compose down
### 編譯
sudo docker run -ti --rm -w /app -v /home/chris/order_golang/work:/app -p 8080:8080 --name some-golang golang:1.24-bullseye bash
```

``` bash
## 測試
go run -race .
## 編譯
go build -o ./dist/order .
```
