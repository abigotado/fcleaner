# fcleaner

Утилита для очистки Flutter проекта и переустановки зависимостей.

## Что делает

1. Выполняет `flutter clean`
2. Выполняет `flutter pub get`
3. Удаляет `ios/Podfile.lock`
4. Удаляет `ios/Pods`
5. Выполняет `pod install --repo-update` в папке `ios`

## Способы запуска

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