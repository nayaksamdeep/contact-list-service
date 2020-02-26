package Models

import (
	"../Config"
	"fmt"
	_ "github.com/lib/pq"
)

func GetAllContacts(contact *[]Contact) (err error) {
	if err = Config.DB.Find(contact).Error; err != nil {
		return err
	}
	return nil
}

func CreateAContact(contact *Contact) (err error) {
	if err = Config.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func GetAContact(contact *Contact, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(contact).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAContact(contact *Contact, id string) (err error) {
	fmt.Println(contact)
	Config.DB.Save(contact)
	return nil
}

func DeleteAContact(contact *Contact, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(contact)
	return nil
}