# Jobhun-Backend-Technical-Test

## Overview

Merupakan sebuah RESTful Web Service menggunakan bahasa Golang yang terdiri dari 5 endpoint utama, yaitu:

- <b>/get-all-mahasiswa</b> digunakan untuk mendapatkan seluruh data mahasiswa yang terdaftar di dalam database. Menggunakan metode <b>GET</b>.
- <b>/get-mahasiswa-on-id/{id}</b> digunakan untuk mendapatkan data mahasiswa sesuai dengan id yang diberikan. Akan mengembalikan error apabila id yang dimasukkan salah. Menggunakan metode <b>GET</b>.
- <b>/insert-mahasiswa</b> digunakan untuk memasukkan data mahasiswa baru ke dalam database. Menggunakan metode <b>POST</b>.
- <b>/update-mahasiswa</b> digunakan untuk melakukan perubahan terhadap data mahasiswa tertentu. Akan mengembalikan error apabila id yang dimasukkan salah. Menggunakan metode <b>PUT</b>.
- <b>/delete-mahasiswa/{id}</b> digunakan untuk menghapus data mahasiswa sesuai dengan id yang diberikan. Akan mengembalikan error apabila id yang dimasukkan salah. Menggunakan metode <b>DELETE</b>.

Seluruh endpoint ini menggunakan port 8080.

## Cara menjalankan
1. Clone repository ini ke dalam komputer Anda.
2. Buatlah file .env baru dengan isi sesuai .env.example, sesuaikan dengan database yang Anda miliki.
3. Jalankan perintah berikut
```
go run main.go
```
4. Program sudah berhasil dijalankan.

## Contoh struktur JSON
- Untuk endpoint insert mahasiswa, struktur JSON adalah sebagai berikut

```
{
    "nama": "Nama mahasiswa",
    "usia": "Usia mahasiswa",
    "gender": "Gender mahasiswa",
    "jurusan": {
        "nama": "Nama Jurusan"
    },
    "hobi": {
        "nama": "Nama hobi"
    }
}
```

- Untuk endpoint update mahasiswa, struktur JSON adalah sebagai berikut

```
{
    "id_mahasiswa":"ID mahasiswa",
    "nama": "Nama mahasiswa",
    "usia": "Usia mahasiswa",
    "gender": "Gender mahasiswa",
    "jurusan": {
        "nama": "Nama Jurusan"
    },
    "hobi": {
        "nama": "Nama hobi"
    }
}
```
