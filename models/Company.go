package models

import (
    "fmt"
)

type Company struct {
    ID int `gorm:"primary_key";"AUTO_INCREMENT"`
    RemoteId int `gorm:"column:remoteId"`
    Name string
    Ticker string
    LinkedInId int `gorm:"column:linkedInId"`
    Status string
    ApprovalStatus string `gorm:"column:approvalStatus"`
}

func (Company *Company) TableName() string {
    return "company"
}

type CompanyDuplicateError struct {
    name string
}

func (c *CompanyDuplicateError) Error() string {
    return fmt.Sprintf("another company with name: '%s' already exists", c.name)
}

type DuplicateCompanyRemoteIdError struct {
    remoteId int
}

func (c *DuplicateCompanyRemoteIdError) Error() string {
    return fmt.Sprintf("another company with remote id: '%d' already exists", c.remoteId)
}