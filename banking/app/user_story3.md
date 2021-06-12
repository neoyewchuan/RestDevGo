Write an API to create a new account for an existing customer.

Acceptance criteria:

    - A new account can only be opened with a minimum deposit of S$50.00
    - Account can only be of saving or current type
    - In case of an unexpected error, API should return status code 500 (Internal Server Error) along with the error message
    - The API should return the new account id, when the new account is opened with the status code as 201 (CREATED)
