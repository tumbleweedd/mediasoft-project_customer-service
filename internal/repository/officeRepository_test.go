package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"regexp"
	"testing"
	"time"
)

const (
	createOfficeQuery = "insert into offices"
)

func TestOfficeRepository_CreateOffice(t *testing.T) {
	db, mock, r := setupMockDB()
	defer db.Close()

	type args struct {
		uuId   uuid.UUID
		office model.Office
	}

	type mockBehavior func(args args)

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
				office: model.Office{
					Name:    "test name",
					Address: "test address",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(createOfficeQuery).
					WithArgs(args.uuId, args.office.Name, args.office.Address).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "MISSING UUID",
			args: args{
				office: model.Office{
					Name:    "test name",
					Address: "test address",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(createOfficeQuery).
					WithArgs(args.uuId, args.office.Name, args.office.Address).
					WillReturnError(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "MISSING NAME OFFICE",
			args: args{
				uuId: uuid.New(),
				office: model.Office{
					Address: "test address",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(createOfficeQuery).
					WithArgs(args.uuId, args.office.Name, args.office.Address)
			},
			wantErr: true,
		},
		{
			name: "MISSING ADDRESS",
			args: args{
				uuId: uuid.New(),
				office: model.Office{
					Name: "test name",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(createOfficeQuery).
					WithArgs(args.uuId, args.office.Name, args.office.Address)
			},
			wantErr: true,
		},
		{
			name: "MISSING NAME AND ADDRESS",
			args: args{
				uuId:   uuid.New(),
				office: model.Office{},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec(createOfficeQuery).
					WithArgs(args.uuId, args.office.Name, args.office.Address)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err := r.CreateOffice(testCase.args.uuId, testCase.args.office)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOfficeRepository_GetOffice(t *testing.T) {
	db, mock, r := setupMockDB()
	defer db.Close()

	type args struct {
		officeUuid uuid.UUID
	}

	type mockBehavior func(args args, expectedOffice model.Office)

	testTable := []struct {
		name           string
		mockBehavior   mockBehavior
		args           args
		expectedOffice model.Office
		wantErr        bool
	}{
		{
			name: "OK",
			args: args{
				officeUuid: uuid.New(),
			},
			mockBehavior: func(args args, expectedOffice model.Office) {
				mock.ExpectQuery("select \\* from customer.offices where uuid=\\$1").
					WithArgs(args.officeUuid).
					WillReturnRows(
						sqlmock.NewRows([]string{"uuid", "name", "address", "created_at"}).
							AddRow(expectedOffice.Uuid, expectedOffice.Name, expectedOffice.Address, expectedOffice.CreatedAt),
					)
			},
			expectedOffice: model.Office{
				Uuid:      uuid.New(),
				Name:      "Test Office",
				Address:   "Test Address",
				CreatedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "INVALID UUID",
			args: args{
				officeUuid: uuid.MustParse("f47ac10b-58cc-4372-a567-000000000000"),
			},
			mockBehavior: func(args args, expectedOffice model.Office) {
				mock.ExpectQuery("select * from customer.offices where uuid=$1").
					WithArgs(args.officeUuid).
					WillReturnError(errors.New("invalid UUID"))
			},
			expectedOffice: model.Office{},
			wantErr:        true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.expectedOffice)

			office, err := r.GetOffice(testCase.args.officeUuid)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOffice, office)
			}
		})
	}
}

func TestOfficeRepository_GetOfficesList(t *testing.T) {
	db, mock, r := setupMockDB()
	defer db.Close()

	type mockBehavior func(expectedOfficesList []*model.Office)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedOfficesList []*model.Office
		wantErr             bool
	}{
		{
			name: "OK",
			mockBehavior: func(expectedOfficesList []*model.Office) {
				rows := sqlmock.NewRows([]string{"uuid", "name", "address", "created_at"})

				for _, office := range expectedOfficesList {
					rows.AddRow(office.Uuid, office.Name, office.Address, office.CreatedAt)
				}

				mock.ExpectQuery("select uuid, name, address, created_at from offices").
					WillReturnRows(rows)
			},
			expectedOfficesList: []*model.Office{
				{
					Uuid:      uuid.New(),
					Name:      "Office 1",
					Address:   "Address 1",
					CreatedAt: time.Now(),
				},
				{
					Uuid:      uuid.New(),
					Name:      "Office 2",
					Address:   "Address 2",
					CreatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name: "EMPTY RESULT",
			mockBehavior: func(expectedOfficesList []*model.Office) {
				rows := sqlmock.NewRows([]string{"uuid", "name", "address", "created_at"})

				mock.ExpectQuery("select uuid, name, address, created_at from offices").
					WillReturnRows(rows)
			},
			expectedOfficesList: nil,
			wantErr:             false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.expectedOfficesList)

			offices, err := r.GetOfficesList()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOfficesList, offices)
			}
		})
	}
}

func TestOfficeRepository_GetOfficeByUserUUID(t *testing.T) {
	db, mock, r := setupMockDB()
	defer db.Close()

	type args struct {
		userUUID uuid.UUID
	}

	type mockBehavior func(expectedOffice model.Office, args args)

	testTable := []struct {
		name           string
		mockBehavior   mockBehavior
		args           args
		expectedOffice model.Office
		wantErr        bool
	}{
		{
			name: "OK",
			mockBehavior: func(expectedOffice model.Office, args args) {
				const query = `select o.uuid, o.name, o.address
					from customer.offices o
							 join users u on o.uuid = u.office_uuid
					where u.uuid = $1`
				escapedQuery := regexp.QuoteMeta(query)
				mock.ExpectQuery(escapedQuery).
					WithArgs(args.userUUID).
					WillReturnRows(
						sqlmock.NewRows([]string{"uuid", "name", "address"}).
							AddRow(expectedOffice.Uuid, expectedOffice.Name, expectedOffice.Address))
			},
			args: args{
				userUUID: uuid.New(),
			},
			expectedOffice: model.Office{
				Uuid:    uuid.New(),
				Name:    "Test Office",
				Address: "Test Address",
			},
			wantErr: false,
		},
		{
			name: "User does not exist or not associated with an office",
			mockBehavior: func(expectedOffice model.Office, args args) {
				const query = `select o.uuid, o.name, o.address
					from customer.offices o
							 join users u on o.uuid = u.office_uuid
					where u.uuid = $1`
				escapedQuery := regexp.QuoteMeta(query)
				mock.ExpectQuery(escapedQuery).
					WithArgs(args.userUUID).
					WillReturnError(errors.New("some error"))
			},
			args: args{
				userUUID: uuid.New(),
			},
			expectedOffice: model.Office{},
			wantErr:        true,
		},
	}

	for _, testCase := range testTable {
		testCase.mockBehavior(testCase.expectedOffice, testCase.args)

		office, err := r.GetOfficeByUserUUID(testCase.args.userUUID)
		if testCase.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedOffice, office)
		}
	}
}

func setupMockDB() (*sqlx.DB, sqlmock.Sqlmock, *OfficeRepository) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	r := NewOfficeRepository(db)

	return db, mock, r
}
