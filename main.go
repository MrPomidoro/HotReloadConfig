package main

func main() {

}

//type ConfigWatcher struct {
//	log *logrus.Logger
//	s   *server.Server
//}
//
//func (cw *ConfigWatcher) Update() {
//	newServer, err := cw.s.ServerUpdate()
//	if err != nil {
//		cw.log.Fatalf("Ошибка обновления сервера: %v", err)
//	}
//
//	time.Sleep(5 * time.Second)
//
//	newServer.Start()
//
//	cw.log.Info("Сервер успешно запущен")
//	cw.s = newServer
//	cw.log.Info("Объек сервера переопределен")
//}
//
//type Publisher interface {
//	Notify()
//	AddObserver(ConfigWatcher)
//}
//
//type ConfigPublisher struct {
//	observers []ConfigWatcher
//}
//
//func (cp *ConfigPublisher) AddObserver(o ConfigWatcher) {
//	cp.observers = append(cp.observers, o)
//}
//
//func (cp *ConfigPublisher) Notify() {
//	for _, observer := range cp.observers {
//		observer.Update()
//	}
//}

//func main() {
//// создание сервера и обозревателя настроек
//// TODO Переделать viper чтобы переопределять структуру сервера
//var log *logrus.Logger
//log = logger.New()
//log.Info("Начало работы")
//
//viper.SetConfigName("config")
//viper.AddConfigPath("../")
//log.Info("Чтение конфигурационного файла...")
//if err := viper.ReadInConfig(); err != nil {
//	log.Fatalf("Ошибка чтения файла конфигурации: %s \n", err)
//}
//
//cfg := &server.CServer{Port: viper.GetInt("server.port"), Domain: viper.GetString("server.domain")}
//log.Info("Конфигурационный файл успешно прочитан, конфиг сервера ", cfg)
//
//newServer := server.NewServer(log)
//
//newServer.Start()
//
//cw := ConfigWatcher{
//	log: log,
//	s:   newServer,
//}
//
//publisher := &ConfigPublisher{}
//p := Publisher(publisher)
//p.AddObserver(cw)
//
//var lastConfigChange time.Time
//// запуск мониторинга изменений файла конфигурации
//viper.WatchConfig()
//viper.OnConfigChange(func(e fsnotify.Event) {
//	if e.Op == fsnotify.Write {
//		fileStat, _ := os.Stat(viper.ConfigFileUsed())
//		if !lastConfigChange.Equal(fileStat.ModTime()) {
//			lastConfigChange = fileStat.ModTime()
//			log.Info("Файл конфигурации изменился:", e.Name)
//			p.Notify()
//		}
//	}
//})
//
//// Ожидание сигнала для завершения работы программы
//log.Println("Press Ctrl+C to exit")
//signalChan := make(chan os.Signal, 1)
//signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
//<-signalChan
//
//// Остановка сервера
//if err := newServer.Stop(); err != nil {
//	log.Errorf("Server shutdown error: %v", err)
//}

//}
