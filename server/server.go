package server

//func main() {
//	builder := NewBuilder()
//	director := NewDirector(builder)
//	server, err := director.Construct(
//		":8080",
//		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello, World!\n")) }),
//		time.Second*10,
//		time.Second*15,
//		time.Second*3,
//	)
//
//	if err = server.ListenAndServe(); err != nil {
//		log.Fatal("err", err)
//	}
//
//}

//func NewServer(log *logrus.Logger) *Server {
//	cs := &CServer{}
//	cs.setPort()
//	cs.setDomain()
//	return &Server{
//		log: log,
//		cs:  cs,
//	}
//}
//
//func (s *Server) Start() {
//	server := &http.Server{
//		Addr: fmt.Sprintf("%s:%d", s.cs.Domain, s.cs.Port),
//		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			w.Write([]byte("Hello, World!\n"))
//		}),
//	}
//
//	go func() {
//		err := server.ListenAndServe()
//		if err != nil {
//			s.log.Warningf("Ошибка запуска сервера: %v", err)
//			return
//		}
//	}()
//}
//
//func (s *Server) Stop() error {
//	s.log.Info("Остановка сервера...")
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	err := s.server.Shutdown(ctx)
//	if err != nil {
//		return err
//	}
//
//	// Ожидание остановки сервера
//	for {
//		s.log.Info("Ожидание остановки сервера...")
//		if err = s.server.ListenAndServe(); err != http.ErrServerClosed {
//			s.log.Errorf("Ошибка ожидания остановки сервера: %s\n", err)
//			time.Sleep(1 * time.Second)
//			return err
//		} else {
//			s.log.Info("Сервер успешно остановлен")
//			break
//		}
//	}
//	return nil
//}
//
//func (s *Server) ServerUpdate() (*Server, error) {
//	s.log.Info("Перезагрузка сервера...")
//
//	err := s.Stop()
//	if err != nil {
//		return nil, err
//	}
//
//	return NewServer(s.log), nil
//}
//
//func (cs *CServer) setPort() {
//	cs.Port = viper.GetInt("server.port")
//}
//
//func (cs *CServer) setDomain() {
//	cs.Domain = viper.GetString("server.domain")
//}
