package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Author    string
	Name      string
	PageCount int
}

"""
gorm.Model을 Embedded 구조체로 사용하면, 
자동으로 ID, CreatedAt, UpdatedAt, DeletedAt 컬럼과 
idx_users_deleted_at 인덱스를 가진 테이블이 생성된다.
"""


