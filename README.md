<h3> Running the application </h3>

1. 22
2. 222
3. 22
4. 22

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