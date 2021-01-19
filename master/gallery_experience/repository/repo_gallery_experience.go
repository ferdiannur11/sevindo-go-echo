package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/master/gallery_experience"
	"github.com/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type galleryexperienceRepository struct {
	Conn *sql.DB
}



// NewuserRepository will create an object that represent the article.repository interface
func NewGalleryExperienceRepository(Conn *sql.DB) gallery_experience.Repository {
	return &galleryexperienceRepository{Conn}
}
func (m *galleryexperienceRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.GalleryExperience, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.GalleryExperience, 0)
	for rows.Next() {
		t := new(models.GalleryExperience)
		err = rows.Scan(
			&t.Id,
			&t.CreatedBy,
			&t.CreatedDate,
			&t.ModifiedBy,
			&t.ModifiedDate,
			&t.DeletedBy,
			&t.DeletedDate,
			&t.IsDeleted,
			&t.IsActive,
			&t.ExperienceName,
			&t.ExperienceDesc,
			&t.ExperiencePicture,
			&t.Longitude,
			&t.Latitude,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
func (m *galleryexperienceRepository) GetByID(ctx context.Context, id string) (res *models.GalleryExperience, err error) {
	query := `SELECT * FROM gallery_experiences WHERE `

	if id != "" {
		query = query + ` id = '` + id + `' `
	}

	list, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (m *galleryexperienceRepository) Update(ctx context.Context, a *models.GalleryExperience) error {
	query := `UPDATE gallery_experiences set modified_by=?, modified_date=?,
	 experience_name=?, experience_desc=?, experience_picture=?, longitude=?, latitude=? WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, a.ModifiedBy, time.Now(), a.ExperienceName, a.ExperienceDesc,
		a.ExperiencePicture, a.Longitude, a.Latitude, a.Id)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}

func (m *galleryexperienceRepository) Delete(ctx context.Context, id string, deleted_by string) error {
	query := `UPDATE gallery_experiences SET deleted_by=? , deleted_date=? , is_deleted=? , is_active=? WHERE id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, deleted_by, time.Now(), 1, 0, id)
	if err != nil {
		return err
	}

	//lastID, err := res.RowsAffected()
	if err != nil {
		return err
	}

	//a.Id = lastID
	return nil
}


func (m *galleryexperienceRepository) Insert(ctx context.Context, a *models.GalleryExperience) error {
	query := `INSERT gallery_experiences SET id=? , created_by=? , created_date=? , modified_by=?, modified_date=? , deleted_by=? , deleted_date=? , is_deleted=? , is_active=? ,
	 experience_name=?, experience_desc=?,experience_picture=?, longitude=?, latitude=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, a.Id, a.CreatedBy, time.Now(), nil, nil, nil, nil, 0, 1,  a.ExperienceName, a.ExperienceDesc, a.ExperiencePicture, a.Longitude, a.Latitude)
	if err != nil {
		return err
	}

	//lastID, err := res.RowsAffected()
	if err != nil {
		return err
	}

	//a.Id = lastID
	return nil
}

func (m *galleryexperienceRepository) Count(ctx context.Context) (int, error) {
	query := `SELECT count(*) AS count FROM gallery_experiences WHERE is_deleted = 0 and is_active = 1`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	count, err := checkCount(rows)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return count, nil
}

func checkCount(rows *sql.Rows) (count int, err error) {
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (m *galleryexperienceRepository) List(ctx context.Context, limit, offset int) ([]*models.GalleryExperience, error) {
	query := `SELECT * FROM gallery_experiences WHERE is_deleted = 0 and is_active = 1 `

	query = query + ` LIMIT ? OFFSET ?`
	list, err := m.fetch(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return list, nil
}
