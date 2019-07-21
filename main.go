package main

func main() {
    a := App{} 
    // You need to set your Username and Password here
    a.Initialize("ecom", "ecom123", "ecommerce_example")

    a.Run(":8080")
}