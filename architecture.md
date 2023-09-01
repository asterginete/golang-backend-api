+---------------------+     +---------------------+     +---------------------+
|                     |     |                     |     |                     |
|    API Gateway      +----->    Application     +----->    Database         |
| (Rate Limiting,     |     | (Go-Gin, Handlers,  |     | (Data Storage,      |
|  Logging, CORS)     |     |  Middleware, JWT)   |     |  Migrations)        |
|                     |     |                     |     |                     |
+----------+----------+     +----------^----------+     +----------^----------+
           |                        |                       |
           |                        |                       |
           |                        |                       |
+----------v----------+     +--------+-----------+     +------v--------------+
|                     |     |                    |     |                     |
|    File Storage     |     |   Cache Server     |     |   Email Server      |
| (Book Covers, etc.) |     | (Redis, go-cache)  |     | (Password Resets)   |
|                     |     |                    |     |                     |
+---------------------+     +--------------------+     +---------------------+
