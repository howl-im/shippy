package go_micro_srv_user
import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("create uuid failed: %v\n", err)
	}

	log.Printf("set id by BeforeCreate func\n")
	return scope.SetColumn("Id", uuid.String())
}
