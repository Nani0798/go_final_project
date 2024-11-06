package scheduler

import (
	"fmt"
	"strconv"
	"time"

	"go_final_project/pkg/models"
	"go_final_project/storage/sqlite"
)

type Scheduler struct {
	Storage *sqlite.Storage
}

func NewScheduler(dataBase *sqlite.Storage) *Scheduler {
	return &Scheduler{
		Storage: dataBase,
	}
}

type TaskScheduler interface {
	SaveTask(*models.Task) (int64, error)
	GetTasks(string) ([]*models.Task, error)
	GetTaskByID(string) (*models.Task, error)
	UpdateTask(*models.Task) error
	MarkTaskCompleted(string) error
	DeleteTask(string) error
}

func (s *Scheduler) SaveTask(task *models.Task) (int64, error) {
	if err := check(task); err != nil {
		return 0, err
	}

	id, err := s.Storage.SaveTask(task)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Scheduler) GetTasks(search string) ([]*models.Task, error) {

	var sr_st sqlite.Search
	if search != "" {
		sr_st.Active = true
		date, err := time.Parse("02.01.2006", search)
		if err != nil {
			sr_st.Search = search
		} else {
			sr_st.Date = date.Format(d)
		}
	}

	tasks, err := s.Storage.GetTasks(&sr_st)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Scheduler) GetTaskByID(id string) (*models.Task, error) {
	task, err := s.Storage.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Scheduler) UpdateTask(task *models.Task) error {
	if task.ID == "" {
		return fmt.Errorf("id is empty")
	}

	if _, err := strconv.Atoi(task.ID); err != nil {
		return fmt.Errorf("id is not a number: %w", err)
	}

	if err := check(task); err != nil {
		return err
	}

	if err := s.Storage.UpdateTask(task); err != nil {
		return err
	}
	return nil
}

func check(task *models.Task) error {
	if task.Title == "" {
		return fmt.Errorf("empty title field")
	}

	var now = time.Now().Truncate(24 * time.Hour)
	var nextdate string

	if task.Date == "" {
		nextdate = now.Format(d)
	} else {
		date, err := time.Parse(d, task.Date)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		nextdate = date.Format(d)
		if date.Before(now) {
			if task.Repeat == "" {
				nextdate = now.Format(d)
			} else {
				nextdate, err = NextDate(now, task.Date, task.Repeat)
				if err != nil {
					return fmt.Errorf("%w", err)
				}
			}
		}
	}
	task.Date = nextdate

	return nil
}

func (s *Scheduler) MarkTaskCompleted(id string) error {
	var now = time.Now().Truncate(24 * time.Hour)
	task, err := s.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task.Repeat == "" {
		if err := s.Storage.DeleteTask(id); err != nil {
			return err
		}
	} else {
		nextdate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
		task.Date = nextdate

		err = s.Storage.UpdateTask(task)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Scheduler) DeleteTask(id string) error {
	if err := s.Storage.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
