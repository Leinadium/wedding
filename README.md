# wedding / api

## API Spec

```text
POST /v1/guest/
<- {name, phone}
-> 201 ""
-> 400 {error}
-> 500 {error}

GET /v1/product
-> 200 {products: {id, name, imageURL, priceBRL}}

POST /v1/product
<- products: {name, imageURL, priceBRL}
-> {id}

GET /v1/product/<id>/payment
-> 200 {url}

GET /v1/purchase/
-> 200 []{email, productID, price}

POST /v1/purchase/
<- {?}
-> 201 ""
-> 500 {error}
```

## DB Spec

```text
Product:
- stripe_id
- name
- image_url
- price_brl
- price_stripe_id
- purchased

Guest:
- name
- phone
- registerDate
```
