package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"fullstack2024/db"
	"fullstack2024/model"
	"fullstack2024/redis"
	"fullstack2024/utils"

	"github.com/gorilla/mux"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var c model.Client
	_ = json.NewDecoder(r.Body).Decode(&c)

	c.ClientLogo = utils.UploadToS3Mock(c.ClientLogo)
	c.CreatedAt = time.Now().Format(time.RFC3339)

	query := `INSERT INTO my_client (name, slug, is_project, self_capture, client_prefix, client_logo, address, phone_number, city, created_at)
	          VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	err := db.DB.QueryRow(query, c.Name, c.Slug, c.IsProject, c.SelfCapture, c.ClientPrefix, c.ClientLogo, c.Address, c.PhoneNumber, c.City, c.CreatedAt).Scan(&c.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(c)
	redis.Rdb.Set(redis.Ctx, c.Slug, data, 0)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	var c model.Client
	_ = json.NewDecoder(r.Body).Decode(&c)
	params := mux.Vars(r)
	id := params["id"]
	c.UpdatedAt = time.Now().Format(time.RFC3339)

	query := `UPDATE my_client SET name=$1, slug=$2, is_project=$3, self_capture=$4, client_prefix=$5, client_logo=$6,
	          address=$7, phone_number=$8, city=$9, updated_at=$10 WHERE id=$11`
	_, err := db.DB.Exec(query, c.Name, c.Slug, c.IsProject, c.SelfCapture, c.ClientPrefix,
		utils.UploadToS3Mock(c.ClientLogo), c.Address, c.PhoneNumber, c.City, c.UpdatedAt, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(c)
	redis.Rdb.Set(redis.Ctx, c.Slug, data, 0)

	w.WriteHeader(http.StatusOK)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	deletedAt := time.Now().Format(time.RFC3339)

	_, err := db.DB.Exec("UPDATE my_client SET deleted_at=$1 WHERE id=$2", deletedAt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slug := r.URL.Query().Get("slug")
	redis.Rdb.Del(redis.Ctx, slug)
	w.WriteHeader(http.StatusOK)
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, slug FROM my_client WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clients []model.Client
	for rows.Next() {
		var c model.Client
		rows.Scan(&c.ID, &c.Name, &c.Slug)
		clients = append(clients, c)
	}
	json.NewEncoder(w).Encode(clients)
}
