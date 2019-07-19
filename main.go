package main

func main() {
    a := App{} 
    // You need to set your Username and Password here
    a.Initialize("", "", "ecommerce_example")

    a.Run(":8080")
}