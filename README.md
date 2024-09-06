# MyPass (Password Manager Backend)

MyPass adalah aplikasi backend yang dibangun menggunakan Go (Golang) dengan pendekatan Clean Architecture. Aplikasi ini dirancang untuk mengelola data pengguna, mendukung penyimpanan dan pengambilan informasi pengguna dengan aman dan efisien. Backend ini siap untuk diintegrasikan dengan frontend Password Manager yang bertujuan untuk membantu pengguna mengelola kata sandi mereka secara aman.

Fitur
Manajemen Pengguna:

Mendapatkan data pengguna berdasarkan ID melalui endpoint RESTful.
Menyimpan data pengguna baru ke dalam database PostgreSQL.
Endpoint RESTful:

Endpoint GET /user/{id} untuk mengambil data pengguna berdasarkan ID.
Keamanan:

Praktik keamanan terbaik diterapkan untuk memastikan data pengguna dikelola dengan aman.
Teknologi yang Digunakan
Go (Golang): Bahasa pemrograman utama untuk backend, dikenal karena kinerja dan efisiensinya.
GORM: ORM untuk berinteraksi dengan database PostgreSQL, mempermudah operasi CRUD.
PostgreSQL: Database relasional yang digunakan untuk menyimpan data pengguna.
Gorilla Mux: Router HTTP yang mendukung parameter path dinamis dan middleware.
Clean Architecture: Struktur yang memisahkan logika bisnis, data, dan presentasi, memudahkan pemeliharaan dan pengembangan