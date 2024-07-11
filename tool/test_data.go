package tool

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
)

func prepareTestData(txx *sqlx.Tx) error {
	ctx := context.Background()
	f := faker.New()
	userIDs, err := createUsers(ctx, txx, &f, 1000)
	if err != nil {
		return txx.Rollback()
	}

	videoIDs := make([]string, 0)
	for i := 1; i <= 10; i++ {
		ids, err := createVideos(ctx, txx, &f, 1000)
		if err != nil {
			return txx.Rollback()
		}
		videoIDs = append(videoIDs, ids...)
	}

	genreIDs, err := createGenres(ctx, txx, &f, 100)
	if err != nil {
		return txx.Rollback()
	}

	if err := createVideoGenres(ctx, txx, videoIDs, genreIDs); err != nil {
		return txx.Rollback()
	}

	// create user_interests

	// create user_no_interests

	// create_watch_history
}

func createUsers(ctx context.Context, txx *sqlx.Tx, f *faker.Faker, amount int) ([]string, error) {
	values := strings.Repeat(", (?, ?, ?)", amount-1)
	q := fmt.Sprintf(`INSERT INTO user (user_id, username, email) VALUES (?, ?, ?) %s`, values)

	args := make([]string, 0, amount*3)
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

func createVideos(ctx context.Context, txx *sqlx.Tx, f *faker.Faker, amount int) ([]string, error) {
	values := strings.Repeat(", (?, ?, ?)", amount-1)
	q := fmt.Sprintf(`INSERT INTO video (video_id, title, description) VALUES (?, ?, ?) %s`, values)

	args := make([]string, 0, amount*3)
	videoIDs := make([]string, amount)
	for i := 0; i < amount; i++ {
		videoID := f.UUID().V4()
		args = append(args, videoID, f.Lorem().Sentence(10), f.Lorem().Paragraph(5))
		videoIDs[i] = videoID
	}

	if _, err := txx.ExecContext(ctx, q, args); err != nil {
		return nil, err
	}

	return videoIDs, nil
}

func createGenres(ctx context.Context, txx *sqlx.Tx, f *faker.Faker, amount int) ([]string, error) {
	values := strings.Repeat(", (?, ?)", amount-1)
	q := fmt.Sprintf(`INSERT INTO genre (genre_id, genre_name) VALUES (?, ?) %s`, values)

	args := make([]string, 0, amount*2)
	genreIDs := make([]string, amount)
	for i := 0; i < amount; i++ {
		genreID := f.UUID().V4()
		args = append(args, genreID, f.Music().Genre()) // use music genre names instead of video ones 'cause no video genres in faker
		genreIDs[i] = genreID
	}

	if _, err := txx.ExecContext(ctx, q, args); err != nil {
		return nil, err
	}

	return genreIDs, nil
}

func createVideoGenres(ctx context.Context, txx *sqlx.Tx, videoIDs, genreIDs []string) error {
	args := make([]string, 0)
	valueAmount := 0
	for _, videoID := range videoIDs {
		genreCount := rand.Intn(10)
		for i := 0; i < genreCount; i++ {
			args = append(args, videoID, genreIDs[rand.Intn(len(genreIDs))])
		}
	}

	values := strings.Repeat(", (?, ?)", valueAmount-1)
	q := fmt.Sprintf(`INSERT INTO video_genre (video_id, genre_id) VALUES (?, ?) %s`, values)

	if _, err := txx.ExecContext(ctx, q, args); err != nil {
		return err
	}

	return nil
}

func createUserInterests(userIDs, genreIDs []string) {

}

func createUserNoInterests(userIDs, genreIDs []string) {

}

func createWatchHistories(userIDs, videoIDs []string) {

}
