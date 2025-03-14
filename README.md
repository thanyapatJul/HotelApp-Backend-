# Hotel Booking API

## 📌 Overview
This is a **hotel booking system** built with **Golang (Gin)** for the backend and **React (Chakra UI, FullCalendar)** for the frontend. Users can register, book hotels, and leave reviews.

---

## ⚡ Features
### ✅ **User Management**
- Register/Login (JWT Authentication)
- Session Management

### ✅ **Hotel Management**
- View available hotels
- View hotel images & ratings
- Display available rooms with a calendar

### ✅ **Booking System**
- Select date range & book a hotel
- Check room availability before booking
- View user-specific bookings

### ✅ **Review System**
- Leave reviews & ratings for hotels
- Fetch & display reviews per hotel

---

## 🛠️ **Tech Stack**
### **Backend (Golang + Gin)**
- Gin Gonic
- GORM (MySQL)
- JWT Authentication

## 🔒 **Security**
- JWT Authentication
- CORS Authentication



### **Database (MySQL)**
- User Management
- Hotel Listings
- Booking Records
- Reviews & Ratings

---



## 📡 **API Endpoints**
### 🧑‍💻 **Authentication**
| Method | Endpoint        | Description            |
|--------|----------------|------------------------|
| POST   | `/register`     | Register a new user   |
| POST   | `/login`        | Authenticate user     |
| GET    | `/checksession` | Validate session      |

### 🏨 **Hotels**
| Method | Endpoint          | Description               |
|--------|------------------|---------------------------|
| GET    | `/hotel/list`     | Get all hotels           |
| GET    | `/hotel/:id`      | Get hotel by ID          |
| POST   | `/hotel/book`     | Book a hotel room        |
| GET    | `/hotel/bookings/:id` | Get hotel bookings |

### ⭐ **Reviews**
| Method | Endpoint            | Description            |
|--------|--------------------|------------------------|
| GET    | `/hotel/review/:id` | Get reviews by hotel  |
| POST   | `/hotel/review`     | Add a new review      |






