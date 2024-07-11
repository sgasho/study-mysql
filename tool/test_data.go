package tool

import (
	"context"
	"fmt"
	"strings"

	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
)

func prepareTestData(txx *sqlx.Tx) error {
	ctx := context.Background()
	f := faker.New()
	userIDs, err := createUsers(ctx, txx, f, 1000)
	if err != nil {
		return txx.Rollback()
	}

	// create videos

	// create genres

	// create video_genres

	// create user_interests

	// create user_no_interests

	// create_watch_history
}

func createUsers(ctx context.Context, txx *sqlx.Tx, f faker.Faker, amount int) ([]string, error) {
	values := strings.Repeat(", (?, ?, ?)", amount-1)
	q := fmt.Sprintf(`INSERT INTO user (user_id, username, email) VALUES (?, ?, ?) %s`, values)

	args := make([]string, 0, amount*4)
	userIDs := make([]string, amount)
	for i := 0; i < amount; i++ {
		userID := f.UUID().V4()
		userName := fmt.Sprintf("%s%s", f.Person().FirstName(), f.Person().LastName())
		args = append(args, userID, userName, fmt.Sprintf("%s@email.com", userName))
		userIDs[i] = userID
	}

	if _, err := txx.ExecContext(ctx, q, args); err != nil {
		return nil, err
	}

	return userIDs, nil
}

func createVideos(f faker.Faker, amount int) []string {
	return nil
}

func createGenres(f faker.Faker, amount int) []string {
	return nil
}

func createVideoGenres(videoIDs, genreIDs []string) {

}

func createUserInterests(userIDs, genreIDs []string) {

}

func createUserNoInterests(userIDs, genreIDs []string) {

}

func createWatchHistories(userIDs, videoIDs []string) {

}
