package account

import (
	"crud_api/entities"
	mocks "crud_api/mocks/modules/account"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestUseCase_Activate(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		body entities.Activate
		ID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Activate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Activate(tt.args.body, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Activate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Activate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Approval(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		body entities.Approval
		ID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Approval
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Approval(tt.args.body, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Approval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approval() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		actor *entities.Actors
	}
	err := errors.New("Database Error")
	nilreq := entities.Actors{}
	reqcase := entities.Actors{
		Model:      gorm.Model{},
		Username:   "admin",
		Password:   "123456",
		RoleID:     2,
		IsVerified: "false",
		IsActive:   "false",
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
			if err := u.Create(tt.args.actor); (err != nil) != tt.wantErr {
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
		want    *entities.Actors
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

func TestUseCase_Login(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Login(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Read(t *testing.T) {
	type fields struct {
		repo Repository
	}
	err := errors.New("Database Error")
	req := []entities.Actors{
		{
			Model:      gorm.Model{},
			Username:   "admin",
			Password:   "123456",
			RoleID:     2,
			IsVerified: "false",
			IsActive:   "false",
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
		want    []entities.Actors
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
	nilreq := []entities.Actors{}
	req := []entities.Actors{
		{
			Model: gorm.Model{
				ID:        10,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Username:   "admin10",
			Password:   "admin10",
			RoleID:     2,
			IsVerified: "false",
			IsActive:   "false",
		},
	}
	mockRepository := mocks.NewRepository(t)
	//test gagal
	mockRepository.EXPECT().
		FindByID("").
		Return(nilreq, err).
		Once()
	mockRepository.EXPECT().
		FindByID("10").
		Return(req, nil).
		Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "failed",
			fields:  fields{repo: mockRepository},
			args:    args{ID: ""},
			want:    nilreq,
			wantErr: true,
		}, {
			name:    "success",
			fields:  fields{repo: mockRepository},
			args:    args{ID: "10"},
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
	err := errors.New("ID not insterted")
	nilreq := entities.Actors{}
	req := entities.Actors{
		Username:   "admin",
		Password:   "123456",
		RoleID:     2,
		IsVerified: "false",
		IsActive:   "false",
	}
	res := entities.Actors{
		Username:   "admin",
		Password:   "123456",
		RoleID:     2,
		IsVerified: "false",
		IsActive:   "false",
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
	type args struct {
		body entities.Actors
		ID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actors
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
