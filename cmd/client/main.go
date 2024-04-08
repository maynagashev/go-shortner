package main

import (
	"bufio"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	flags := parseFlags()
	var contextTimeout = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	data, _ := readURLDataFromConsole()
	body := sendRequestToShortner(ctx, flags.GetServerURL(), data)
	slog.Info(string(body))
}

func sendRequestToShortner(ctx context.Context, endpoint string, data url.Values) []byte {
	var err error

	client := &http.Client{}

	// Тело запроса должно быть источником потокового чтения io.Reader
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	// В заголовках запроса указываем кодировку.
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Отправляем запрос и получаем ответ.
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// выводим код ответа
	slog.Info("Статус-код", "status_code", response.Status)

	// читаем поток из тела ответа
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func readURLDataFromConsole() (url.Values, error) {
	// контейнер данных для запроса
	data := url.Values{}

	slog.Info("Введите длинный URL")
	// открываем потоковое чтение из консоли
	reader := bufio.NewReader(os.Stdin)
	// читаем строку из консоли
	long, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	long = strings.TrimSuffix(long, "\n")
	data.Set("url", long)

	return data, err
}
