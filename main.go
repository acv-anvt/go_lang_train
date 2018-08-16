package main
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
  "strconv"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}
const (
  host     = "localhost"
  port     = 5432
  user     = "acv"
  password = ""
  dbname   = "proj_doshin_client_dev_3"
)

func main() {
  psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
    user, password, host, dbname)

  db, err := gorm.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  for i := 1; i < 10; i++ {
    db.Create(&Product{Code: fmt.Sprintf("L1212%s",strconv.Itoa(i)), Price: 10000})
  }
  

  // Read
  var product Product
  db.First(&product, 1) // find product with id 1
  db.First(&product, "code = ?", "L1212") // find product with code l1212

  // Update - update product's price to 2000
  db.Model(&product).Update("Price", 2000)

  // Delete - delete product
  db.Delete(&product)
}