package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	cache "github.com/psevdocoder/InMemoryCacheTTL"
	"log"
	"strings"
	"time"
)

type CacheResponse struct {
	Body       []byte
	StatusCode int
	Headers    map[string]string
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)                  // Записываем переданные байты в поле body
	return r.ResponseWriter.Write(b) // Вызываем метод Write исходного ResponseWriter
}

func CacheMiddleware(cache *cache.Cache, cacheTTL time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
			return
		}
		cacheKey := c.Request.URL.Path
		if cacheItem, found := cache.Get(cacheKey); found {
			cachedResponse := cacheItem.(CacheResponse)
			for key, value := range cachedResponse.Headers {
				c.Writer.Header().Set(key, value)
			}
			c.Writer.WriteHeader(cachedResponse.StatusCode)
			if _, err := c.Writer.Write(cachedResponse.Body); err != nil {
				log.Println(err)
				return
			}
			c.Abort()
			return
		}

		body := &bytes.Buffer{}
		// Создаем обертку для ResponseWriter, чтобы захватить ответ
		wrappedWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           body,
		}
		c.Writer = wrappedWriter

		c.Next()

		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 || c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
			cacheResponse := CacheResponse{
				Body:       body.Bytes(),
				StatusCode: c.Writer.Status(),
				Headers:    make(map[string]string),
			}
			for key, values := range c.Writer.Header() {
				cacheResponse.Headers[key] = strings.Join(values, "; ")
			}
			cache.Set(cacheKey, cacheResponse, cacheTTL)
		}
	}
}
