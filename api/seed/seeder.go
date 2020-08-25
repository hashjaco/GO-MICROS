package seed

import(
	"github.com/hashjaco/GO-MICROS/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models.User{
	{
		FirstName: "Jack",
		LastName:  "Ripper",
		Email:     "jackripper187@murder.you",
		Username:  "ripperJack2020",
		Password:  "password",
		},
	{
		FirstName: "Darla",
		LastName:  "Singer",
		Email:     "dsinger1up@mindspring.com",
		Username:  "dsinger1up",
		Password:  "password",
		},
	{
		FirstName: "Dude",
		LastName:  "Sweet",
		Email:     "sweetdude@dope.net",
		Username:  "dudesweet444",
		Password:  "password",
		},

}


func Load(db *gorm.DB){

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}