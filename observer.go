package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Observer interface, представляющий наблюдателя
type Observer interface {
	Notify()
}

// Subject interface, представляющий объект, за которым наблюдает наблюдатель
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConfigReader представляет объект, который будет уведомлять о
// необходимости перезагрузки и обновлении конфигурационного файла
type ConfigReader struct {
	observers []Observer
}

// AddObserver добавляет наблюдателя в список наблюдателей
func (c *ConfigReader) AddObserver(observer Observer) {
	c.observers = append(c.observers, observer)
}

// RemoveObserver удаляет наблюдателя из списка наблюдателей
func (c *ConfigReader) RemoveObserver(observer Observer) {
	for i, obs := range c.observers {
		if obs == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers уведомляет всех наблюдателей об изменениях в конфигурационном файле
func (c *ConfigReader) NotifyObservers() {
	for _, observer := range c.observers {
		observer.Notify()
	}
}

// ConfigurableObject представляет объект, который нуждается в обновлении конфигурационного файла
type ConfigurableObject struct {
	// ...
}

// Notify оповещает объект о необходимости перезагрузки и обновлении конфигурационного файла
func (c *ConfigurableObject) Notify() {
	// Обновление конфигурационного файла
	// Перезагрузка объекта
}

func watchConfig(server *http.Server) {
	var lastConfigChange time.Time

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if e.Op == fsnotify.Write {
			fileStat, _ := os.Stat(viper.ConfigFileUsed())
			if !lastConfigChange.Equal(fileStat.ModTime()) {
				lastConfigChange = fileStat.ModTime()
				// Остановка текущего сервера
				fmt.Println("Stopping server...")
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				if err := server.Shutdown(ctx); err != nil {
					fmt.Printf("Error stopping server: %s\n", err)
				}

				// Ожидание остановки сервера
				for {
					if err := server.ListenAndServe(); err != http.ErrServerClosed {
						fmt.Printf("Error waiting for server to stop: %s\n", err)
						time.Sleep(1 * time.Second)
					} else {
						break
					}
				}

				// Ожидание полного завершения текущего сервера
				time.Sleep(1 * time.Second)

				// Запуск нового сервера с обновленным портом
				newServer := &http.Server{
					Addr: fmt.Sprintf(":%d", viper.GetInt("server.port")),
					Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("Hello, World!\n"))
					}),
				}
				log.Printf("Starting server on port %d...", viper.GetInt("server.port"))
				go func() {
					if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("Error starting server: %s\n", err)
					}
				}()
				// Обновление ссылки на сервер
				server = newServer
			}
		}
	})
}

func main() {
	// Инициализация конфигурации
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	// Запуск сервера на указанном порту
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	// Обработка изменений в файле конфигурации
	watchConfig(server)

	// Ожидание сигнала для завершения работы программы
	log.Println("Press Ctrl+C to exit")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Остановка сервера
	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
