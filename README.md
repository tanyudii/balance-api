# **Balance API**

A simple balance management API using **PostgreSQL**.

## Demo Project 🎬
* [Demo Video (G-Drive)](https://drive.google.com/file/d/1YvA4103FcINrooEnKjl6IJp2JBSpb3o6/view?usp=sharing)
* [Demo Video (Without Explanation) (Github)](https://github.com/tanyudii/balance-api/blob/main/docs/Balance%20API%20Recording.mp4)


## **🚀 Getting Started**

### **1. Clone the Repository**
```bash
git clone https://github.com/tanyudii/balance-api.git
cd balance-api
```

### **2. Set Up Environment Variables**
Copy the example environment file:
```bash
cp .env.example .env
```
Modify `.env` to configure database credentials and other settings.

### **3. Initialize the Database**
Ensure PostgreSQL is running, then execute:
```bash
psql -U <your-db-user> -d <your-db-name> -f db/migrations/init.sql
```
Replace `<your-db-user>` and `<your-db-name>` accordingly.

### **4. Start the Application**
Run the project using Docker:
```bash
docker compose up -d
```
- The `-d` flag runs containers in the background.
- Check logs if needed:
  ```bash
  docker compose logs -f
  ```

---

## **📌 API Endpoints**

### **1. Register an Account**
```http
POST /daftar
Content-Type: application/json

{
    "nama": "John Doe",
    "nik": "1234567890",
    "no_hp": "08123456789",
    "no_rekening": "9876543210"
}
```

### **2. Deposit Funds**
```http
POST /tabung
Content-Type: application/json

{
    "no_rekening": "9876543210",
    "nominal": 1000
}
```

### **3. Withdraw Funds**
```http
POST /tarik
Content-Type: application/json

{
    "no_rekening": "9876543210",
    "nominal": 1000
}
```

### **4. Check Balance**
```http
GET /saldo/{no_rekening}
```

