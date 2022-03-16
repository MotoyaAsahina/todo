package router

import (
	"github.com/MotoyaAsahina/todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

type PostTaskRequest struct {
	GroupID     uuid.UUID `json:"group_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"due_date"`
	Tags        []string  `json:"tags"`
}

type TaskResponse struct {
	ID          uuid.UUID   `json:"id"`
	GroupID     uuid.UUID   `json:"group_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Done        bool        `json:"done"`
	DueDate     time.Time   `json:"due_date"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DoneAt      time.Time   `json:"done_at"`
	Tags        []uuid.UUID `json:"tags"`
}

func GetTasks(c echo.Context) error {
	tasks, err := model.GetTasks(c.Request().Context())
	if err != nil {
		return err
	}

	tagMap, err := model.GetTagMaps(c.Request().Context())
	if err != nil {
		return err
	}

	var response []TaskResponse
	for _, task := range tasks {
		response = append(response, TaskResponse{
			ID:          task.ID,
			GroupID:     task.GroupID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			DueDate:     task.DueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			DoneAt:      task.DoneAt,
			Tags:        tagMap[task.ID],
		})
	}

	return c.JSON(200, response)
}

func PostTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	task, err := model.PostTask(c.Request().Context(), &model.Task{
		ID:          uuid.New(),
		GroupID:     req.GroupID,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
		DueDate:     time.Now(), // TODO: parse
	})

	// tag map
	tagMaps := make([]*model.TagMap, 0)
	for _, tag := range req.Tags {
		tagMaps = append(tagMaps, &model.TagMap{
			ID:     uuid.New(),
			TaskID: task.ID,
			TagID:  uuid.MustParse(tag),
		})
	}
	if len(tagMaps) > 0 {
		err = model.PostTagMaps(c.Request().Context(), tagMaps)
		if err != nil {
			return err
		}
	}

	return c.JSON(200, task)
}

func PutTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	id := uuid.MustParse(c.Param("id"))

	registeredTagMaps, err := model.GetTagMapsByTaskID(c.Request().Context(), id)
	registeredTags := make([]uuid.UUID, 0)
	for _, tagMap := range registeredTagMaps {
		registeredTags = append(registeredTags, tagMap.TagID)
	}

	reqTags := make([]uuid.UUID, 0)
	for _, tag := range req.Tags {
		reqTags = append(reqTags, uuid.MustParse(tag))
	}

	newTags := make([]*model.TagMap, 0)
	for _, tagID := range reqTags {
		if !isRegisteredTag(tagID, registeredTags) {
			newTags = append(newTags, &model.TagMap{
				ID:     uuid.New(),
				TaskID: id,
				TagID:  tagID,
			})
		}
	}
	deletedTags := make([]*model.TagMap, 0)
	for _, tagMap := range registeredTagMaps {
		if !isRegisteredTag(tagMap.TagID, reqTags) {
			deletedTags = append(deletedTags, &model.TagMap{
				ID:     tagMap.ID,
				TaskID: id,
				TagID:  tagMap.TagID,
			})
		}
	}

	err = model.PutTask(c.Request().Context(), &model.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     time.Now(),
	})
	if err != nil {
		return err
	}

	if len(newTags) > 0 {
		err = model.PostTagMaps(c.Request().Context(), newTags)
		if err != nil {
			return err
		}
	}
	if len(deletedTags) > 0 {
		err = model.DeleteTagMaps(c.Request().Context(), deletedTags)
		if err != nil {
			return err
		}
	}

	return c.JSON(200, nil)
}

func isRegisteredTag(tagID uuid.UUID, registeredTags []uuid.UUID) bool {
	for _, registeredTag := range registeredTags {
		if tagID.String() == registeredTag.String() {
			return true
		}
	}
	return false
}

func DeleteTask(c echo.Context) error {
	return c.JSON(200, "DeleteTasks")
}

func PutTaskDone(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	err := model.PutTaskDone(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}

func PutTaskUndone(c echo.Context) error {
	return c.JSON(200, "PutTaskUndone")
}
