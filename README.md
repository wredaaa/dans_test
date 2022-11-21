# Query Test

1. Recapitulation of number of accounts owned by every customer.
```
select customer.*, count(account.acc_number) as recap_acc 
from customer
inner join account on customer.cust_id = account.acc_owner
group by cust_id
```

2. All transactions created by John Michael sorted by account number and transaction date

```
select CONCAT(customer.cust_firstname, ' ', customer.cust_lastname) as full_name, account.acc_number, transaction.trs_date, transaction.trs_amount, transaction.trs_type  
from customer
inner join account on customer.cust_id = account.acc_owner
inner join `transaction` on account.acc_number = transaction.trs_from_account
where CONCAT(customer.cust_firstname, ' ', customer.cust_lastname) = 'John Michael'
order by account.acc_number asc, transaction.trs_date
```

# Golang Test 

1. Create database `dans` or change database name in `db.go line 16`
2. run project


## API Reference

#### Register

```http
  POST /signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. |
| `email` | `string` | **Required**. |
| `password` | `string` | **Required**. |
| `role` | `string` | **Required**. |


#### Login

```http
  POST /signin
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. |
| `password` | `string` | **Required**. |



#### Get list job

```http
  GET /jobs
  GET /jobs?description=python&location=berlin
  GET /jobs?page=1
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Token`      | `string` | **Required**. add Token in Headers |


#### Get job detail

```http
  GET jobs/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. |
| `Token`      | `string` | **Required**. add Token in Headers |

