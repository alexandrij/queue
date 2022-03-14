# Очередь

Очередь (англ. queue) - структура данных в информатике, в которой элементы хранятся в порядке их добавления. Добавление
новых элементов(enqueue)
осуществляется в начало списка. А удаление элементов (dequeue)
осуществляется с конца. Таким образом очередь реализует принцип
"первым вошёл - первым вышел" (FIFO). Часто реализуется операция чтения головного элемента (peek), которая возвращает
первый в очереди элемент, при этом не удаляя его. Очередь является примером линейной структуры данных или
последовательной коллекции.

Иллюстрация работы с очередью.

![Очередь](https://upload.wikimedia.org/wikipedia/commons/5/52/Data_Queue.svg)

## Пример

```go
package main

import (
	"fmt"
	"github.com/alexandrij/queue"
)

func main() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		fmt.Println(v)
	}
	fmt.Println("end")
}
```

## References

- [Wikipedia](https://ru.wikipedia.org/wiki/%D0%9E%D1%87%D0%B5%D1%80%D0%B5%D0%B4%D1%8C_(%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5))
- [YouTube](https://www.youtube.com/watch?v=GRsVMTlBIoE)