
package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
};

func main()  {
	// 创建一个新的结构体, 可以使用key => value形式 忽略的字段为 0 或 空
	var Books1 Books
	Books1 = Books{"go", "haha", "go haha",123 };
	// Books2 = Books{ title: "go", author: "auth", book_id: 212 }
	// Books3 = Books{title: "Go 语言", author: "www.runoob.com"};

	printBook(Books1)
}

func printBook(book Books) {
	fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}