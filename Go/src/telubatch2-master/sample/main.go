package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Customer struct (Model) ...
type Customer struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}

// Get all orders

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customers []Customer

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var customer Customer
		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		customerID := r.FormValue("CustomerID")
		companyName := r.FormValue("CompanyName")
		contactName := r.FormValue("ContactName")
		contactTitle := r.FormValue("ContactTitle")
		address := r.FormValue("Address")
		city := r.FormValue("City")
		region := r.FormValue("Region")
		postalCode := r.FormValue("PostalCode")
		country := r.FormValue("Country")
		phone := r.FormValue("Phone")
		fax := r.FormValue("Fax")

		stmt, err := db.Prepare("INSERT INTO customers (CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax) VALUES (?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(customerID, companyName, contactName, contactTitle, address, city, region, postalCode, country, phone, fax)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []Customer
	params := mux.Vars(r)

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers WHERE CustomerID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var customer Customer

	for result.Next() {

		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}

		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newCompanyName := r.FormValue("CompanyName")

		stmt, err := db.Prepare("UPDATE customers SET CompanyName = ? WHERE CustomerID = ?")

		_, err = stmt.Exec(newCompanyName, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Customer with CustomerID = %s was updated", params["id"])
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM customers WHERE CustomerID = ?")

	_, err = stmt.Exec(params["id"])

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "Customer with ID = %s was deleted", params["id"])
}

func delCustomer(w http.ResponseWriter, r *http.Request) {

	customerID := r.FormValue("CustomerID")
	companyName := r.FormValue("CompanyName")

	stmt, err := db.Prepare("DELETE FROM customers WHERE CustomerID = ? AND CompanyName = ?")

	_, err = stmt.Exec(customerID, companyName)

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "Customer with ID = %s was deleted", customerID)
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customers []Customer

	CustomerID := r.FormValue("CustomerID")
	CompanyName := r.FormValue("CompanyName")

	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers WHERE CustomerID = ? AND CompanyName = ?`

	result, err := db.Query(sql, CustomerID, CompanyName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var customer Customer

	for result.Next() {

		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}

		customers = append(customers, customer)
	}

	json.NewEncoder(w).Encode(customers)

}

//Book is ...
type Book struct {
	BookID        string `json:"CustomerID"`
	Title         string `json:"Title"`
	Author        string `json:"Author"`
	DatePublished string `json:"DatePublished"`
	Publisher     string `json:"Publisher"`
	Edition       string `json:"Edition"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	sql := `SELECT
				BookID,
				IFNULL(Title,''),
				IFNULL(Author,'') Author,
				IFNULL(DatePublished,'') DatePublished,
				IFNULL(Publisher,'') Publisher,
				IFNULL(Edition,'') Edition
			FROM book`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var book Book
		err := result.Scan(&book.BookID, &book.Title, &book.Author,
			&book.DatePublished, &book.Publisher, &book.Edition)

		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		bookID := r.FormValue("BookID")
		title := r.FormValue("Title")
		author := r.FormValue("Author")
		datePublished := r.FormValue("DatePublished")
		publisher := r.FormValue("Publisher")
		edition := r.FormValue("Edition")

		stmt, err := db.Prepare("INSERT INTO book (BookID,Title,Author,DatePublished,Publisher,Edition) VALUES (?,?,?,?,?,?)")

		_, err = stmt.Exec(bookID, title, author, datePublished, publisher, edition)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	params := mux.Vars(r)

	sql := `SELECT
				BookID ,
				IFNULL(Title,''),
				IFNULL(Author,'') Author,
				IFNULL(DatePublished,'') DatePublished,
				IFNULL(Publisher,'') Publisher,
				IFNULL(Edition,'') Edition
			FROM book WHERE BookID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var book Book

	for result.Next() {

		err := result.Scan(&book.BookID, &book.Title, &book.Author,
			&book.DatePublished, &book.Publisher, &book.Edition)

		if err != nil {
			panic(err.Error())
		}

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM book WHERE BookID = ?")

	_, err = stmt.Exec(params["id"])

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "book with ID = %s was deleted", params["id"])
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newTitle := r.FormValue("Title")
		// newAuthor := r.FormValue("Author")
		// newDatePublished := r.FormValue("DatePublished")
		// newPublisher := r.FormValue("Publisher")
		// newEdition := r.FormValue("Edition")

		stmt, err := db.Prepare("UPDATE book SET Title = ? WHERE BookID = ?")
		//stmt, err := db.Prepare("UPDATE book SET Title = ?, Author = ? , DatePublished = ?, Publisher = ?, Edition = ? WHERE BookID = ?")
		
		_, err = stmt.Exec(newTitle, params["id"])
		//_, err = stmt.Exec(newTitle,newAuthor,newDatePublished,newPublisher,newEdition, params["id"])
		// _, err = stmt.Exec(newAuthor, params["id"])
		// _, err = stmt.Exec(newDatePublished, params["id"])
		// _, err = stmt.Exec(newPublisher, params["id"])
		// _, err = stmt.Exec(newEdition, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Book with BookID = %s was updated", params["id"])
	}
}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	r.HandleFunc("/delcustomer", delCustomer).Methods("POST")
	//New
	r.HandleFunc("/getcustomer", getPost).Methods("POST")

	//Fungsi CRUD Untuk tabel book
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
