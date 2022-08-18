package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("need config file")
	}
	log.Println("start application")

	config, err := NewConfig(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	app := NewServer()

	server := &http.Server{
		Addr:         config.Server.Host + ":" + config.Server.Port,
		Handler:      app.NewRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Run server, host:", config.Server.Host, "port:", config.Server.Port)
	log.Fatalln(server.ListenAndServe())
}

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err = json.Unmarshal(file, config); err != nil {
		return nil, err
	}

	return config, nil
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(date []byte) error {
	if string(date) == "" || string(date) == "null" {
		*d = Date{time.Now()}
		return nil
	}

	tm, err := time.Parse(`"`+"2006-01-02"+`"`, string(date))
	*d = Date{tm}
	return err
}

type Model struct {
	UserId string `json:"user_id"`
	Date   Date   `json:"date"`
	ID     string `json:"event_id"`
	Desc   string `json:"desc"`
}

type App struct {
	sync.RWMutex
	Items map[string][]Model
}

func NewServer() *App {
	res := App{Items: make(map[string][]Model)}

	return &res
}

func (a *App) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", middleware(http.HandlerFunc(a.create)))
	mux.HandleFunc("/update_event", middleware(http.HandlerFunc(a.update)))
	mux.HandleFunc("/delete_event", middleware(http.HandlerFunc(a.delete)))
	mux.HandleFunc("/events_for_day", middleware(http.HandlerFunc(a.forDay)))
	mux.HandleFunc("/events_for_week", middleware(http.HandlerFunc(a.forWeek)))
	mux.HandleFunc("/events_for_month", middleware(http.HandlerFunc(a.forMonth)))

	return mux
}

func (a *App) create(w http.ResponseWriter, r *http.Request) {
	event, err := postEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.createItems(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (a *App) createItems(model *Model) ([]Model, error) {
	temp := a.Items[model.UserId]
	var err error

	a.RLock()
	for _, ev := range temp {
		if ev.ID == model.ID {
			err = errors.New("model already exist")
			break
		}
	}
	a.RUnlock()

	if err != nil {
		return nil, err
	}

	a.Lock()
	a.Items[model.UserId] = append(a.Items[model.UserId], *model)
	a.Unlock()
	return a.Items[model.UserId], nil
}

func (a *App) update(w http.ResponseWriter, r *http.Request) {
	event, err := postEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := a.updateEvent(event)

	writeResult(w, result)
}

func (a *App) updateEvent(model *Model) []Model {
	temp := a.Items[model.UserId]

	for i, t := range temp {
		if model.ID == t.ID {
			a.Lock()
			temp[i].Date = model.Date
			temp[i].Desc = model.Desc
			a.Unlock()
			break
		}
	}

	return a.Items[model.UserId]
}

func (a *App) delete(w http.ResponseWriter, r *http.Request) {
	event, err := postEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := a.deleteEvent(event)

	writeResult(w, result)
}

func (a *App) deleteEvent(model *Model) []Model {
	temp := a.Items[model.UserId]

	for i, t := range temp {
		if model.ID == t.ID {
			a.Lock()
			temp = append(temp[:i], temp[i+1:]...)
			a.Unlock()
		}
	}

	a.Lock()
	a.Items[model.UserId] = temp
	a.Unlock()
	return temp
}

func (a *App) forDay(w http.ResponseWriter, r *http.Request) {
	event, err := getEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.eventForDay(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (a *App) eventForDay(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := a.Items[model.UserId]

	a.RLock()
	defer a.RUnlock()
	for _, mod := range models {
		if mod.Date == model.Date {
			result = append(result, mod)
		}
	}

	return result, nil
}

func (a *App) forWeek(w http.ResponseWriter, r *http.Request) {
	event, err := getEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.eventForWeek(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (a *App) eventForWeek(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := a.Items[model.UserId]
	targetYear, targetWeek := model.Date.ISOWeek()

	a.RLock()
	defer a.RUnlock()

	for _, mod := range models {
		year, week := mod.Date.ISOWeek()
		if targetYear == year && targetWeek == week {
			result = append(result, mod)
		}
	}

	return result, nil
}

func (a *App) forMonth(w http.ResponseWriter, r *http.Request) {
	event, err := getEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.eventForMonth(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (a *App) eventForMonth(model *Model) ([]Model, error) {
	result := make([]Model, 0)
	models := a.Items[model.UserId]
	targetYear, targetMonth := model.Date.Year(), model.Date.Month()

	a.RLock()
	defer a.RUnlock()

	for _, mod := range models {
		year, month := mod.Date.Year(), mod.Date.Month()
		if targetYear == year && targetMonth == month {
			result = append(result, mod)
		}
	}

	return result, nil
}

func middleware(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: Method " + r.Method + ", Url " + r.RequestURI)
		fn.ServeHTTP(w, r)
	}
}

func postEvents(r *http.Request) (*Model, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("invalid method: " + r.Method)
	}

	return jsonDecode(r)
}

func jsonDecode(r *http.Request) (*Model, error) {
	event := Model{}

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func writeError(w http.ResponseWriter, msg string, code int) {
	err := struct {
		Err string `json:"error"`
	}{Err: msg}

	marshal, _ := json.Marshal(err)
	http.Error(w, string(marshal), code)
	log.Println("error:", msg)
}

func writeResult(w http.ResponseWriter, res []Model) {
	result := struct {
		Result []Model `json:"result"`
	}{res}

	marshal, _ := json.Marshal(result)
	_, err := w.Write(marshal)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadGateway)
	}
}

func getEvent(r *http.Request) (*Model, error) {
	timeReq, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		return nil, err
	}

	// собираем модель эвента из queryString
	m := Model{
		Date:   Date{timeReq},
		UserId: r.URL.Query().Get("user_id"),
	}
	return &m, nil
}
