## Instalasi
> docker-compose up

atau secara **background**
> docker-compose up -d

## Fitur Aplikasi

|       | | |
| ----------- | ----------- |---|
| Info Tagihan | GET| http://127.0.0.1:4321/v1/demo/info |
| Cek Tagihan| GET| http://127.0.0.1:4321/v1/demo/inquiry/{nomor_kontrak} |
| Bayar Tagihan| POST| http://127.0.0.1:4321/v1/demo/payment |
| Token| POST| http://127.0.0.1:4321/v1/demo/token |
| Detail Pembayaran| GET| http://127.0.0.1:4321/v1/demo/payment/{id_transaksi} |
|       | | |

## Asset Pendukung

> [Postman Collection](assets/DDD-Demo.postman_collection.json)
> [File SQL](assets/db_demo.sql)

## Arsitekture Aplikasi
```
![arsitekture aplikasi](assets/diagram-architecture.jpg)
```