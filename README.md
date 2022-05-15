# nipa-interview

# วิธีการ clone .

  - เปิด terminal
  - ทำการโคลน
  ```
  git clone git@github.com:pond18143/nipa-interview.git
  ```
# สร้าง env สำหรับ connect DB
  - สร้างไฟล์ config.yaml
  ```
  cd nipa-interview&&mkdir configure
  ```
  ```
  touch config.yaml
  ```
  - ใส่ค่า ใน config.yaml
  ```
  mssql :
    databaseType    : "mssql"
    server          : "<server ip>"
    port            : <port>
    user            : "<user>"
    password        : "<password>"
    database        : "<database name>"
  ```
# วิธี run
  ```
  go run .
  ```
  - LINK : [swagger](http://localhost:8080/swagger/index.html)
