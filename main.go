package main

import (
	"fmt"
)

type Observer interface {
	handleEvent(v []string)
}

type Observable interface {
	subscribe()
	unsubscribe()
	sendAll()
}

type User struct {
	name string
}

type JobSite struct {
	vacancies   []string
	subscribers []User
}

func (u *User) handleEvent(v []string) {
	fmt.Printf("Hello, dear %v!\n", u.name)
	for _, value := range v {
		fmt.Printf("- %v\n", value)
	}
}

func (js *JobSite) sendAll() {
	for _, value := range js.subscribers {
		value.handleEvent(js.vacancies)
	}
}

func (js *JobSite) addVacancy(vacancy string) {
	js.vacancies = append(js.vacancies, vacancy)
	js.sendAll()
}

func (js *JobSite) removeVacancy(vacancy string) {
	tempVacancies := make([]string, 0)
	for _, value := range js.vacancies {
		if value != vacancy {
			tempVacancies = append(tempVacancies, value)
		}
	}
	js.vacancies = tempVacancies
	js.sendAll()
}

func (js *JobSite) subscribe(u User) {
	js.subscribers = append(js.subscribers, u)

}

func (js *JobSite) unsubscribe(u User) {
	tempSubscribers := make([]User, 0)
	for _, value := range js.subscribers {
		if value != u {
			tempSubscribers = append(tempSubscribers, value)
		}
	}
	js.subscribers = tempSubscribers

}

func main() {
	itgo := JobSite{
		subscribers: []User{},
		vacancies:   []string{},
	}
	user1 := User{name: "Erzhan"}
	user2 := User{name: "Nathan"}
	user3 := User{name: "Oleg"}
	itgo.subscribe(user1)
	itgo.subscribe(user2)
	itgo.subscribe(user3)
	itgo.addVacancy("UI/UX Designer")
	itgo.addVacancy("English Teacher")
	itgo.unsubscribe(user2)
	itgo.unsubscribe(user3)
	itgo.addVacancy("Aquaman")
	itgo.removeVacancy("English Teacher")

}
