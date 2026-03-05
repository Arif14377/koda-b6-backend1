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
https://github.com/swaggo/swag

### Cara menjalankan swagger (current dir: /cmd$):
1. Genereate comment swaggo
```bash
swag init -g main.go -d .,../internal -o docs --parseInternal
```
2. Jalankan project
```bash
go run main.go
```

### Masalah swag terminal (suka ngilang)
1. Pastikan binary sudah terinstall
```bash
ls $(go env GOPATH)/bin | grep swag
```
2. Cek gopath & gobin
```bash
go env GOPATH
go env GOBIN
```
3. Tambahkan ke path
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
4. Reload
```bash
source ~/.bashrc
```
5. Cek swag
```bash
swag --version
```

## TODO
- Swagger masih belum tergroup
- request kirim body json belum berhasil
- Password belum dihash