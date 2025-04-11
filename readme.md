# FCLI - Flutter Komut Satırı Aracı

FCLI, Flutter geliştirme iş akışınızı hızlandırmak için tasarlanmış güçlü bir
komut satırı aracıdır. Bu araç, sık kullanılan Flutter komutlarını daha kısa ve
kullanışlı hale getirir.

## Özellikler

- Flutter projelerini temizleme ve bağımlılıkları yükleme (`fcli use`)
- Flutter uygulamasını çalıştırma (`fcli r`)
- Flutter doktor kontrolü (`fcli doctor`)
- Bağlı cihazları listeleme (`fcli devices`)
- Paket güncelleştirme (`fcli u`)
- Paket yükleme (`fcli i [paket_adı]`)
- Proje oluşturma (`fcli c [proje_adı]`)
- Dosya oluşturma (`fcli createfile [dosya_adı] [şablon_tipi]`)
- Uygulama derleme (`fcli build`)
- Sürüm bilgilerini görüntüleme (`fcli versions`)

## Kurulum

FCLI'yı kurmak için aşağıdaki adımları izleyin:

1. Projeyi derleyin:
   ```
   go build -o fcli
   ```

2. Sisteminize yükleyin:
   ```
   sudo cp fcli /usr/local/bin/
   ```

## Kullanım

FCLI, Flutter projenizin içindeyken varsayılan olarak mevcut dizini kullanır:
