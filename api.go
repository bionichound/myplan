package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type TimeStamp struct {
	CheckOut time.Time
	CheckIn  time.Time
}

type Task struct {
	Created    time.Time
	Modified   time.Time
	TimeStamps []TimeStamp
	Open       bool
	Title      string
	Message    string
}

func NewTask(title string, message string) *Task {
	return &Task{
		Created:  time.Now(),
		Modified: time.Now(),
		Open:     true,
		Title:    title,
		Message:  message,
	}
}

type Store struct {
	Tasks []*Task
	Done  []*Task
}

var store *Store
var storeFileName string

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	storeFileName = dirname + "/.myplan/store"
	_, err = os.Stat(storeFileName)
	if os.IsNotExist(err) {
		os.Create(storeFileName)
	}
	store = &Store{
		Tasks: []*Task{},
		Done:  []*Task{},
	}
	ReadStore()
}

func PrintStore() {
	fmt.Println("Items for today")
	for _, item := range store.Tasks {
		fmt.Println(item.Modified.Format(time.Stamp), "-", item.Title, ":", item.Message)
	}
	fmt.Println("Items done")
	for _, item := range store.Done {
		fmt.Println(item.Modified.Format(time.Stamp), "-", item.Title, ":", item.Message)
	}
}

func PrintEnumerated() {
	for ind, item := range store.Tasks {
		fmt.Println(ind, item.Title, ":", item.Message)
	}
}

func MarkAsDone(ind int) {
	store.Done = append(store.Done, store.Tasks[ind])
	RemoveFromStore(ind)
}

func SaveStore() {
	content, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(storeFileName, content, 0644)
}

func ReadStore() {
	content, err := os.ReadFile(storeFileName)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	if string(content) != "" {
		err = json.Unmarshal(content, &store)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func AddToStore(item *Task) {
	store.Tasks = append(store.Tasks, item)
	SaveStore()
}

func RemoveFromStore(ind int) {
	store.Tasks[ind] = store.Tasks[len(store.Tasks)-1]
	store.Tasks = store.Tasks[:len(store.Tasks)-1]
}
