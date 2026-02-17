package main
import(
	"fmt"
	"hello/greetings"  // import path = go module specified in go.mod + file paths
	"log"
	"rsc.io/quote"
)
func main() {
    fmt.Println("Hello, World!")
	// calling functions from external package
	// go mod tidy which located and downloaded the rsc.io/quote module that contains the package i imported
	 fmt.Println(quote.Go())
	
	 log.SetPrefix("greetings: ")
     log.SetFlags(0)
	 message, err := greetings.Hello("Shubhra") // error handling
	  if err != nil {
        log.Fatal(err)
    }

     fmt.Println(message)


	 names := []string{"Gladys", "Samantha", "Darrin"}
	 messages, err2 := greetings.Hellos(names)
	 if(err2!=nil){
		log.Fatal(err2);
	 }
	 fmt.Println(messages);
	 
}