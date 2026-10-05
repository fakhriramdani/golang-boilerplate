# Golang Boilerplate

Boilerplate REST API Go dengan Echo, GORM (PostgreSQL), dan autentikasi JWT (access + refresh token).

## Tech Stack

- [Echo](https://echo.labstack.com/) — HTTP framework
- [GORM](https://gorm.io/) + PostgreSQL — ORM
- [go.uber.org/dig](https://github.com/uber-go/dig) — dependency injection
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) — JWT
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — password hashing

## Struktur Proyek

```
cmd/            entry point aplikasi
config/         load config dari environment (.env)
di/             gabungan registrasi dependency injection
internal/
  controller/   HTTP handler (per domain + impl)
  usecase/      business logic (per domain + impl)
  repository/   akses database (per domain + impl)
  entity/       model / struct database
  middleware/   JWT auth middleware + token manager
  router/       definisi routing
resource/
  database.go   koneksi database
  migrate.go    runner migrasi
  migrations/   file migrasi SQL (up/down)
```

## Setup

1. Buat database PostgreSQL.
2. Salin `.env.example` ke `.env` lalu sesuaikan nilainya:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=golang_boilerplate
   APP_PORT=8080
   JWT_SECRET=change-me
   ```

3. Jalankan:

   ```bash
   go run ./cmd
   ```

Migrasi dijalankan otomatis saat startup (menerapkan file `.up.sql` di `resource/migrations/` yang belum diterapkan).

## API Endpoint

Semua prefix `/api/v1`.

### Auth

| Method | Path                | Body                                                  | Deskripsi                     |
|--------|---------------------|-------------------------------------------------------|-------------------------------|
| POST   | `/auth/register`    | `{ "name", "email", "password" }`                     | Buat user baru                |
| POST   | `/auth/login`       | `{ "email", "password" }`                             | Login, dapat access + refresh |
| POST   | `/auth/refresh`     | `{ "refresh_token" }`                                 | Dapat pasangan token baru     |

`login` mengembalikan `access_token`, `refresh_token`, dan `user`. `refresh` mengembalikan pasangan token baru.

### Users (perlu autentikasi)

Header: `Authorization: Bearer <access_token>`

| Method | Path         | Deskripsi         |
|--------|--------------|-------------------|
| GET    | `/users`     | Daftar semua user |
| GET    | `/users/:id` | Detail user       |
| POST   | `/users`     | Buat user         |
| PUT    | `/users/:id` | Update user       |
| DELETE | `/users/:id` | Hapus user        |

## Catatan

- Refresh token bersifat stateless (tanpa tabel revoke). Tambahkan tabel refresh token jika butuh logout/revoke.
- Runner migrasi saat ini up-only (rollback belum di-wire).
