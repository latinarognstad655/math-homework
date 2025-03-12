
  package main
  
import "fmt"
  import "math/rand"
  func main() {
      fmt.Println("Welcome to Math Homework")
      
     rand.Seed(time.Now().UnixNano())
      var x int = rand.Intn(100) + 1
     var y int = rand.Intn(100) + 1
      fmt.Println("Random numbers: ", x, "and", y)
      
    var sum int = x + y
     var difference int = x - y
      var product int = x * y
       var quotient int = x / y
        if x > y {
         fmt.Println("The greater number is: ", x)
          } else {
           fmt.Println("The greater number is: ", y)
            }
             fmt.Println("Sum:", sum)
              fmt.Println("Difference:", difference)
               fmt.Println("Product:", product)
                fmt.Println("Quotient:", quotient)
                  }