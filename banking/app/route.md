Routes

admin can do all 
> Get All Customers     GET /customer
> Get Customer by IDs   GET /customer/{customer_id}
> Create new account    POST /customer/{customer_id}/account
> Make a transaction    POST /customer/{customer_id}/transaction

user role can do 
> Get Customer by IDs   GET /customer/{customer_id}
> Make a transaction    POST /customer/{customer_id}/transaction