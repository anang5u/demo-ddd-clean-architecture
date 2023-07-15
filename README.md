## Instalasi
> git clone git@github.com:anang5u/demo-ddd-clean-architecture.git
> cd demo-ddd-clean-architecture
> docker-compose up

atau
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

## Arsitektur Aplikasi
![arsitektur aplikasi](assets/diagram-architecture.jpg)