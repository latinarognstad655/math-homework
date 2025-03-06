
import "math/rand"

func main() {
    // Generate a random number between 1 and 100
    n := rand.Intn(100) + 1
    fmt.Println("Random number:", n)
}