**Submit Tax**
----
Submit a tax object to the database.
  
  * **URL**
  
    `/v1/taxes`
  
  * **Method:**
  
    `POST`
  
  * **Success Response:**
  
    * **Code:** 200 <br />
      **Content:** 
      ```
      {
          "success": true,
          "status_code": 200,
          "message": ""
      }
      ```
   
  * **Error Response:**
  
      * **Code:** 400 Bad Request <br />
      **Content:**
      ```
      {
          "success": false,
          "status_code": 400,
          "message": "Tax code is invalid"
      }
      ```
    
      OR    
  
      * **Code:** 405 Method Not Allowed <br />
      **Content:**
      ```
      {
          "success": false,
          "status_code": 405,
          "message": "Invalid method"
      }
      ```
  
      OR
      
      * **Code:** 500 Internal Server Error <br />
          **Content:**
      ```
      {
          "success": false,
          "status_code": 500,
          "message": "Failed to insert bills. Err:<some error messages>"
      }
      ```
  
  * **Sample Call:**
  
    ```
    curl -X POST \
      http://localhost:8080/v1/taxes \
      -H 'cache-control: no-cache' \
      -F name=Games \
      -F code=3 \
      -F price=50
    ```

----

**Get Taxes**
----
Get bills from the database and calculate the taxes.

* **URL**

  `/v1/taxes`

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    {
        "success": true,
        "message": "",
        "status_code": 200,
        "data": [
            {
                "id": 1,
                "name": "Big Mac",
                "code": 1,
                "price": 1000,
                "is_refundable": true,
                "tax": 100,
                "amount": 1100
            },
            {
                "id": 2,
                "name": "Lucky Stretch",
                "code": 2,
                "price": 1000,
                "is_refundable": false,
                "tax": 30,
                "amount": 1030
            },
            {
                "id": 3,
                "name": "Movie",
                "code": 3,
                "price": 150,
                "is_refundable": false,
                "tax": 0.5,
                "amount": 150.5
            },
            {
                "id": 4,
                "name": "Games",
                "code": 3,
                "price": 50,
                "is_refundable": false,
                "tax": 0,
                "amount": 50
            }
        ]
    }
    ```
 
* **Error Response:**

    * **Code:** 405 Method Not Allowed <br />
    **Content:**
    ```
    {
        "success": false,
        "status_code": 405,
        "message": "Invalid method"
    }
    ```

    OR
    
    * **Code:** 500 Internal Server Error <br />
        **Content:**
    ```
    {
        "success": false,
        "status_code": 500,
        "message": "Failed to get bills. Err:<some error messages>"
    }
    ```

* **Sample Call:**

  ```
  curl -X GET \
    http://localhost:8080/v1/taxes \
    -H 'cache-control: no-cache'
  ```