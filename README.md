# 🔗 URL Shortener App

A simple fullstack URL shortener application that allows users to convert long URLs into shorter, shareable links.

---

## 🚀 Features

- Shorten long URLs into unique short links
- Redirect short URLs to the original destination
- Input validation (empty, format, scheme, host)
- Base62 encoding for short code generation
- RESTful API with clean architecture
- Simple and responsive frontend UI

---

## 🧱 Tech Stack

### Backend

- Go (Golang)
- Gin Framework
- PostgreSQL
- SQLX

### Frontend

- Next.js (React)
- Tailwind CSS

### Deployment

- Frontend: Vercel
- Backend: Local server exposed via ngrok
- Database: Supabase (PostgreSQL)

---

## 🏗️ Project Structure

```
url-shortener/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   └── model/
│   └── .env
│
├── frontend/
│   ├── src/
│   └── ...
```

---

## ⚙️ How It Works

1. User inputs a long URL from the frontend
2. Frontend sends request to backend API
3. Backend:
   - Validates URL
   - Stores original URL in database
   - Generates short code using Base62

4. Backend returns short URL
5. User accesses short URL → redirected to original URL

---

## 📡 API Endpoints

### Create Short URL

```
POST /v1/shorten
```

Request:

```json
{
  "original_url": "https://example.com"
}
```

Response:

```json
{
  "code": 201,
  "message": "URL shortened successfully",
  "data": {
    "short_url": "http://localhost:8000/abc123",
    "original_url": "https://example.com"
  }
}
```

---

### Redirect URL

```
GET /:short_code
```

---

## 🌐 Live Demo

Frontend:
👉 https://url-shortener-jade-ten.vercel.app/

⚠️ Note:
Backend is running locally and exposed using **ngrok**, so the API endpoint may change when the server restarts.

---

## 🧪 Local Development

### Backend

```bash
cd backend
go run cmd/main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

---

## 🔐 Environment Variables

### Backend (`.env`)

```
DATABASE_URL=your_supabase_connection_string
```

---

## ⚠️ Notes

- Backend is not deployed to a cloud server (uses ngrok for public access)
- ngrok URL changes on restart
- For production, backend can be deployed to services like Render, Fly.io, or VPS

---

## 📌 Future Improvements

- Custom alias for short URL
- URL analytics (click count)
- Expiration time for links
- Authentication system
- Deployment of backend to cloud server

---

## 👨‍💻 Author

Fajry Alvin
LinkedIn: _https://www.linkedin.com/in/fajryalvinhidayat_
GitHub: _https://github.com/fajryalvin12_

---
