package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	chInt := make(chan int)
	chUser := make(chan user, n)
	for i := 0; i < int(pool); i++ {
		go func(chInt <-chan int, chUser chan<- user) {
			for {
				chUser <- getOne(int64(<-chInt))
			}
		}(chInt, chUser)
	}
	for i := 0; i < int(n); i++ {
		chInt <- i
	}
	for len(res) < int(n) {
		res = append(res, <-chUser)
	}
	return
}
