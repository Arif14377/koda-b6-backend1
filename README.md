# Backend 1
minitask backend 1, membuat rest api dengan data product dan user hardcode di requests.

## Task
- CRUD email & password

## Update Task 1
1. membuat endpoint u/ auth dan register
2. validasi data sederhana
3. kembalikan status error jika tidak sesuai

## Update Task 2
- Hashing password ketika register dan update password
    - hashing pakai library matthewhartstonge/argon2

## Update Task 3
- Implementasi swagger

Cara menjalankan swagger (current dir: /cmd$):
1. Genereate comment swaggo
```bash
swag init -g main.go -d .,../internal -o docs --parseInternal
```
2. Jalankan project
```bash
go run main.go
```


## TODO
- Swagger masih belum tergroup
- request kirim body json belum berhasil
- Password belum dihash