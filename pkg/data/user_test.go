package data

import (
	"reflect"
	"testing"

	"github.com/bukhavtsov/user-directory/db"
	"github.com/jinzhu/gorm"
)

var (
	host     = "localhost"
	port     = "postgres"
	user     = "postgres"
	dbname   = "postgres"
	password = "postgres"
	sslmode  = "disable"
)

func TestNewUserData(t *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *UserData
	}{
		{
			name: "NewUserData with valid db connection",
			args: args{db: conn},
			want: &UserData{db: conn},
		},
		{
			name: "NewUserData with nil db connection",
			args: args{db: nil},
			want: &UserData{db: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserData(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t1 *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()

	lastUser := &User{}

	conn.Last(lastUser)
	expectedId := lastUser.Id + 1

	if err := conn.Last(lastUser).Error; err != nil {
		t1.Error("getLast user:", err)
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Create with valid params",
			fields: fields{db: conn},
			args: args{
				user: &User{
					FirstName: "test_first_name",
					LastName:  "test_last_name",
					Img:       "assets/images/user_icon_1.png",
				}},
			want:    expectedId,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := UserData{
				db: tt.fields.db,
			}
			got, err := t.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t1 *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()

	lastUser := &User{}

	conn.Last(lastUser)

	user := &User{
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Img:       "assets/images/user_icon_1.png",
	}

	data := NewUserData(conn)
	id, err := data.Create(user)
	if err != nil {
		t1.Error("create user err:", err)
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "delete user with valid args",
			fields:  fields{db: conn},
			args:    args{id: id},
			want:    id,
			wantErr: false,
		},
		{
			name:    "delete user with incorrect args",
			fields:  fields{db: conn},
			args:    args{id: -1},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := UserData{
				db: tt.fields.db,
			}
			got, err := t.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRead(t1 *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()

	lastUser := &User{}

	conn.Last(lastUser)
	user := &User{
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Img:       "assets/images/user_icon_1.png",
	}

	data := NewUserData(conn)
	id, err := data.Create(user)
	user.Id = id
	if err != nil {
		t1.Error("create user err:", err)
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "get user with valid args",
			fields:  fields{db: conn},
			args:    args{id: id},
			want:    user,
			wantErr: false,
		},
		{
			name:    "get user with incorrect args",
			fields:  fields{db: conn},
			args:    args{id: -1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := UserData{
				db: tt.fields.db,
			}
			got, err := t.Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t1 *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()

	lastUser := &User{}

	conn.Last(lastUser)
	user := &User{
		FirstName: "test_first_name",
		LastName:  "test_last_name",
		Img:       "assets/images/user_icon_1.png",
	}

	data := NewUserData(conn)
	id, err := data.Create(user)
	user.Id = id
	if err != nil {
		t1.Error("create user err:", err)
	}

	new := &User{
		Id:        user.Id,
		FirstName: "updated_first_name",
		LastName:  "updated_last_name",
		Img:       "assets/images/user_icon_2.png",
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id  int64
		new *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:   "update user with valid args",
			fields: fields{db: conn},
			args: args{
				new: new,
			},
			want:    new,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := UserData{
				db: tt.fields.db,
			}
			got, err := t.Update(tt.args.new)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
