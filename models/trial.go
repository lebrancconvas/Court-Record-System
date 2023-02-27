package model

import (
	"lebrancconvas/courtrecord/db"
	form "lebrancconvas/courtrecord/forms"
	"log"
)

func CreateTrialTable() {
	db := db.GetDB()
	err := db.AutoMigrate(form.Trial{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetTrials() {

}

func GetTrialByID(id int) {

}

func CreateTrial() {

}

func UpdateTrial(id int) {

}

func DeleteTrial(id int) {
	
}