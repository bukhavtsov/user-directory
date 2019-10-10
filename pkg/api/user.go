package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bukhavtsov/user-directory/pkg/data"
	"github.com/gorilla/mux"
)

type UserData interface {
	Create(user *data.User) (int64, error)
	Read(id int64) (*data.User, error)
	ReadAll() ([]*data.User, error)
	Update(user *data.User) (*data.User, error)
	Delete(id int64) (int64, error)
}

type userAPI struct {
	data UserData
}

func ServeUserResource(r *mux.Router, data UserData) {
	api := &userAPI{data: data}
	r.HandleFunc("/users", api.getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", api.getUser).Methods("GET")
	r.HandleFunc("/users", api.createUser).Methods("POST")
	r.HandleFunc("/users/{id}", api.updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", api.deleteUser).Methods("DELETE")
	r.HandleFunc("/upload/users/{id}", api.updateIcon).Methods("PUT")
	r.HandleFunc("/", api.serveTemplate).Methods("GET")
	r.PathPrefix("/assets/images").Handler(http.StripPrefix("/assets/images", http.FileServer(http.Dir("./assets/images/"))))

}

func (api userAPI) getUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := api.data.ReadAll()
	if err != nil {
		log.Println("users haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api userAPI) getUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := api.data.Read(id)
	if err != nil {
		log.Println("user hasn't been read")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api userAPI) createUser(writer http.ResponseWriter, request *http.Request) {
	user := new(data.User)
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if user == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := api.data.Create(user)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/users/%d", userId))
	writer.WriteHeader(http.StatusCreated)
}

func (api userAPI) updateUser(writer http.ResponseWriter, request *http.Request) {
	user := new(data.User)
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Id = id
	updatedUser, err := api.data.Update(user)
	if err != nil {
		log.Println("user hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(updatedUser)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api userAPI) deleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	_, err = api.data.Delete(id)
	if err != nil {
		log.Println("user hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (api userAPI) updateIcon(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error Retrieving the File: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	tempFile, err := ioutil.TempFile("assets/images", "upload-*.png")
	if err != nil {
		log.Printf("failed method ioutil.TempFile: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("failed method ioutil.ReadFile: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = tempFile.Write(fileBytes)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filename := tempFile.Name()

	user, err := api.data.Read(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Img = filename
	_, err = api.data.Update(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (api userAPI) serveTemplate(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
