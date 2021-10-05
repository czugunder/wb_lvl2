package listing5

func Start() {
	mymain()
}

type customError struct {
	msg string
}

func (e *customError) Error() string { // являет собой соответствие интерфейсу error, поэтому в 24 строчке присваивается
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func mymain() {
	var err error   // переменная тип error
	err = test()    // вернулся интерфейс, динамическое значение nil, тип *customErr, раз есть тип то != nil
	if err != nil { // заходит сюда т.к. см. строчку выше
		println("error") // строка error в stdout
		return           // завершается main
	}
	println("ok")
}

// error

// Подробно описал в listing/listing3, кратко описал в комментариях
