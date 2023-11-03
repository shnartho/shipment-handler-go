package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"www.github.com/shnartho/shipment-handler-go/pkg/data"
)

var d = data.NewData()

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/index.html")
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/order.html")
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/add.html")
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/update.html")
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/remove.html")
}

func OrderPackHandler(w http.ResponseWriter, r *http.Request) {
	inputValue := r.FormValue("inputValue")
	intValue, _ := strconv.Atoi(inputValue)
	packsNeeded := data.PacksNeeded(d, intValue)

	temp, err := template.ParseFiles("../templates/orderpack.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, packsNeeded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func AddPackHandler(w http.ResponseWriter, r *http.Request) {
	addValue := r.FormValue("inputValue")
	intValue, err := strconv.Atoi(addValue)
	if err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	if err := data.AddToSlice(d, intValue); err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	http.ServeFile(w, r, "../templates/success.html")
}

func UpdatePackHandler(w http.ResponseWriter, r *http.Request) {
	cs := r.FormValue("currentsize")
	us := r.FormValue("updatesize")

	cs_int, err := strconv.Atoi(cs)
	if err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	up_int, err := strconv.Atoi(us)
	if err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	if err := data.UpdateSlice(d, cs_int, up_int); err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	http.ServeFile(w, r, "../templates/success.html")
}

func RemovePackHandler(w http.ResponseWriter, r *http.Request) {
	rp := r.FormValue("inputValue")
	intValue, err := strconv.Atoi(rp)
	if err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	if err := data.RemoveFromSlice(d, intValue); err != nil {
		http.ServeFile(w, r, "../templates/invalid.html")
		return
	}

	http.ServeFile(w, r, "../templates/success.html")
}

func GetSliceHandler(w http.ResponseWriter, r *http.Request) {
	sliceAsString := data.GetSliceValues(d)
	w.Write([]byte("Values:" + sliceAsString))
}
