package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/database/postgres"
	"testing"
	"time"
)

func connectToTestPostgresDB() (*sqlx.DB, error) {
	db, err := postgres.NewPostgresDB(&postgres.Config{
		PgHost:         "localhost",
		PgPort:         "5432",
		PgUser:         "myusername",
		PgDBName:       "mydatabase",
		PgDBSchemaName: "public",
		PgPwd:          "mypassword",
		PgSSLMode:      "disable",
	})

	return db, err
}

func TestOfficeRepository_CreateOffice(t *testing.T) {
	db, err := connectToTestPostgresDB()
	assert.NoError(t, err, "Failed to connect to database")

	officeRepo := NewOfficeRepository(db)

	type args struct {
		uuId   uuid.UUID
		office model.Office
	}

	var testTable = []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				uuId: uuid.New(),
				office: model.Office{
					Name:    "test name",
					Address: "test address",
				},
			},
			wantErr: false,
		},
		{
			name: "Missing uuid",
			args: args{
				office: model.Office{
					Name:    "test name",
					Address: "test address",
				},
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := officeRepo.CreateOffice(testCase.args.uuId, testCase.args.office)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOfficeRepository_GetOffice(t *testing.T) {
	db, err := connectToTestPostgresDB()
	assert.NoError(t, err, "Failed to connect to database")

	officeRepo := NewOfficeRepository(db)

	expectedTimeString := "2023-06-13T12:34:56Z"
	expectedTime, _ := time.Parse(time.RFC3339, expectedTimeString)
	expectedTime = expectedTime.In(time.FixedZone("", 0))

	validUUID := uuid.New()

	type args struct {
		officeUUID uuid.UUID
	}

	var testTable = []struct {
		name           string
		args           args
		expectedResult model.Office
		wantErr        bool
	}{
		{
			name: "OK",
			args: args{
				officeUUID: validUUID,
			},
			expectedResult: model.Office{
				Uuid:      validUUID,
				Name:      "New Office",
				Address:   "New Address",
				CreatedAt: expectedTime,
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Здесь этого быть не должно. Вынести предварительное добавление данных в другое место
			/*const createOfficeQuery = `insert into offices (uuid, name, address, created_at) values ($1, $2, $3, $4)`
			_, err := officeRepo.db.Exec(createOfficeQuery,
				testCase.expectedResult.Uuid, testCase.expectedResult.Name, testCase.expectedResult.Address, testCase.expectedResult.CreatedAt,
			)
			assert.NoError(t, err)*/

			office, err := officeRepo.GetOffice(testCase.args.officeUUID)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedResult, office)
			}
		})
	}
}

func TestOfficeRepository_GetOfficesList(t *testing.T) {
	db, err := connectToTestPostgresDB()
	assert.NoError(t, err, "Failed to connect to database")

	officeRepo := NewOfficeRepository(db)

	validUUID := uuid.New()
	expectedTimeString := "2023-06-13T12:34:56Z"
	expectedTime, _ := time.Parse(time.RFC3339, expectedTimeString)
	expectedTime = expectedTime.In(time.FixedZone("", 0))

	var testTable = []struct {
		name           string
		expectedResult []*model.Office
		wantErr        bool
	}{
		{
			name: "OK",
			expectedResult: []*model.Office{
				{
					Uuid:      validUUID,
					Name:      "Test name 1",
					Address:   "Test address 1",
					CreatedAt: expectedTime,
				},
				{
					Uuid:      validUUID,
					Name:      "Test name 2",
					Address:   "Test address 2",
					CreatedAt: expectedTime,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			offices, err := officeRepo.GetOfficesList()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedResult, offices)
			}
		})
	}
}
