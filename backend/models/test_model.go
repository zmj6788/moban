package models

// TestModel 用户表
type TestModel struct {
	MODEL
	Name string `gorm:"size:36" json:"name"` // 用户名
	Age  string `gorm:"size:36" json:"age"`  // 用户名
}
