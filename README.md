# Product_APP

Bu proje, PostgreSQL veritabanında ürünlerin depolandığı ve Go Echo framework'ünü kullanarak bu ürün bilgilerini güncelleyen bir API sunucusudur.

## Kurulum

1. Öncelikle, projeyi bilgisayarınıza klonlayın:

   ```bash
   git clone https://github.com/rahimeT/Product_APP.git
   ```

2. Gerekli bağımlılıkları yüklemek için projenin kök dizinine gidin ve `go mod tidy` komutunu çalıştırın:

   ```bash
   cd Product_APP
   go mod tidy
   ```

3. PostgreSQL veritabanınızı hazırlayın ve bağlantı bilgilerini test_db.sh dosyasına

4. Proje dizinindeyken aşağıdaki komutu çalıştırarak projeyi başlatın:

   ```bash
   go run .
   ```

## Kullanım

- API'ye erişmek için `localhost:port` adresine HTTP istekleri gönderebilirsiniz. Proje varsayılan olarak 8080 portunda çalışır.

- Örnek bir GET isteği:

  ```
  GET /products
  ```

- Örnek bir POST isteği:

  ```
  POST /products
  ```

  ```json
  {
    "name": "Ürün adı",
    "price": 10.99,
    "description": "Ürün açıklaması"
  }
  ```

- Diğer API rotaları ve istekleri için lütfen kodu inceleyin veya API dokümantasyonunu kontrol edin.

## Katkıda Bulunma

1. Bu depoyu klonlayın:

   ```bash
   git clone https://github.com/rahimeT/Product_APP.git
   ```

2. Değişikliklerinizi yapın ve yeni bir dal oluşturun:

   ```bash
   git checkout -b feature/newFeature
   ```

3. Değişikliklerinizi yapın, commit'leyin ve dalınıza itin:

   ```bash
   git add .
   git commit -m "Yaptığınız değişikliklerin açıklayıcı bir mesajı"
   git push origin feature/newFeature
   ```

4. Bir birleştirme isteği (pull request) açın ve değişikliklerinizi değerlendirin.
