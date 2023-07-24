package middlewares

import (
	"context"
	"errors"
	"fmt"
	"interview/app/constants"
	redisclient "interview/app/providers/redis"
	"net/http"

	"github.com/google/uuid"
)

func RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId := r.Header.Get("X-USER-ID")
		requestId := uuid.New().String()

		err := storeRequestIdForUser(ctx, requestId, userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		errr := getTotalRequestForUserIdInCurrentTimeFrame(ctx, userId)
		if errr != nil {
			http.Error(w, "Rate Limit Reached", http.StatusInternalServerError)
		}

		next.ServeHTTP(w, r)
	})
}

func storeRequestIdForUser(ctx context.Context, requestId string, userId string) error {
	key := getKeyForRateLimiting(userId)
	client := redisclient.GetClient()
	err := client.SetDataForRequestRateLimiting(ctx, key, requestId, constants.TTL_FOR_RATE_LIMIT)
	if err != nil {
		return err
	}

	return nil
}

func getTotalRequestForUserIdInCurrentTimeFrame(ctx context.Context, userId string) error {
	client := redisclient.GetClient()
	key := getKeyForRateLimiting(userId)
	totalRequest, err := client.GetTotalRequestCountForRateLimiting(ctx, key)
	if err != nil {
		return err
	}

	if totalRequest < constants.MAX_REQUEST_ALLOWED_PER_MIN {
		return nil
	}

	return errors.New("Rate Limit Reached")
}

func getKeyForRateLimiting(userId string) string {
	return fmt.Sprintf("api:request_rate_limit:%s", userId)
}
