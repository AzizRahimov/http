package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

// помнишь это type Handlerfucn func(responsWriter, request)
type HandlerFunc func(conn net.Conn)

type Server struct {
	addr string
	mu sync.RWMutex
	handlers map[string]HandlerFunc
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandlerFunc)}
}

func (s *Server) Register(path string, handler HandlerFunc)  {
	// вот тут у нас значит проблема
	s.mu.Lock()
	defer s.mu.Unlock()
	// мы даем ему ключ который возьмем из пути и значением будет handler
	s.handlers[path] = handler

}

//func (s *Server) Register(path string, handler HandlerFunc)  {
//	s.mu.Lock()
//	defer s.mu.Unlock()
//	s.handlers[path]=handler
//}

//func (s *Server) Start() error  {
//	//TODO: 	start server on host & port
//	//TODO: что я тут должен начать??
//	listener, err := net.Listen("tcp", s.addr)
//	if err != nil {
//		log.Println("ошибка при подключении")
//		return err
//	}
//
//	defer func() {
//		err = listener.Close()
//		if err != nil {
//			log.Print(err)
//		}
//	}()
//
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println("клиент не смог подключится, идем дальше...")
//			continue
//		}
//
//		// conn который мы даем вот это и нужно присвоить
//		go s.handle(conn)
//	}
//	return nil
//}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}

	defer func() {
		err = listener.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return err
		}

		go s.handle(conn)
	}

	return nil
}






//func (s *Server)  handle(conn net.Conn)  {
//
//	defer func() {
//		err := conn.Close()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//	}()
//	//TODO: тут  говорится, что я могу сразу  искать... но ведь это не так
//	buf := make([]byte, 4096)
//	n, err := conn.Read(buf)
//	if err == io.EOF{
//		fmt.Printf("%s", buf[:n])
//		return // помоему я должен это остановить.. не?
//	}
//	log.Printf("%s", buf[:n])
//	data :=  buf[:n]
//	requestLineDelim := []byte{'\r', '\n'}
//	requestLineEnd := bytes.Index(data, requestLineDelim)
//	if requestLineEnd == -1{
//		log.Println("-1")
//		return
//	}
//	// тут мы должны получить 3 данных
//	requestLine := string(data[:requestLineEnd]) // она возвращает 3 элемента, только они слитыне
//	parts := strings.Split(requestLine, " ")
//	if len(parts) != 3{
//		fmt.Println("длинна не 3")
//		return
//	}
//	method, path, version := parts[0], parts[1], parts[2]
//	if method != "GET"{
//		fmt.Println("проблема с методом")
//		return
//	}
//	if version != "HTTP/1.1"{
//					fmt.Println("проблема с версией")
//					return
//				}
//				//s.mu.RUnlock()
//				//
//				//handler := HandlerFunc(func(conn net.Conn) {
//				//	err = conn.Close()
//				//	log.Print(err, "это то что было внутри handler")
//				//})
//				//for _, _ = range s.handlers{
//				//
//				//	handl1 := s.handlers[path]
//				//	if handl1 != nil{
//				//		//
//				//		handler = handl1
//				//	}
//				//
//				//}
//				//
//				//s.mu.RUnlock()
//				//handler(conn)
//
//	if path != path{
//		conn.Close()
//	}
//	fmt.Println(path, "по этим путям зареган")
//
//	handler := func(conn net.Conn) {
//		err := conn.Close()
//		if err != nil {
//			log.Print(err)
//		}
//	} // какую он играет роль???
//	s.mu.RLock()
//	for i := 0; i < len(s.handlers); i++ {
//		fmt.Println(i, "это I")
//		handl1 := s.handlers[path] // about
//		// Nil != nil
//		if handl1 != nil{
//			handler = handl1
//		}
//
//		break
//	}
//	s.mu.RUnlock()
//	handler(conn)
//	// хмм тут только что мне только
//}


func (s *Server) handle(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	buffer := make([]byte, 4096)
	//for {
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			log.Printf("%v", buffer[:n])
			return
		}
		log.Print(err)
		return
	}

	data := buffer[:n]
	requestLineDelim := []byte{'\r', '\n'}
	requestLineEnd := bytes.Index(data, requestLineDelim)
	if requestLineEnd == -1 {
		return
	}

	requestLine := string(data[:requestLineEnd])
	parts := strings.Split(requestLine, " ")
	if len(parts) != 3 {
		return
	}

	path, version := parts[1], parts[2]

	if version != "HTTP/1.1" {
		return
	}
	if path != path{
		conn.Close()
	}
	fmt.Println(path, "по этим путям зареган")

	handler := func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	} // какую он играет роль???
	s.mu.RLock()
	for i := 0; i < len(s.handlers); i++ {
		fmt.Println(i, "это I")
		handl1 := s.handlers[path] // about
		// Nil != nil
		if handl1 != nil{
			handler = handl1
		}

		break
	}
	s.mu.RUnlock()
	handler(conn)
}

