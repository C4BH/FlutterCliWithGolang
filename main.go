package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)
func checkError(err error) {
	if err != nil {
		fmt.Println("Hata:", err)
	}
}

func runCommand(command string, dir string, args ...string) error {
	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var rootCmd = &cobra.Command{
	Use:   "fcli",
	Short: "Flutter CLI flutter komutlarımı kullanabilmek için bir CLI",
}

var projectPath string

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectPath, "path", "p", "", "Komutların çalıştırılacağı proje dizini (boş bırakılırsa mevcut dizin kullanılır)")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(doctorCMD)
	rootCmd.AddCommand(devicesCMD)
	rootCmd.AddCommand(createfileCMD)
	rootCmd.AddCommand(buildCMD)
	rootCmd.AddCommand(installCMD)
	rootCmd.AddCommand(upgradeCMD)
	rootCmd.AddCommand(createCMD)
	rootCmd.AddCommand(versionCMD)
}

var runCmd = &cobra.Command{
	Use:   "r",
	Short: "Flutter uygulamasını çalıştırır",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Flutter uygulaması çalıştırılıyor...")
		err := runCommand("flutter", projectPath, "run")
		checkError(err)
	},
}
var doctorCMD = &cobra.Command{
	Use:   "doctor",
	Short: "Flutter uygulamasının sağlık durumunu kontrol eder",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "doctor")
		checkError(err)
	},
}

var devicesCMD = &cobra.Command{
	Use:   "devices",
	Short: "Flutter uygulamasının cihazlarını listeler",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "devices")
		if err != nil {
			fmt.Println("Hata:", err)
		}
	},
}
var upgradeCMD = &cobra.Command{
	Use:   "u",
	Short: "Flutter uygulamasını günceller",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath,"pub", "upgrade")
		checkError(err)
	},
}
var installCMD = &cobra.Command{
	Use:   "i",
	Short: "Flutter uygulamasına bağımlılık ekler",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "pub", "add", args[0])
		checkError(err)
	},
}

var createfileCMD = &cobra.Command{
	Use:   "createfile",
	Short: "Flutter dosyası oluşturur",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Hata: createfile için proje ismi ve template tipi girilmelidir.")
			return
		  }
		err := runCommand("flutter", projectPath, "create", args[0], args[1])
		checkError(err)
	},
}

var buildCMD = &cobra.Command{
	Use:   "build",
	Short: "Flutter uygulamasını derler",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "build")
		checkError(err)
	},
}

var createCMD = &cobra.Command{
	Use:   "c",
	Short: "Flutter projesini oluşturur",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "create", args[0])
		checkError(err)
	},
}
var versionCMD = &cobra.Command{
	Use:   "versions",
	Short: "FCLI ve Flutter sürümünü gösterir",
	Run: func(cmd *cobra.Command, args []string) {
		err := runCommand("flutter", projectPath, "--version")
		checkError(err)
	},
}
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Flutter bağımlılıklarını yükler ve temizlik yapar",
	Run: func(cmd *cobra.Command, args []string) {

		// Sonra clean çalıştır
		fmt.Println("Flutter projesi temizleniyor...")
		err := runCommand("flutter", projectPath, "clean")
		if err != nil {
			fmt.Println("Clean hatası:", err)
			return
		}
		// Önce pub get çalıştır
		fmt.Println("Flutter bağımlılıkları yükleniyor...")
		err = runCommand("flutter", projectPath, "pub", "get")
		if err != nil {
			fmt.Println("Pub get hatası:", err)
			return
		}

		fmt.Println("İşlemler başarıyla tamamlandı!")
	},
}

func main() {
	// Eğer proje yolu belirtilmediyse, mevcut dizini kullan
	if projectPath == "" {
		currentDir, err := os.Getwd()
		if err == nil {
			projectPath = currentDir
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
