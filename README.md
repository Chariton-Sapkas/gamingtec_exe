<h3>  Prerequisites</h3>

1. docker
2. PostMan (for testing APIs)

<h3> Running the application </h3>

#### For the following steps I used a machine running Windows 11 and it worked

1. Open your terminal  
2. git clone <repository-url> --> it's attached on the email
3. cd <repository-directory> --> use the name of the repo just cloned
4. docker build -t user-microservice .
5. docker run -p 5050:5050 user-microservice
6. Make sure you see on your terminal the message "2024/08/11 21:08:15 gRPC server listening on :5050"
7. Navigate to PostMan
8. Click New(Located on the left panel)
9. Select the gRPC option
10. For the url(text box with the lock emoji placeholder) type "localhost:5050"
11. Now click enter and make sure the lock emoji is crossed with a line (symbolizing the http)
12. In drop-down menu(next to Cancel button) you should be able to select different methods to Invoke
13. If this does not work try to repeat the process from step 4 by making sure you execute the steps correctly

#### If you are struggling to test the APIs in PostMan you should try the following:
1. After executing step 11, Navigate to tab "Service Definition"
2. Click import a .proto file
3. choose the proto file under api>proto>user.proto
4. The methods should appear on the drop-down menu
5. With process, you won't be able to invoke health check

<h4>Notes</h4>
 - The payload for **creating** or **updating** a user should look like this(Make sure you include the correct id when updating):
   {
      "user": {
         "country": "UK",
         "email": "example@example.com",
         "first_name": "John",
         "last_name": "Smith",
         "nickname": "MySampleNickName",
         "password": "MyStrongPassword"
      }
   }
 - The payload for **deleting** or **getting** a user should look like this:
   {
      "id": "someID"
   }
 - The payload for **listing** users should look like this:
   {
       "country": "UK",
       "page": 1,
       "page_size": 5
   }

<h3> Design Choices and Assumptions </h3>

1. gRPC and Protocol Buffers -->  higher-performance communication over HTTP/2 
2. In-Memory Storage --> required 
3. Dependency Injection --> eg obj UserStore into the UserService 
4. Health Checks --> required 
5. Logging + comments --> essential for debugging and readability

<h3>Possible Extensions and Improvements</h3>

1. Persistent Storage
    - Use a no SQL database or RDBMS for storing user data.
2. Error Handling and Validation
    - Implement request validations for having meaningful error messages and status codes that will assist in debugging and readability.   
3. Security Enhancements
    - Implement an authentication system that will also give access only for the requests containing valid tokens.
4. Testing
   - Build end-to-end tests for validating all incoming changes and making sure that each change satisfies all scenarios. 