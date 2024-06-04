package main

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"github.com/ecodeclub/ekit/sqlx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestValuerScanner(t *testing.T) {
	// 有没有想过，我的 Go 语言的类型是怎么传递过去数据库上
	// 数据库上的数据，传回来之后，我这边是解析的
	// 比如说 sql.NullString{} 怎么变成 varchar?
	// Valuer: Go 类型 => 数据库/驱动可处理类型
	// Scanner: 数据库/驱动类型 => Go 类型

	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	// assert 是断言，不成立还是会继续执行
	// require 是要求，必须满足，不成立会 panic
	require.NoError(t, err)
	err = db.AutoMigrate(&UserV1{})
	require.NoError(t, err)
	err = db.Create(&UserV1{
		Id:    1,
		Name:  "Tom",
		Phone: "12345678901",
		Address: sqlx.JsonColumn[Address]{
			Valid: true,
			Val: Address{
				Province: "广东",
			},
		},
		Labels: sqlx.JsonColumn[[]string]{
			Valid: true,
			Val:   []string{"帅气", "穷苦"},
		},
	}).Error
	require.NoError(t, err)

	var u1 UserV1
	err = db.Where("id = ?", 1).First(&u1).Error
	require.NoError(t, err)

	var u2 UserV1
	err = db.Where("phone = ?", PhoneNumber("12345678901")).First(&u2).Error
	require.NoError(t, err)

	err = db.Exec("TRUNCATE table user_v1;").Error
	require.NoError(t, err)
}

type UserV1 struct {
	Id   int64 `gorm:"primaryKey,autoIncrement"`
	Name string
	// 假定说，我现在要你对 Phone 进行一个加密存储
	// 1. 可以解密的加密方案
	// 2. 在哪里加密解密 - 直接利用数据库，或者应用代码里面手动加密

	Phone PhoneNumber `gorm:"type:varchar(512)"`

	Address sqlx.JsonColumn[Address]  `gorm:"type:varchar(1024)"`
	Labels  sqlx.JsonColumn[[]string] `gorm:"type:varchar(1024)"`
}

type Address struct {
	Province string
}

type PhoneNumber string

// 这里绝对是指针
// src 是什么呢？是数据库过来的东西，具体类型，不要猜测
func (n *PhoneNumber) Scan(src any) error {
	switch val := src.(type) {
	case []byte:
		data, err := hex.DecodeString(string(val))
		if err != nil {
			return err
		}
		*n = PhoneNumber(data)
		return nil
	case string:
		data, err := hex.DecodeString(val)
		if err != nil {
			return err
		}
		*n = PhoneNumber(data)
		return nil
	}
	return errors.New("未知类型")
}

// 可以放回“任意”类型，但是这个类型必须要能够被 驱动处理
// 基本类型，[]byte，string，
func (n PhoneNumber) Value() (driver.Value, error) {
	// 我在这里加密，而后 sql.DB 拿到我的返回值，它会直接存进去数据库里面
	// 注意：对同一个数据加密，是同一个值
	return hex.EncodeToString([]byte(n)), nil
}
