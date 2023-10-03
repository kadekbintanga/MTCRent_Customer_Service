package main

import (
	"Service/App/Models/Activity/Support"
	"Service/App/Models/Testing"
	Server "Service/App/Services/Constant/Activity"
	"Service/Config"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.Init()

	var testing Testing.Testing
	var activity Support.UseActivity

	err = Config.PgSQL.First(&testing, 2).Error
	if err != nil {
		fmt.Println(err)
	}

	activity = activity.SetSubFeature("testing-activity").
		SetModel(testing).SetOldProperty(Server.ACTION_GENERAL, "onlyName")

	testing.Name = "Fourth Update"
	err = Config.PgSQL.Save(&testing).Error
	if err != nil {
		fmt.Println(err)
	}

	activity.SetModel(testing).SetNewProperty(Server.ACTION_GENERAL, "onlyName").
		Save("Object update activity")
}
