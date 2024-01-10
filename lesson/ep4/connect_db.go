package lesson

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"      // or the Docker service name if running in another container
	port     = 5432             // default PostgreSQL port
	user     = "mydearuser"     // as defined in docker-compose.yml
	password = "mydearpassword" // as defined in docker-compose.yml
	dbname   = "mydeardatabase" // as defined in docker-compose.yml
)

var db *sql.DB

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Supplier_id int    `json:"supplier_id"`
}

type Suppliers struct {
	ID          int
	Name        string
}

type ProductWithSupplier struct {
	ProductID        int
	ProductName      string
	Price            int
	SupplierName     string
  }
func Connect_DB() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb


	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}


//lesson.Connect_DB()

	//lesson.Ep4_createProduct("nooode", 555, 1)

	//lesson.Ep4_selectProduct(8)

	//lesson.Ep4_updateProduct(11,"jjjjj",1111)

	//lesson.Ep4_deleteProductt(19)

	//lesson.Ep4_getAllProduct()

	//lesson.Ep4_getJoinProduct()