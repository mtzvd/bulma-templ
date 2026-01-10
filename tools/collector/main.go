package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const outputFileName = "all_code.txt"

func main() {
	// Получаем текущую директорию
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения текущей директории: %v\n", err)
		os.Exit(1)
	}

	// Получаем имя исполняемого файла
	exeName := filepath.Base(os.Args[0])

	// Загружаем правила .gitignore
	ignorePatterns := loadGitignore(currentDir)

	// Собираем список всех файлов
	var files []string
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускаем директории
		if info.IsDir() {
			return nil
		}

		// Получаем относительный путь
		relPath, err := filepath.Rel(currentDir, path)
		if err != nil {
			return err
		}

		// Нормализуем путь для сравнения (используем прямые слэши)
		relPath = filepath.ToSlash(relPath)

		// Пропускаем сам исполняемый файл
		if filepath.Base(path) == exeName || filepath.Base(path) == exeName+".exe" {
			return nil
		}

		// Пропускаем выходной файл
		if filepath.Base(path) == outputFileName {
			return nil
		}

		// Пропускаем .git директорию
		if strings.Contains(relPath, ".git/") || strings.HasPrefix(relPath, ".git/") {
			return nil
		}

		// Проверяем по правилам .gitignore
		if shouldIgnore(relPath, ignorePatterns) {
			return nil
		}

		files = append(files, relPath)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка обхода директорий: %v\n", err)
		os.Exit(1)
	}

	// Создаём выходной файл
	outFile, err := os.Create(filepath.Join(currentDir, outputFileName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка создания выходного файла: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Записываем заголовок
	writer.WriteString("СПИСОК ФАЙЛОВ\n")
	writer.WriteString(strings.Repeat("=", 80) + "\n\n")

	// Записываем список файлов
	for _, file := range files {
		writer.WriteString(file + "\n")
	}

	writer.WriteString("\n" + strings.Repeat("=", 80) + "\n\n")

	// Записываем содержимое каждого файла
	for _, relPath := range files {
		fullPath := filepath.Join(currentDir, filepath.FromSlash(relPath))

		writer.WriteString(fmt.Sprintf("ФАЙЛ: %s\n", relPath))
		writer.WriteString(strings.Repeat("-", 80) + "\n")

		// Читаем содержимое файла
		content, err := os.ReadFile(fullPath)
		if err != nil {
			writer.WriteString(fmt.Sprintf("[Ошибка чтения файла: %v]\n", err))
		} else {
			writer.Write(content)
			// Добавляем перенос строки если файл не заканчивается на \n
			if len(content) > 0 && content[len(content)-1] != '\n' {
				writer.WriteString("\n")
			}
		}

		writer.WriteString(strings.Repeat("=", 80) + "\n\n")
	}

	fmt.Printf("Успешно создан файл %s с %d файлами\n", outputFileName, len(files))
}

// loadGitignore загружает правила из .gitignore
func loadGitignore(dir string) []string {
	gitignorePath := filepath.Join(dir, ".gitignore")
	file, err := os.Open(gitignorePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Пропускаем пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		patterns = append(patterns, line)
	}

	return patterns
}

// shouldIgnore проверяет, должен ли файл быть проигнорирован
func shouldIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		// Убираем начальный слэш если есть
		pattern = strings.TrimPrefix(pattern, "/")

		// Простая проверка на совпадение
		if matchPattern(path, pattern) {
			return true
		}
	}
	return false
}

// matchPattern выполняет простое сопоставление паттернов
func matchPattern(path, pattern string) bool {
	// Если паттерн заканчивается на /, это директория
	if strings.HasSuffix(pattern, "/") {
		pattern = strings.TrimSuffix(pattern, "/")
		// Проверяем, начинается ли путь с этой директории
		return strings.HasPrefix(path, pattern+"/") || path == pattern
	}

	// Если паттерн содержит /, проверяем полный путь
	if strings.Contains(pattern, "/") {
		matched, _ := filepath.Match(pattern, path)
		return matched
	}

	// Иначе проверяем имя файла
	fileName := filepath.Base(path)
	matched, _ := filepath.Match(pattern, fileName)
	if matched {
		return true
	}

	// Также проверяем, есть ли совпадение в любой части пути
	pathParts := strings.Split(path, "/")
	for _, part := range pathParts {
		matched, _ := filepath.Match(pattern, part)
		if matched {
			return true
		}
	}

	return false
}
