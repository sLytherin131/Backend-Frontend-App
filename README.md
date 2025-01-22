Admin Reservasi Lapangan Futsal (Backend Frontend API)
Studi Kasus: Olahraga futsal adalah kegiatan yang banyak digemari banyak orang, aplikasi seperti reservasi lapangan futsal yang mampu membantu admin dalam mengelola data lapangan, pelanggan, dan reservasi secara efisien. Membangun sistem reservasi lapangan futsal bisa membantu admin untuk mengelola data lapangan, pelanggan, dan reservasi secara digital. Sistem ini dirancang agar mudah digunakan, terkoneksi dengan database, dan mampu menangani operasi CRUD (Create, Read, Update, Delete).

Program ini memiliki database, backend, dan frontend. Backend menggunakan bahasa GO sebagai bahasa pemrogramannya, dengan menggunakan gofiber sebagai framework backendnya. Backend terkoneksi dengan database reservasi_futsal.sql. Untuk frontendnya menggunakan html, javascript, css, dan bootstrap sebagai framework. Koneksi antara backend dan frontend sudah menggunakan API HTTP. Program ini memiliki fungsi CREATE, READ, UPDATE, dan DELETE.

Backend:
![image](https://github.com/user-attachments/assets/50e1cbff-3acf-4cf3-b662-2211030598af)
Untuk Backend menggunakan bahasa GO sebagai bahasa pemrogramannya, dengan menggunakan gofiber sebagai framework backendnya. Backend terkoneksi dengan database reservasi_futsal.sql.

Frontend:
Lapangan: ![image](https://github.com/user-attachments/assets/219fa86e-db91-4dbd-a41c-40380cb6cad2)
Ada kolom isian untuk Nama lapangan, Tipe lapangan, dan Harge per jamnya. Ada fitur Tabel untuk Read data, Delete untuk hapus data, Edit untuk update data, Tambah lapangan untuk Create data.

Pelanggan: ![image](https://github.com/user-attachments/assets/170406dc-12ad-45fa-ad77-0fe5aaee4048)
Ada kolom isian untuk Nama pelanggan dan Nomor telepon. Ada fitur Tabel untuk Read data, Delete untuk hapus data, Edit untuk update data, Tambah pelanggan untuk Create data.

Reservasi: ![image](https://github.com/user-attachments/assets/9b472e93-8fe7-4332-a595-03f24fe41cc2)
Ada dropdown Lapangan untuk memilih lapangan yang sudah di input di database, dropdown Pelanggan untuk memilih pelanggan yang sudah di input ke database, Waktu Mulai dengan menggunakan fitur kalender untuk memilih tanggal dan jam mulai, dan kolom isian Durasi(jam) untuk berapa jam lapangan akan disewa. Ada fitur Tabel untuk Read data, Delete untuk hapus data, Tambah reservasi untuk Create data.

Edit Data: ![image](https://github.com/user-attachments/assets/bb1482bf-dca2-494f-8616-fe4d0d7d011e)

Database:
![image](https://github.com/user-attachments/assets/9e13a119-89d2-4324-86c7-b253ee3bf7b5)
![image](https://github.com/user-attachments/assets/9730f851-8ab5-4cd1-960c-9ca47ce25c41)
![image](https://github.com/user-attachments/assets/22beeff8-a9c9-459e-8e27-60f90530ad2e)
![image](https://github.com/user-attachments/assets/a3081166-7425-426b-989b-f2d4a681ce1e)


Cara install: 
Download reservasi_futsal.sql, main.go, frontend.html, go.mod, go.sum. 
Import database reservasi_futsal.sql ke database seperti phpmyadmin.
Jalankan server lewat Visual Studio Code dengan menjalankan file main.go di terminal dengan command "go run main.go".
Setelah server berjalan, buka frontend.html di browser
Selesai, program akan berjalan dengan benar.


