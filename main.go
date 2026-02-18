package main
import(
	//"fmt"
	"net/http"
	//"hello/greetings"  // import path = go module specified in go.mod + file paths
	//"log"
	//"rsc.io/quote"
	 "github.com/gin-gonic/gin"
)

// album model
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

//data
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
//handler to return album items
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	err:=c.BindJSON(&newAlbum)
	if err!=nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated,newAlbum)
}

func getAlbumByID(c *gin.Context){
	reqID := c.Param("id")
	for _,val := range albums {
    if val.ID == reqID {
		c.IndentedJSON(http.StatusOK,val)
		return
	}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func main() {
	 router := gin.Default()  // initializes a gin router using gin.Default()
    router.GET("/albums", getAlbums)  // GET function to associate the getAlbums http method to the /albums path
	router.POST("/albums",postAlbums)
	router.GET("/albums/:id",getAlbumByID)
    router.Run("localhost:8080")
}