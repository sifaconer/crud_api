**CRUD API Go**

CRUD API in go gRPC server

**Run project:**

1. Clone the project.  

2. Create a .env file.  
> Example .env file  
```python
# SERVER
SERVER_PORT=8998
# POSTGRES
POSTGRES_HOST=db-pg
POSTGRES_PORT=5432
POSTGRES_USER=crud
POSTGRES_PASSWORD=eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81
POSTGRES_DB=crud_database
POSTGRES_SSL_MODE=disable
# REDIS
REDIS_HOST=db-redis
REDIS_PORT=6379
REDIS_PASSWORD=eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81
```

3. Inside the project build with docker-compose 

> Command: docker-compose  
```bash
docker-compose build 
```

4. Run
> Command: docker-compose  
```bash
docker-compose up -d
```

