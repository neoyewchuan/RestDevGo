Write an API to retrieve a customer by ID

Acceptance criteria:
(Part 1)
    - The URL should only accept numeric IDs
    - API should return customer as a JSON object
 (Part 2)
    - In case the customer id does not exist, API should return http status code 404 (Not Found)
    - In case of an unexpected error, API should return status code 500 (Internal Server Error) along with the error message