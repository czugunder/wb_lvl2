package listing4

func Start() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

// deadlock, так как после 10 проходов цикла writer'a - он умрет, а receiver нет, соответственно
// рутина с writer'ом умрет, а main с ридером останется и будет ожидать что в канал что-то кинут
