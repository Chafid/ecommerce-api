# ecommerce-api
This is a simple e-commerce cart api with four function:
- Add item to the cart
- Delete item from the cart
- Update item from the cart
- Show items on the cart

SETTING UP THE API
1. Clone the API from github
   - go to the local directory for src where the GOPATH is configured
   - run the command
     # git clone https://github.com/chafid/ecommerce-api
2. Set up the mysql database for this api
   - Create a database name ecommerce-example
   - set up the username/password for the database as ecom/ecom123
   - create the table with the SQL:
   
   CREATE TABLE `cart` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `cart_id` int(11) NOT NULL,
     `item_name` varchar(255) NOT NULL,
     `total_price` int(11) NOT NULL,
     `num_of_item` int(11) NOT NULL,
      PRIMARY KEY (`id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1

3. go the src directory (just a level above the ecommerce-api directory) and run this command:
   # go run ecommerce-api

4. Test the api using POSTMAN
   - To add item to the cart:
     choose POST and enter the new item data in json format. For example:
     {"cart_id":1,"item_name":"table","total_price":300000,"num_of_item":1}
     it will add 1 table to the cart 1, with the price 300.000
    - To delete item in the cart:
      choose DELETE and use the cart id and item id that you want to delete
      /cart/{cart_id}/{id}. For example:
      /cart/1/2
     - To get item from cart
      choose GET and use the cart id with this format: /cart/{cart_id}. For example:
      /cart/1
      - Update the item on the cart
      choose PUT and use the cart id and the item id with this format /cart/{cart_id}/{id}. For example:
      /cart/1/2. And use the new data in json format, like this example: {"total_price":50000,"num_of_time":2}

