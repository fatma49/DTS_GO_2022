package taskmodel

import (
	"database/sql"

	"github.com/fatma49/DTS_GO_2022/config"
	"github.com/fatma49/DTS_GO_2022/entities"
)

type TaskModel struct {
	db *sql.DB
}

func New() *TaskModel {
	db, err := config.DBConnection()

	if err != nil {
		panic(err)
	}
	return &TaskModel{db: db}
}

func (m *TaskModel) FindAll(task *[]entities.Task) error {
	rows, err := m.db.Query("select * from task")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Task
		rows.Scan(
			&data.id,
			&data.detail,
			&data.pegawai,
			&data.deadline)

		*task = append(*task, data)
	}
	return nil
}

func (m *TaskModel) Create(task *entities.Task) error {
	result, err := m.db.Exec("insert into task (detail, pegawai, deadline) values(?,?,?)",
		task.detail, task.pegawai, task.deadline)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	task.id = lastInsertId
	return nil
}

func (m *TaskModel) Find(id int64, task *entities.Task) error {
	return m.db.QueryRow("select * from task where id = ?", id).Scan(
		&task.id,
		&task.detail,
		&task.pegawai,
		&task.deadline)
}

func (m *TaskModel) Update(task entities.Task) error {

	_, err := m.db.Exec("update task set detail = ?, pegawai = ?, deadline = ? where id = ?",
		task.detail, task.pegawai, task.deadline, task.id)

	if err != nil {
		return err
	}

	return nil
}

func (m *TaskModel) Delete(id int64) error {
	_, err := m.db.Exec("delete from task where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
