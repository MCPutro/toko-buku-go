# toko-buku using golang REST API and Template 

### Web UI
http://localhost:8080/  

### Rest API 
1. Sign Up
   - Method : POST
   - Endpoint : http://localhost:8080/user/SignUp
   - Request :
    ```json
    {
      "Email" : "email@test.com",
      "UserName": "test-new",
      "Password": "123456789",
      "UserType": 1
    }
    ```
   - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": {
          "ID": 4,
          "Email": "email@test.com",
          "UserName": "user-name",
          "UserType": 1 
      } 
   }
   ```
2. Sign In
    - Method : POST
    - Endpoint : http://localhost:8080/user/SignUp
    - Request :
    ```json
    {
      "Email" : "email@test.com",
      "Password": "123456789"
    }
    ```
    - Response :
    ```json
    {
        "Status": "success",
        "Message": "",
        "Data": {
            "ID": 4,
            "UserName": "user-name", 
            "Email": "email@test.com",
            "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6ImVtYWlsQHRlc3QuY29tIiwiUm9sZSI6MSwiaXNzIjoiTUNQdXRybyIsImV4cCI6MTY0NzI1OTQ3NiwiaWF0IjoxNjQ3MTczMDc2fQ.oFoiLaipcVluTJD3fFbC7PmGyRS9ST3D7JQOo75QZUA",
            "UserType": 1
        } 
   }
   ```
3. Get Books
    - Method : GET
    - Endpoint : http://localhost:8080/book/{bookId}
    - Param : bookId - string
    - Response :
    ```json
    {
        "Status": "success",
        "Message": "",
        "Data": {
            "Id": "5509625f-84fc-4ae5-a61f-ed5181f25101",
            "Title": "buku makan - minum 3",
            "Author": "yuk gas keun",
            "Stock": 40,
            "Price": 10500,
            "Discount": 30
        } 
   }
   ```
4. Add Book
   - Method : POST
   - Endpoint : http://localhost:8080/book/Add
   - Request :
   ```json
   {
      "Title": "buku makan - minum 3",
      "Author": "yuk gas keun", 
      "Stock": 40,"Price": 10500, 
      "Discount": 30
   }
   ```
   - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": {
          "ID": 4,
          "UserName": "user-name", 
          "Email": "email@test.com",
          "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6ImVtYWlsQHRlc3QuY29tIiwiUm9sZSI6MSwiaXNzIjoiTUNQdXRybyIsImV4cCI6MTY0NzI1OTQ3NiwiaWF0IjoxNjQ3MTczMDc2fQ.oFoiLaipcVluTJD3fFbC7PmGyRS9ST3D7JQOo75QZUA",
          "UserType": 1
      } 
   }
   ```
5. Get All Book
   - Method : GET
   - Endpoint : http://localhost:8080/book/All
   - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": [
        {
          "Id": "c7402a74-88a0-45b9-8bae-1360557036ed",
          "Title": "22222",
          "Author": "222222",
          "Stock": 2,
          "Price": 2222,
          "Discount": 100
        },
        {
          "Id": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
          "Title": "buku makan - minum 3",
          "Author": "yuk gas keun",
          "Stock": 22,
          "Price": 10500,
          "Discount": 30
        },
        {
          "Id": "486ed146-3e51-4c38-be1a-7b7e6b485183",
          "Title": "buku makan - minum 3",
          "Author": "yuk gas keun",
          "Stock": 40,
          "Price": 10500,"Discount": 30
        }
      ]
   }
   ```
6. Delete Book
   - Method : GET
   - Endpoint : http://localhost:8080/book/Delete/{BookId}
   - param : BookId - string
   - Response :
   ```json
   {"Status": "success","Message": "","Data": null}
   ```
7. Update Book
   - Method : POST
   - Endpoint : http://localhost:8080/book/Update/{BookId}
   - Request : 
   ```json
   {
      "Title": "111111111111",
      "Author": "string",
      "Stock": 9,
      "Price": 32,
      "Discount": 0
   }
   ```
   - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": {
          "Id": "5509625f-84fc-4ae5-a61f-ed5181f25101",
          "Title": "111111111111",
          "Author": "string",
          "Stock": 9,
          "Price": 32,
          "Discount": 0
      }
   }
   ```
8. Buy a Book
    - Method : POST
    - Endpoint : http://localhost:8080/book/Update/{BookId}
    - Request :
   ```json
   {
      "Customer": "email@test.com",
      "BookID": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
      "Quantity": 1
   }
   ```
    - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": {
          "Id": "0165a270-8b8d-4f53-a3ed-c593687f6478",
          "Date": "2022-03-15T19:48:39.4613267+07:00",
          "Customer": "email@test.com",
          "BookID": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
          "BookTitle": "",
          "Price": 10500,
          "Quantity": 1,
          "Discount": 30,
          "Total": 7350
      }
   }
   ```
9. Get Transaction List
    - Method : GET
    - Endpoint : localhost:8080/transaction/history/{email}
    - param : email(Customer Email) - string 
    - Response :
   ```json
   {
      "Status": "success",
      "Message": "",
      "Data": [
        {
          "Id": "8776b817-8e5f-4288-b52d-31524cc29b23",
          "Date": "2022-03-15T19:45:48.186363+07:00",
          "Customer": "email@test.com",
          "BookID": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
          "BookTitle": "buku makan - minum 3",
          "Price": 10500,
          "Quantity": 1,
          "Discount": 30,
          "Total": 7350
        },
        {
          "Id": "ad5145af-f602-491f-b298-0e8e5afe3bc6",
          "Date": "2022-03-15T19:45:48.596034+07:00",
          "Customer": "email@test.com",
          "BookID": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
          "BookTitle": "buku makan - minum 3",
          "Price": 10500,
          "Quantity": 1,
          "Discount": 30,
          "Total": 7350
        },
        {
          "Id": "0165a270-8b8d-4f53-a3ed-c593687f6478",
          "Date": "2022-03-15T19:48:39.461326+07:00",
          "Customer": "email@test.com",
          "BookID": "720c51ac-2b86-4f4c-bed4-8ef6708777c9",
          "BookTitle": "buku makan - minum 3",
          "Price": 10500,
          "Quantity": 1,
          "Discount": 30,
          "Total": 7350
        }
      ]
   }
   ```

