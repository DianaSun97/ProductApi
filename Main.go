package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type Product struct {
	Id   int
	Name string
}

var database *sql.DB // global variable for database access

// listProducts from the database.
func listProducts() ([]Product, error) {
	rows, err := database.Query("select * from productdb.product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// IndexHandler lists all products and show them.
func IndexHandler(c *gin.Context) {
	products, err := listProducts()
	if err != nil {
		log.Println(err)
		return
	}
	tmpl, _ := template.ParseFiles("templates/index.gohtml")
	tmpl.Execute(c.Writer, products)
}

// CreateGetHandler shows a form for adding data.
func CreateGetHandler(c *gin.Context) {
	c.File("templates/create.gohtml")
	//http.ServeFile(w, r, "templates/create.gohtml")
}

// CreatePostHandler accepts POST request with a form, extract the necessary elements and adds it to the database.
func CreatePostHandler(c *gin.Context) {
	Id, _ := c.GetPostForm("Id")
	Name, _ := c.GetPostForm("Name")

	_, err := database.Exec("insert into productdb.product (id, name) values (?, ?)",
		Id, Name)

	if err != nil {
		log.Println(err)
		return
	}

	// Redirect to site root
	c.Redirect(http.StatusMovedPermanently, "/")
	//http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// ProductsHandler is an API to list products.
func ProductsHandler(c *gin.Context) {
	products, err := listProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, products)
}

// ProductsAddHandler is an API to add products.
func ProductsAddHandler(c *gin.Context) {
	// Reading product request.
	var p Product
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// Inserting to the database.
	res, err := database.Exec("insert into productdb.product (name) values (?)", p.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Get new ID from the database.
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Update ID in the product variable.
	p.Id = int(id)
	// Send it as a response.
	c.JSON(http.StatusOK, p)
}

func main() {
	// Open a database connection.
	db, err := sql.Open("mysql", "root:admin@/productdb")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// After the connection is opened, the value of the database variable is set.
	database = db

	// Setting all routes.
	router := gin.Default()
	// Routes for the forms.
	router.GET("/", IndexHandler)
	router.GET("/create", CreateGetHandler)
	router.POST("/create", CreatePostHandler)
	// Routes for the API.
	router.GET("/api/products", ProductsHandler)
	router.POST("/api/products", ProductsAddHandler)

	// Old code of the test task using net/http package.
	//http.HandleFunc("/", IndexHandler)
	//http.HandleFunc("/create", CreateHandler)

	fmt.Println("Server is listening...")
	router.Run(":8181")
	//http.ListenAndServe(":8181", nil)
}
