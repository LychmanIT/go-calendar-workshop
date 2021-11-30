package db

import (
	"calendarWorkshop/internal/util"
	"calendarWorkshop/models"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
)

type dbService struct {
	DBConn *pgxpool.Pool
}

//NewService returns new dbService instance.
func NewService(conn *pgxpool.Pool) Service { return &dbService{DBConn: conn} }

//GetUser returns user ID if user exists in the database and error if not.
func (d *dbService) GetUser(ctx context.Context, auth *models.Auth) (string, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return "", errors.New("error with db connection")
	}
	defer conn.Release()

	passHash := md5.New()
	passHash.Write([]byte(auth.Password))
	password := hex.EncodeToString(passHash.Sum(nil))

	row := conn.QueryRow(context.Background(),
		"SELECT id FROM users WHERE username = $1 AND password = $2", auth.Username, password)

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", errors.New("no user found")
	}

	return id, nil
}

//AddEvent receives event model and returns status 200 if it was successfully stored
//and status 500 if there is an error.
func (d *dbService) AddEvent(ctx context.Context, e *models.Event) (int, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO events (id, title, description, time, timezone, duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		e.ID, e.Title, e.Description, e.Time, e.Timezone, e.Duration)

	var id string
	err = row.Scan(&id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	for _, text := range e.Notes {
		rowNote := conn.QueryRow(context.Background(),
			"INSERT INTO event_notes (id, event_id, text) VALUES ($1, $2, $3) RETURNING id",
			uuid.New().String(), id, text)

		var noteID string
		err = rowNote.Scan(&noteID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

//AllEvents receives filters and returns list of events that stored in the database.
func (d *dbService) AllEvents(ctx context.Context, filters ...models.Filter) ([]models.Event, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var events []models.Event
	var query = "select * from events"

	if filters != nil {
		query = util.JoinQueryFilter(query, filters)
	}

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatal("error while executing query in AllEvents")
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset in AllEvents")
			return nil, err
		}

		id := values[0].(string)
		title := values[1].(string)
		description := values[2].(string)
		timeField := values[3].(string)
		timezone := values[4].(string)
		duration := values[5].(int32)

		notes, err := d.findEventNotesByID(ctx, id)
		if err != nil {
			return nil, err
		}

		events = append(events,
			models.Event{
				ID:          id,
				Title:       title,
				Description: description,
				Time:        timeField,
				Timezone:    timezone,
				Duration:    duration,
				Notes:       notes,
			})
	}

	return events, nil
}

//ShowEvent receives an ID of the event and returns that event from the database.
//WARNING! returns nil event instance if there is no event.
func (d *dbService) ShowEvent(ctx context.Context, ID string) (*models.Event, error) {
	event, err := d.findEventByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

//UpdateEvent receives an ID of the event we want to update and the new instance of event that we want to
//replace it with. Returns status 200 if everything is ok and 404 or 500 if there is problems.
func (d *dbService) UpdateEvent(ctx context.Context, ID string, e *models.Event) (int, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer conn.Release()

	row, err := conn.Query(context.Background(),
		"SELECT id FROM events WHERE id = $1", ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !util.IsRowsAffected(row) {
		return http.StatusNotFound, nil
	}

	_, err = d.DeleteEvent(ctx, ID)
	if err != nil {
		return http.StatusNotFound, nil
	}

	_ = conn.QueryRow(context.Background(),
		"INSERT INTO events (id, title, description, time, timezone, duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		ID, e.Title, e.Description, e.Time, e.Timezone, e.Duration)
	conn.Release()

	for _, text := range e.Notes {
		conn, err := d.DBConn.Acquire(ctx)
		rowNote := conn.QueryRow(context.Background(),
			"INSERT INTO event_notes (id, event_id, text) VALUES ($1, $2, $3) RETURNING id",
			uuid.New().String(), ID, text)

		var noteID string
		err = rowNote.Scan(&noteID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusOK, nil
}

//DeleteEvent receives an ID of the event we want to delete. Returns 200 if everything is ok and 404 or 500
//if there is problems.
func (d *dbService) DeleteEvent(ctx context.Context, ID string) (int, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer conn.Release()

	row, err := conn.Query(context.Background(),
		"SELECT id FROM events WHERE id = $1", ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !util.IsRowsAffected(row) {
		return http.StatusNotFound, nil
	}

	_, err = conn.Query(context.Background(), "DELETE FROM events WHERE id = $1", ID)
	if err != nil {
		log.Fatal("error while executing query in DeleteEvent")
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

//ServiceStatus returns service current health.
func (d *dbService) ServiceStatus(_ context.Context) (int, error) {
	return http.StatusOK, nil
}

//findEventByID receives an ID of needed event and finds it in the database
func (d *dbService) findEventByID(ctx context.Context, ID string) (*models.Event, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var result models.Event

	rows, err := conn.Query(context.Background(), "SELECT * FROM events WHERE id = $1", ID)
	if err != nil {
		log.Fatal("error while executing query in findEventByID")
	}

	for rows.Next() {

		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset in findEventByID")
		}
		result.ID = values[0].(string)
		result.Title = values[1].(string)
		result.Description = values[2].(string)
		result.Time = values[3].(string)
		result.Timezone = values[4].(string)
		result.Duration = values[5].(int32)
	}

	result.Notes, err = d.findEventNotesByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

//findEventNotesByID receives an ID of needed event and finds its notes in the database
func (d *dbService) findEventNotesByID(ctx context.Context, ID string) ([]string, error) {
	conn, err := d.DBConn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var result []string

	notesRows, err := conn.Query(context.Background(), "SELECT * FROM event_notes WHERE event_id = $1", ID)
	if err != nil {
		log.Fatal("error while executing query in findEventNotesByID")
	}

	for notesRows.Next() {
		notesValues, err := notesRows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset in findEventNotesByID")
		}
		result = append(result, notesValues[2].(string))
	}
	return result, nil
}
