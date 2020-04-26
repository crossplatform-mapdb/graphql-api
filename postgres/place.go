package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/zackartz/go-graphql-api/models"
)

type PlacesRepo struct {
	DB *pg.DB
}

func (p *PlacesRepo) GetPlaces() ([]*models.Place, error) {
	var places []*models.Place
	err := p.DB.Model(&places).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return places, nil
}

func (p *PlacesRepo) CreatePlace(place *models.Place) (*models.Place, error) {
	_, err := p.DB.Model(place).Returning("*").Insert()

	return place, err
}

func (p *PlacesRepo) GetById(id string) (*models.Place, error) {
	var place models.Place
	err := p.DB.Model(&place).Where(" id = ?", id).First()
	return &place, err
}

func (p *PlacesRepo) Update(place *models.Place) (*models.Place, error) {
	_, err := p.DB.Model(place).Where("id = ?", place.ID).Update()
	return place, err
}

func (p *PlacesRepo) Delete(place *models.Place) error {
	_, err := p.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}

func (p *PlacesRepo) GetPlacesForUser(user *models.User) ([]*models.Place, error) {
	var places []*models.Place
	err := p.DB.Model(&places).Where("user_id = ?", user.ID).Order("id").Select()
	return places, err
}
