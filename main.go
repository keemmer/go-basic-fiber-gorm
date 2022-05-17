package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=====================================================\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:P@ssw0rd@tcp(192.168.1.8:3306)/gorm_basic?parseTime=true"
	dial := mysql.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		// DryRun: true,
	})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(Gender{}, Test{})
	// CreateGender("Female")
	// GetGenders()
	// GetGender(1)
	// GetGender(10)
	// GetGenderByname("Male")
	// UpdateGender(1, "xxxxx")
	// UpdateGender2(1, "rrr")
	// DeleteGender(1)

	// CreateTest(0, "test1")
	// CreateTest(0, "test2")
	// CreateTest(0, "test3")

	// DeleteTest(1)

	// GetTest()

	// db.Migrator().CreateTable(Customer{})
	// db.AutoMigrate(Gender{}, Test{},Customer{})

	// CreateCustomer("mou",3)
	GetCustomer()
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func CreateCustomer(name string, genderID uint) {
	customer := &Customer{
		Name:     name,
		GenderID: genderID,
	}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomer(){
	customers:= []Customer{}
	// tx:=db.Find(&customers)
	// tx:=db.Preload("Gender").Find(&customers)
	tx:=db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _,customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID,customer.Name,customer.Gender.Name)
	}

	fmt.Println(customers)
}







func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenders() {
	genders := []Gender{}
	// tx:=db.Creat(&genders)
	tx := db.Order("id desc").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	genders := []Gender{}
	// tx:=db.Find(&genders)
	tx := db.First(&genders, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGenderByname(name string) {
	genders := []Gender{}
	// tx := db.Order("id").Find(&genders,"name=?",name)
	// tx := db.Find(&genders,"name=?",name)
	tx := db.Where("name=?", name).Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.Order("id").First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	db.Save(&gender)
	GetGender(id)
}
func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&gender).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

func CreateTest(code uint, name string) {
	test := Test{Code: code, Name: name}
	db.Create(&test)
}
func GetTest() {
	tests := []Test{}
	db.Find(&tests)
	fmt.Println(tests)
}
func DeleteTest(id uint) {
	// db.Delete(&Test{}, id)
	db.Unscoped().Delete(&Test{}, id)
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size:20"`
}

type Test struct {
	gorm.Model
	Code uint   `gorm:"primaryKey;comment:this is Code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

// type Gender struct {
// 	gorm.Model
// 	Code uint   `gorm:"primaryKey;comment:this is Code"`
// 	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
// }

// func (g Gender) TableName() string {
// 	return "MyGender"
// }
