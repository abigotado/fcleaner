# fcleaner

Утилита для очистки Flutter проекта и переустановки зависимостей.

## Что делает

1. Выполняет `flutter clean`
2. Выполняет `flutter pub get`
3. Удаляет `ios/Podfile.lock`
4. Удаляет `ios/Pods`
5. Выполняет `pod install --repo-update` в папке `ios`

## 📥 Скачать готовый бинарный файл

**Рекомендуемый способ** - скачать готовый бинарный файл для вашей системы:

### [⬇️ Скачать последнюю версию](../../releases/latest)

Выберите файл для вашей операционной системы:
- **macOS Intel**: `fcleaner-macos-intel`
- **macOS Apple Silicon (M1/M2)**: `fcleaner-macos-arm64`  
- **Linux x64**: `fcleaner-linux-x64`
- **Windows x64**: `fcleaner-windows-x64.exe`

### Установка скачанного файла

1. Скачайте нужный файл
2. Сделайте его исполняемым (macOS/Linux):
   ```bash
   chmod +x fcleaner-*
   ```
3. Переместите в системную папку для глобального использования:
   ```bash
   # macOS/Linux
   sudo mv fcleaner-* /usr/local/bin/fcleaner
   
   # Или добавьте в PATH
   mv fcleaner-* ~/bin/fcleaner  # если ~/bin в PATH
   ```

Теперь можно запускать из любой директории:
```bash
fcleaner
fcleaner -fvm
```

## Способы запуска (для разработчиков)

Если у вас установлен Go, можете запускать из исходного кода:

### 1. Запуск исходного файла

```bash
go run src/main.go
```

С использованием FVM:
```bash
go run src/main.go -fvm
```

### 2. Сборка и запуск бинарного файла

Собрать бинарный файл:
```bash
go build -o fcleaner src/main.go
```

Запустить:
```bash
./fcleaner
```

С использованием FVM:
```bash
./fcleaner -fvm
```

### 3. Установка в систему

Собрать и установить глобально:
```bash
go build -o fcleaner src/main.go
sudo mv fcleaner /usr/local/bin/
```

Теперь можно запускать из любой директории:
```bash
fcleaner
```

С использованием FVM:
```bash
fcleaner -fvm
```

## Опции

- `-fvm` - использовать FVM для выполнения команд Flutter (вместо `flutter` будет использоваться `fvm flutter`)

## Требования

- Go 1.16+
- Flutter SDK (или FVM с установленной версией Flutter)
- CocoaPods (для iOS проектов) 