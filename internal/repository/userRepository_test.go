package repository

/*import (
	"github.com/google/uuid"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserRepository(db)

	type args struct {
		uuId uuid.UUID
		user model.User
	}

	type mockBehavior func(m sqlmock.Sqlmock, args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				uuId: uuid.New(),
				user: model.User{
					Name:       "test name",
					OfficeUuid: uuid.New(),
				},
			},
		},
	}
}*/
