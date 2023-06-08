package customers

import (
	"crud_api/entities"
	mocks "crud_api/mocks/modules/customers"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		customer *entities.Customers
	}
	err := errors.New("Database Error")

	req := entities.Customers{
		Model:     gorm.Model{},
		FirstName: "Yoii",
		LastName:  "Bro",
		Email:     "yoibro@gmail.com",
		Avatar:    "no avatar la",
	}
	mockRepository := mocks.NewRepository(t)
	mockRepository.EXPECT().
		Save(&req).
		Return(err).
		Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on Get All Data",
			fields: fields{
				repo: mockRepository,
			},
			args:    args{&req},
			wantErr: true,
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

func TestUseCase_Delete(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
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

func TestUseCase_Read(t *testing.T) {
	type fields struct {
		repo Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
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
