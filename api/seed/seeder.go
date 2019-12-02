package seed

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/nstoker/apiance1/api/models"
)

var users = []models.User{
	models.User{
		Name:     "Neil Stoker",
		Email:    "neil.stoker@koiostechnology.co.uk",
		Password: "changeme",
	},
	models.User{
		Name:     "Test User",
		Email:    "test.user@example.com",
		Password: "password",
	},
}

// Load loads the seeds
func Load(db *sqlx.DB) error {

	users := []models.User{
		{
			Name:     "Neil Stoker",
			Email:    "neil.stoker@koiostechnology.co.uk",
			Password: "changeme",
		},
	}

	for _, u := range users {
		if db == nil {
			log.Fatal("db is null")
		}
		_, err := u.FindUserByEmail(db, u.Email)
		if err == nil {
			// Found the user, so
			continue
		} else if err == sql.ErrNoRows {
			_, err := u.CreateUser(db)
			if err != nil {
				return nil
			}

			continue

		} else {
			log.Fatal("Seeding stopped %w", err)
		}
		// switch err {
		// case sql.ErrNoRows:
		// 	// Add the record
		// 	_, err := u.CreateUser(db)
		// 	if err != nil {
		// 		return err
		// 	}
		// case nil:
		// 	// User record exists, so leave it alone
		// 	continue
		// default:
		// 	return err
		// }
	}

	return nil
	// err := db.Debug().DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }

	// // err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// // if err != nil {
	// // 	log.Fatalf("attaching foreign key error: %v", err)
	// // }

	// for i := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}
	// 	// posts[i].AuthorID = users[i].ID

	// 	// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	// if err != nil {
	// 	// 	log.Fatalf("cannot seed posts table: %v", err)
	// 	// }
	// }
}
