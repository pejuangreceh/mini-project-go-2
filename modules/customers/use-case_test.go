package customers

import (
	"crud_api/entities"
	mocks "crud_api/mocks/modules/customers"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		customer *entities.Customers
	}
	err := errors.New("Database Error")
	nilreq := entities.Customers{}
	reqcase := entities.Customers{
		Model:     gorm.Model{},
		FirstName: "Yoi",
		LastName:  "Bro",
		Email:     "yoibro@gmail.com",
		Avatar:    "no avatar la",
	}
	mockRepository := mocks.NewRepository(t)
	mockRepository.EXPECT().
		Save(&nilreq).
		Return(err).
		Once()
	mockRepository.EXPECT().
		Save(&reqcase).
		Return(nil).
		Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on Create Data",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{&nilreq},
			wantErr: true,
		},
		{
			name: "success on Create Data",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{&reqcase},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.Create(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_Read(t *testing.T) {
	type fields struct {
		repo Repository
	}
	err := errors.New("Database Error")
	req := []entities.Customers{
		{
			Model:     gorm.Model{},
			FirstName: "Yabes",
			LastName:  "Ganteng",
			Email:     "elloyyabest@gmail.com",
			Avatar:    "no avatar bruh",
		},
	}
	mockRepository := mocks.NewRepository(t)
	//test gagal
	mockRepository.EXPECT().
		GetAll().
		Return(nil, err).
		Once()
	mockRepository.EXPECT().
		GetAll().
		Return(req, nil).
		Once()
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "failed",
			fields:  fields{repo: mockRepository},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "success",
			fields:  fields{repo: mockRepository},
			want:    req,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ReadID(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID string
	}
	err := errors.New("Database Error")
	nilreq := []entities.Customers{}
	req := []entities.Customers{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			FirstName: "Yabes",
			LastName:  "Ganteng",
			Email:     "elloyyabest@gmail.com",
			Avatar:    "no avatar bruh",
		},
	}
	mockRepository := mocks.NewRepository(t)
	//test gagal
	mockRepository.EXPECT().
		FindByID(nil).
		Return(nilreq, err).
		Once()
	mockRepository.EXPECT().
		FindByID(1).
		Return(req, nil).
		Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "failed",
			fields:  fields{repo: mockRepository},
			args:    args{},
			want:    nilreq,
			wantErr: true,
		}, {
			name:    "success",
			fields:  fields{repo: mockRepository},
			args:    args{ID: "1"},
			want:    req,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.ReadID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Update(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		body entities.Customers
		ID   string
	}
	err := errors.New("ID not insterted")
	nilreq := entities.Customers{}
	req := entities.Customers{
		FirstName: "Yabes",
		LastName:  "Ganteng",
		Email:     "elloyyabest@gmail.com",
		Avatar:    "no avatar bruh",
	}
	res := entities.Customers{
		FirstName: "Yabes",
		LastName:  "Ganteng",
		Email:     "elloyyabest@gmail.com",
		Avatar:    "no avatar bruh",
	}

	mockRepository := mocks.NewRepository(t)
	//test gagal
	mockRepository.EXPECT().
		UpdateByID(req, "").
		Return(&nilreq, err).
		Once()
	mockRepository.EXPECT().
		UpdateByID(req, "1").
		Return(&res, nil).
		Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "fail to update",
			fields: fields{
				repo: mockRepository,
			},
			args: args{
				body: req,
				ID:   "",
			},
			want:    &nilreq,
			wantErr: true,
		}, {
			name: "success to update",
			fields: fields{
				repo: mockRepository,
			},
			args: args{
				body: req,
				ID:   "1",
			},
			want:    &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Update(tt.args.body, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_Delete(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID string
	}
	err := errors.New("ID not insterted")
	res := entities.Customers{
		FirstName: "Yabes",
		LastName:  "Ganteng",
		Email:     "elloyyabest@gmail.com",
		Avatar:    "no avatar bruh",
	}

	mockRepository := mocks.NewRepository(t)
	//test gagal
	mockRepository.EXPECT().
		DeleteByID("").
		Return(nil, err).
		Once()
	mockRepository.EXPECT().
		DeleteByID("1").
		Return(&res, nil).
		Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "failed to delete, null id",
			fields:  fields{repo: mockRepository},
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "successfuly delete",
			fields:  fields{repo: mockRepository},
			args:    args{"1"},
			want:    &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Delete(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
