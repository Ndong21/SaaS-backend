# SaaS Backend API Documentation
Here is a list of all the api endpoints that can be used with this backend and what each do.

## Module 1: Stocks

## POST /products
Create a new product

**Request**
```js
{
  "category_id":"uuid",
  "product_name":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"Product created successfully"
}

// if the request is bad
{
  "status":"error",
  "message":"Failed to create product. Product name is required"
}
```
## GET /products
List of all the product

**Response**
```js
[
  {
    "Product_id":"uuid",
    "product_name":"string",
    "category_name":"string",
    "quantity_left":"number"
  }
]
```
## GET /categories
List of all the categories

**Response**
```js
[
  {
    "category_id":"uuid",
    "category_name":"string",
    "category_description":"string"
  }
]
```
## POST /categories
Create a new category

**Request**
```js
{
  "category_name":"string",
  "category_description":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"category created successfully"
}

// if the request is bad
{
  "status":"error",
  "message":"Failed to create category. category name is required"
}
```

## POST /purchases
Create a new product purchases

**Request**
```js
{
  "product_id":"uuid",
  "total_price":"number",
  "quantity":"number",
  "vendor_id":"uuid",
  "cashier_id":"uuid"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"purchase created successfully"
}
```
## GET /purchases
List of all the purchases

**Response**
```js
[
  {
    "product_id":"uuid",
    "product_name":"string",
    "total_price":"number"
    "quantity":"number",
    "vendor_name":"string"
  }
]
```
## GET /sales
List of all the sales made

**Response**
```js
[
  {
    "product_id":"uuid",
    "product_name":"string",
    "unit_price":"number"
    "quantity":"number",
    "total_price":"number"
    "cashier_id":"uuid"
  }
]
```
## POST /sales
Create a new sale

**Request**
```js
{
  "product_id":"uuid",
  "total_price":"number",
  "quantity":"number",
  "cashier_id":"uuid"
}
```
## POST /catalog
Create a new catalog item

**Request**
```js
{
  "product_id":"uuid",
  "unit_price":"number"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"catalog item created successfully"
}
```
## GET /catalog
List of all the catalog items

**Response**
```js
[
  {
    "product_id":"uuid",
    "product_name":"string",
    "unit_price":"number"
  }
]
```
## DELETE /catalog/:product_id
Remove an item from the catalog

**Response**
```js
// sample response
{
  "message":"catalog item successfully deleted",
  "status":"success"
}
```
## Module 1: Block Production Unit

## POST /blocks/products
Create a new block type

**Request**
```js
{
  "product_name":"string",
  "product_description":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"Product created successfully"
}

// if the request is bad
{
  "status":"error",
  "message":"Failed to create product. Product name is required"
}
```
## GET /blocks/products
List of all the products produced

**Response**
```js
{
 products: [
    {
      "product_id":"uuid",
      "product_name":"string",
      "description":"string",
      "quantity_left":"number"
    }
  ],
  "message":"products successfully retrieved",
  "status":"success"
}
```
## POST /blocks/materials
Create a new material

**Request**
```js
{
  "material_name":"string",
  "unit":"string"
  "product_description":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"production material created successfully"
}
```
## GET /blocks/products
List of all the materials

**Response**
```js
{
 materials: [
    {
      "material_id":"uuid",
      "material_name":"string",
      "unit":"string",
      "description":"string"
    }
  ],
  "message":"materials successfully retrieved",
  "status":"success"
}
```
## POST /blocks/session
Create a new working session

**Request**
```js
{
  "session_name":"string",
  "description":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"session created successfully"
}
```
## GET /blocks/session
List of all the sessions

**Response**
```js
{
 sessions: [
    {
      "session_id":"uuid",
      "session_name":"string",
      "description":"string"
    }
  ],
  "message":"sessions successfully retrieved",
  "status":"success"
}
```
## POST /blocks/purchase
Create a new material purchase

**Request**
```js
{
  "material_id":"uuid",
  "quantity":"number",
  "price":"number"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"material purchase created successfully"
}
```
## GET /blocks/purchase
List of all the material purchases

**Response**
```js
{
 purchases: [
    {
      "id":"uuid",
      "material_id":"uuid",
      "quantity":"number",
      "price":"number",
      "unit":"string",
      "purchase_date":"date"
    }
  ],
  "message":"purchases successfully retrieved",
  "status":"success"
}
```
## POST /blocks/team
Create a new team

**Request**
```js
{
  "team_name":"string",
  "phone_number":"string",
  "email":"string"
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"team created successfully"
}
```
## GET /blocks/team
List of all the production team

**Response**
```js
{
 teams: [
    {
      "id":"uuid",
      "team_name":"string",
      "phone_number":"string",
      "email":"string"
    }
  ],
  "message":"teams successfully retrieved",
  "status":"success"
}
```
## POST /blocks/sale
Create a new product sale

**Request**
```js
{
  "product_id":"uuid",
  "selling_price":"number",
  "quantity":"number",
  "cashier_id":"uuid",
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"sale created successfully"
}
```
## GET /blocks/sale
List of all the sales
**Response**
```js
{
 sales: [
    {
      "id":"uuid",
      "product_name":"string",
      "quantity":"number",
      "selling_price":"number"
      "sale_date":"date"
    }
  ],
  "message":"sales successfully retrieved",
  "status":"success"
}
```
## POST /blocks/sale
Create a new product sale

**Request**
```js
{
  "product_id":"uuid",
  "selling_price":"number",
  "quantity":"number",
  "cashier_id":"uuid",
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"sale created successfully"
}
```
## POST /blocks/session/material
Create a new session material

**Request**
```js
{
  "session_id":"uuid",
  "team_id":"uuid",
  "material_id":"uuid",
  "date":"date",
  "quantity":"number",
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"session material created successfully"
}
```
## POST /blocks/session/product
Create a new session product

**Request**
```js
{
  "session_id":"uuid",
  "team_id":"uuid",
  "product_id":"uuid",
  "date":"date",
  "quantity":"number",
}
```
**Response**
```js
// if successfull
{
  "status":"success",
  "message":"session product created successfully"
}
```
## GET /blocks/session/material
List of all the session materials

**Response**
```js
{
 session_materials: [
    {
      "id":"string",
      "session_name":"string",
      "team_name":"string",
      "material_name":"string",
      "quantity":"number",
      "unit":"string",
      "used_date":"date"
    }
  ],
  "message":"session material successfully retrieved",
  "status":"success"
}
```
## GET /blocks/session/product
List of all the session products

**Response**
```js
{
 session_products: [
    {
      "id":"string",
      "session_name":"string",
      "team_name":"string",
      "product_name":"string",
      "quantity":"number",
      "used_date":"date"
    }
  ],
  "message":"session products successfully retrieved",
  "status":"success"
}
```
