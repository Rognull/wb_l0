package db

import (
	"l0/internal/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	log "github.com/sirupsen/logrus"
)

type OrderStorage struct {
	 dataBase     *pg.DB
}



func NewOrderStorage(dBase *pg.DB) *OrderStorage {
	storage := new(OrderStorage)
	storage.dataBase = dBase
	return storage
}

func (db *OrderStorage) AddOrder(order model.Order) error {

	if _, err := db.dataBase.Model(&order).Insert(); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func(db *OrderStorage) MigrateDb() (error, []model.Order){
	models := []interface{}{
		(*model.Order)(nil),
	}

	var err error

	for _, model := range models {
		op := orm.CreateTableOptions{IfNotExists: true}
		err = db.dataBase.Model(model).CreateTable(&op)
		if err != nil {
			log.Fatalln(err," in cashe create migrate")
			return err, nil
		}
	}

	orders := make([]model.Order, 0)
	err =  db.dataBase.Model(&orders).Select()

	if err != nil {
		log.Fatalln(err, "in cashe select migrate")
		return err, orders
	}
	 
	return err, orders
}






// func (storage *VoteStorage) AddOrder(vote model.Vote) error {
// 	query := "INSERT INTO Vote (peer_id, candidate_id, election_id) VALUES ($1, $2, $3)"

// 	_, err := storage.databasePool.Exec(context.Background(),query, vote.Id_voter , vote.Id_candidate , vote.Id_election) //транзакция не нужна, у нас только один запрос

// 	if err != nil {
// 		log.Errorln(err)
// 		return err
// 	}
	
// 	return nil
// }




// func convertJoinedQueryToCar(input userCar) models.Car { //сворачиваем плоскую структуру в нашу рабочую модель
// 	return models.Vote{
// 		Id_voter:Id_voter,
// 		Id_Vote_one:input.Id_Vote_one,
// 		Id_Vote_two:input.Id_Vote_two,
// 	}
// }