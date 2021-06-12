Assignment 3: Make a transaction in bank account

Acceptance Criteria: 

Write an API to create a new transaction for an existing customer
    
    - transaction can only be "withdrawal" or "deposit"
    
    - amount cannot be negative
    
    - withdrawal amount should be available in the account
    
    - successful transaction, should return the updated balance with transaction id response
    
    - error handling should be done for bad request, validation and unexpected errors from the server side and should return the appropriate http status code with message