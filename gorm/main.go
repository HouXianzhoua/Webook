package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm/logger"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

//MAIN使用GORM设置SQLite数据库，迁移产品模型的架构，
//并演示基本的CRUD操作：创建产品，通过主键阅读它
//或特定字段，更新其字段并通过主键删除它。
func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 设置日志级别
  db.Logger = db.Logger.LogMode(logger.LogLevel(4))

  // 迁移 schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})
  println("Created product with code D42 and price 100")

  // Read
  var product Product
  db.First(&product, 1) // 根据整型主键查找
  println("Found product with id 1")
  
  db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
  println("Found product with code D42")

  // Update - 将 product 的 price 更新为 200
  db.Model(&product).Update("Price", 200)
  println("Updated product with id 1, set price to 200")
  // Update - 更新多个字段
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
  println("Updated product with id 1, set price to 200 and code to F42")
  // Update - 更新所有字段
  db.Model(&product).Updates(map[string]any{"Price": 200, "Code": "F42"})
  println("Updated product with id 1, set price to 200 and code to F42")

  // Delete - 删除 product
  db.Delete(&product, 1)
  println("Deleted product with id 1")
}
