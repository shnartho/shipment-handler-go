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
	intValue, _ := strconv.Atoi(addValue)
	data.AddToSlice(d, intValue)
	http.ServeFile(w, r, "../templates/add.html")
}

func UpdatePackHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/updatepack.html")
}

func RemovePackHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../templates/removepack.html")
}

func GetSliceHandler(w http.ResponseWriter, r *http.Request) {
	sliceAsString := data.GetSliceValues(d)
	w.Write([]byte("Values:" + sliceAsString))
}
