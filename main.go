package main
import(
	"fmt"
	"net/http"
)
const portNumber = ":8080"
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(4,5)
	fmt.Fprintf(w,fmt.Sprintf( "This is the about page and 4 + 5 is %d", sum))
}

func addValues(x,y int) int{
	return x + y
}
func main()  {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about",About)

	fmt.Println("Starting application on port ", portNumber)

	http.ListenAndServe(portNumber,nil)
}