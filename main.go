package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=====================================================\n", sql)
}

func main() {
	dsn := "root:P@ssw0rd@tcp(192.168.1.8:3306)/gorm_basic?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		// DryRun: true,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Gender{}, Test{})

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
