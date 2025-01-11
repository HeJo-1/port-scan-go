package main

import (
	"fmt"
	"net"
	"time"
)

// Port türü belirleme fonksiyonu
func getPortType(port int) string {
	switch port {
	case 80:
		return "HTTP"
	case 443:
		return "HTTPS"
	case 21:
		return "FTP"
	case 22:
		return "SSH"
	case 25:
		return "SMTP"
	case 110:
		return "POP3"
	case 3306:
		return "MySQL"
	default:
		return "Bilinmiyor"
	}
}

func main() {
	var ip string
	var startPort, endPort int
	var bekle int

	// Kullanıcıdan hedef IP ve port aralığı alınır
	fmt.Println("===================================")
	fmt.Println("       Port Tarayıcı Uygulaması     ")
	fmt.Println("===================================")
	fmt.Print("\nHedef IP adresini girin: ")
	fmt.Scan(&ip)
	fmt.Print("Başlangıç portunu girin: ")
	fmt.Scan(&startPort)
	fmt.Print("Bitiş portunu girin: ")
	fmt.Scan(&endPort)
	println("Tarama Başlıyor Lütfen Bekleyin...")

	// Açık portlar için başlık
	fmt.Println("\n===================================")
	fmt.Println("            AÇIK PORTLAR           ")
	fmt.Println("===================================")
	openPorts := false // Açık port var mı kontrolü için
	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", ip, port)
		startTime := time.Now() // Bağlantı başlama zamanı
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err != nil {
			// Port kapalı olduğunda hiçbir şey yazdırılmaz
		} else {
			// Port açık olduğunda tür ve süreyi yazdır
			duration := time.Since(startTime)                                                         // Bağlantı süresi
			portType := getPortType(port)                                                             // Port türü
			fmt.Printf("\033[32mPort %d açık - Tür: %s, Süre: %v\033[0m\n", port, portType, duration) // Yeşil renk
			conn.Close()
			openPorts = true
		}
	}
	if !openPorts {
		fmt.Println("\033[31mAçık port bulunamadı.\033[0m") // Eğer açık port yoksa uyarı mesajı
	}

	// Kapalı portlar için başlık
	fmt.Println("\n===================================")
	fmt.Println("            KAPALI PORTLAR         ")
	fmt.Println("===================================")
	closedPorts := false // Kapalı port var mı kontrolü için
	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", address, 1)

		if err != nil {
			// Kapalı portu ekrana yazdır
			fmt.Printf("\033[31mPort %d kapalı\033[0m\n", port) // Kırmızı renk
			closedPorts = true
		} else {
			conn.Close() // Bağlantı açık olursa kapat
		}
	}
	if !closedPorts {
		fmt.Println("\033[32mTüm portlar açık.\033[0m") // Eğer tüm portlar açık ise mesaj
	}

	// Bitirici mesaj
	fmt.Println("\n===================================")
	fmt.Println("         Tarama Tamamlandı!        ")
	fmt.Println("===================================")
	fmt.Scan(&bekle)
}
