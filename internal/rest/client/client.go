package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"gitlab16.skiftrade.kz/libs-go/logger"
	"gitlab16.skiftrade.kz/templates1/go/internal"
	"gitlab16.skiftrade.kz/templates1/go/internal/rest/client/models"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type client struct {
	config     models.Config
	httpClient *http.Client
}

func NewClient(config models.Config) internal.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConnsPerHost = t.MaxIdleConns
	return &client{
		config: config,
		httpClient: &http.Client{
			Transport: otelhttp.NewTransport(
				t,
				otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
					return fmt.Sprintf("HTTP %s %s", r.Method, r.URL.Path)
				})),
			Timeout: config.Timeout,
		},
	}
}

func (c *client) PostingsToCancel(ctx context.Context, token string, req models.PostingsToCancelReq) (response models.PostingsToCancelResp, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.config.Host,
		Path:   models.PathPostingsToCancel,
	}

	// Добавляем query параметр parcelType=MVP
	q := u.Query()
	if req.ParcelType != "" {
		q.Set("parcelType", req.ParcelType)
	}
	if req.IsTerminalCancel != nil {
		q.Set("isTerminalCancel", strconv.FormatBool(*req.IsTerminalCancel))
	}
	u.RawQuery = q.Encode()

	// Формируем HTTP запрос с контекстом
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return response, err
	}

	// Устанавливаем заголовки
	request.Header.Set("Authorization", "Bearer "+token)

	// Отправляем запрос
	resp, err := c.httpClient.Do(request)
	if err != nil {
		slog.ErrorContext(ctx, "failed to do the request", logger.ErrorAttr(err), logger.InputAttr(request))
		return response, err
	}
	defer func(body io.ReadCloser) {
		if err := resp.Body.Close(); err != nil {
			slog.ErrorContext(ctx, "failed to close response body",
				logger.ErrorAttr(err),
				slog.String("url", u.String()))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.ErrorContext(ctx, "failed to read body", logger.InputAttr(request),
			slog.Int("status_code", resp.StatusCode))
		return response, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		slog.ErrorContext(ctx, "unexpected status code", logger.InputAttr(request),
			slog.Int("status_code", resp.StatusCode), slog.String("body", string(body)))
		return response, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Декодируем JSON ответ в структуру
	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (c *client) PostingsCancelResponse(ctx context.Context, token string, req models.PostingsCancelResponseReq) (response models.PostingsCancelResponseResp, err error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.config.Host,
		Path:   models.PathPostingsCancelResponse,
	}

	requestBody, err := json.Marshal(req.Body)
	if err != nil {
		return response, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Формируем HTTP запрос с контекстом
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(requestBody))
	if err != nil {
		return response, err
	}

	// Устанавливаем заголовки
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	// Отправляем запрос
	resp, err := c.httpClient.Do(request)
	if err != nil {
		slog.ErrorContext(ctx, "failed to do the request", logger.ErrorAttr(err), logger.InputAttr(request))
		return response, err
	}
	defer func(body io.ReadCloser) {
		if err := resp.Body.Close(); err != nil {
			slog.ErrorContext(ctx, "failed to close response body",
				logger.ErrorAttr(err),
				slog.String("url", u.String()))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.ErrorContext(ctx, "failed to read body", logger.InputAttr(request),
			slog.Int("status_code", resp.StatusCode))
		return response, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		slog.ErrorContext(ctx, "unexpected status code", logger.InputAttr(request),
			slog.Int("status_code", resp.StatusCode), slog.String("body", string(body)))
		return response, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Декодируем JSON ответ в структуру
	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}
