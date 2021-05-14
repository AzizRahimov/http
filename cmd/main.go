package main

import (
	"fmt"
	"github.com/AzizRahimov/pkg/server"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	host := "localhost"
	port := "9999"

	if err := execute(host, port); err != nil{
		os.Exit(1)

	}
	fmt.Println("server closed")
}

//func execute(host, port string) (err error)  {
//	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer func() {
//		if cerr := listener.Close(); cerr != nil{
//			if err == nil{
//				err = cerr
//				return
//				}
//				log.Print(err)
//
//		}
//		}()
//	//TODO: server code
//
//	for {
//		conn, err := listener.Accept() // он возращает нам интерфейс
//		if  err != nil {
//			log.Print(err)
//			continue
//		}
//		//err = conn.Close()
//		//if err != nil {
//		//	log.Print(err)
//		//	continue
//		//}
//		err  = handle(conn)
//		if err != nil {
//			log.Print(err, "Ошибка возникла тут")
//		}
//		}
//
//
//
//
//	return
//
//}
//
//func handle(conn net.Conn) (err error) {
//	defer func() {
//		if cerr := conn.Close(); cerr != nil{
//			if err == nil{
//				err =cerr
//				return
//			}
//			log.Print(err)
//		}
//	}()
//	//TODO: handle coonection
//
//	buf := make([]byte, 4096)
//
//		//TODO: buf - в него от 0 до какого опредленного элемента внутри него будут данные в виде байт
//		// a N - нам нужен, чтобы понять сколько элементов он прочел, и чтобы мы до туда именно читали....
//
//		n, err := conn.Read(buf)
//		if err  == io.EOF {
//			log.Printf("%s", buf[:n])
//			fmt.Println("прочти то что было в ошибке")
//			return nil
//		}
//
//		if err != nil {
//			return err
//
//		}
//		log.Printf("%s", buf[:n])
//		fmt.Println("а это просто внутри цикла")
//		fmt.Println("ЭТО У НАС N", n)  // сколько та байтов удалось запихнуь  в buf
//		data :=  buf[:n]
//		requestLineDelim := []byte{'\r', '\n'}
//
//		//TODO: помоему эта хрень, нужна чтобы отделять такие вещи как GET / HTTP1.1
//		requestLineEnd := bytes.Index(data, requestLineDelim)
//		fmt.Println(requestLineEnd, "смотри что получили") // 14
//
//		if requestLineEnd == -1{
//			log.Println("ошибка")
//			return err
//		}
//		// тут помоему полученные байты должны преобразовать в тип string чтобы мы смогли увидеть, что там
//
//
//		requestLine := string(data[:requestLineEnd]) // requestLineEnd - он возвращает всего 3 эллемента: метод, путь, протокол
//		// они оказывается слиты и мы должны их разделить
//		//  2 аргумент показывает, если разделить их то что хочешь в итоге поставить....
//		fmt.Println(requestLine, "Тут должно быть 3 вещи")
//		parts := strings.Split(requestLine, " ")
//		if len(parts) != 3{
//			log.Print("Parts не равняется 3", parts)
//			return err
//		}
//
//		method, path, version := parts[0], parts[1], parts[2]
//		if method != "GET"{
//			fmt.Println("твой метод не GET")
//			return err
//		}
//
//		if version != "HTTP/1.1"{
//			fmt.Println("проблема с версией")
//			return err
//		}
//		if path == "/"{
//			body, err := os.ReadFile("static/index.html")
//			if err != nil {
//
//				return fmt.Errorf("can't read index.hmtl: %w", err)
//			}
//
//			_, err = conn.Write([]byte (
//				"HTTP/1.1 200 OK\r\n" +
//				"Content-Length: " + strconv.Itoa(len(body)) +  "\r\n" +
//					"Content-Type: text/html\r\n" +
//					"Connection: close\r\n" +
//					"\r\n" +
//					string(body),
//					))
//			if err != nil {
//				log.Print("ошибка возникла при записи в body")
//				return err
//			}
//
//
//		}
//
//
//
//
//	return err
//}



//func execute(host, port string) (err error) {
//	srv := server.NewServer(net.JoinHostPort(host, port))
//	srv.Register("/", func(conn net.Conn) {
//		body := "Welcome to our web-site"
//		_, err = conn.Write([]byte (
//						"HTTP/1.1 200 OK\r\n" +
//						"Content-Length: " + strconv.Itoa(len(body)) +  "\r\n" +
//							"Content-Type: text/html\r\n"  +
//							"Connection: close\r\n" +
//							"\r\n" +
//							body,
//							))
//		if err != nil {
//			log.Print(err)
//
//		}
//			//TODO: написать логику
//	})
//	//srv.Register("/about", func(conn net.Conn) {
//	//	//TODO: написать логику
//	//	body := "About Golang Academy"
//	//	_, err = conn.Write([]byte (
//	//		"HTTP/1.1 200 OK" + CRLS +
//	//			"Content-Length: " + strconv.Itoa(len(body)) +  CRLS +
//	//			"Content-Type: text/html" + CRLS +
//	//			"Connection: close" + CRLS +
//	//			CRLS +
//	//			body,
//	//	))
//	//
//	//})
//	srv.Register("/about", func(conn net.Conn) {
//		body := "About Golang Academy"
//
//		_, err = conn.Write([]byte(
//			"HTTP/1.1 200 OK \r\n" +
//				"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
//				"Content-Type: text/html\r\n" +
//				"Connection: close\r\n" +
//				"\r\n" +
//				body,
//		))
//		if err != nil {
//			log.Println(err)
//		}
//
//	})
//	return srv.Start()
//
//	// тут просто запускаем сервер и все....
//
//}


func execute(host string, port string) (err error)  {
	srv := server.NewServer(net.JoinHostPort(host, port))
	srv.Register("/", func(conn net.Conn) {
		body := "Welcome to our web-site"

		log.Print(err)
		_, err = conn.Write([]byte(
			"HTTP/1.1 200 OK \r\n" +
				"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
				"Content-Type: text/html\r\n" +
				"Connection: close\r\n" +
				"\r\n" +
				body,
		))


	})
	srv.Register("/about", func(conn net.Conn) {
		body := "About Golang Academy"

		_, err = conn.Write([]byte(
			"HTTP/1.1 200 OK \r\n" +
				"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
				"Content-Type: text/html\r\n" +
				"Connection: close\r\n" +
				"\r\n" +
				body,
		))
		if err != nil {
			log.Println(err)
		}

	})
	return srv.Start()
}
