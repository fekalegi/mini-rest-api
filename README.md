# mini-rest-api

untuk menjalankan aplikasi ini, di perlukan :

GO Versi : 1.17 dan PostgreSQL

Buatlah file local.env terlebih dahulu di folder utama
![image](https://user-images.githubusercontent.com/57470112/161087169-cd72949c-5dea-4644-9a1c-2faf5a6a6502.png)


file ini berisikan settingan untuk database postgre.

```
DB_HOST = localhost
DB_USERNAME = postgres
DB_PASSWORD = admin123
DB_DATABASE = mini_rest_api
DB_PORT = 5432
DB_DEBUG_MODE = false
```

Setelah itu build dan jalankan project
![image](https://user-images.githubusercontent.com/57470112/161089002-078822ff-99b6-443a-8ec0-4032f57262d5.png)

dan jika ingin membuka api documentation bisa ke link swaggernya , untuk di localhost :
localhost:5000/swagger/index.html
![image](https://user-images.githubusercontent.com/57470112/161089332-f3739fc4-4371-47d0-b8c5-2a73ef072415.png)
