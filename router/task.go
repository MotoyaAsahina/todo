package router

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/MotoyaAsahina/todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
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

func (h *Handlers) GetTasks(c echo.Context) error {
	tasks, err := h.Repo.GetTasks(c.Request().Context())
	if err != nil {
		return err
	}

	tagMap, err := h.Repo.GetTagMaps(c.Request().Context())
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

func (h *Handlers) PostTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	dueDate, err := parseDueDate(req.DueDate)
	if err != nil {
		return err
	}

	task, err := h.Repo.PostTask(c.Request().Context(), &model.Task{
		ID:          uuid.New(),
		GroupID:     req.GroupID,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
		DueDate:     dueDate,
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
		err = h.Repo.PostTagMaps(c.Request().Context(), tagMaps)
		if err != nil {
			return err
		}
	}

	// notification
	err = h.scheduleNotification(c.Request().Context(), task)
	if err != nil {
		return err
	}

	return c.JSON(200, task)
}

func (h *Handlers) PutTask(c echo.Context) error {
	var req PostTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	id := uuid.MustParse(c.Param("id"))

	registeredTagMaps, err := h.Repo.GetTagMapsByTaskID(c.Request().Context(), id)
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

	dueDate, err := parseDueDate(req.DueDate)
	if err != nil {
		return err
	}

	task, err := h.Repo.PutTask(c.Request().Context(), &model.Task{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
	})
	if err != nil {
		return err
	}

	if len(newTags) > 0 {
		err = h.Repo.PostTagMaps(c.Request().Context(), newTags)
		if err != nil {
			return err
		}
	}
	if len(deletedTags) > 0 {
		err = h.Repo.DeleteTagMaps(c.Request().Context(), deletedTags)
		if err != nil {
			return err
		}
	}

	// notification
	err = h.Repo.DeleteNotifications(c.Request().Context(), id)
	if err != nil {
		return err
	}
	err = h.scheduleNotification(c.Request().Context(), task)
	if err != nil {
		return err
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

func (h *Handlers) DeleteTask(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))

	err := h.Repo.DeleteTagMapsByTaskID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	err = h.Repo.DeleteTask(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(200, nil)
}

func (h *Handlers) PutTaskDone(c echo.Context) error {
	id := uuid.MustParse(c.Param("id"))
	err := h.Repo.PutTaskDone(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}

func (h *Handlers) PutTaskUndone(c echo.Context) error {
	return c.JSON(200, "PutTaskUndone")
}

func (h *Handlers) scheduleNotification(ctx context.Context, task *model.Task) error {
	notificationTimes, tags := getNotificationTimesFromDescription(task.Description, task.DueDate)
	for i, t := range notificationTimes {
		fromNow := t.Sub(time.Now())
		if fromNow <= 0 {
			continue
		}

		alreadyTimeExists, err := h.Repo.SetNotification(ctx, task.ID, t, tags[i])
		if err != nil {
			return err
		}
		if !alreadyTimeExists {

			timer := time.NewTimer(fromNow)
			go func() {
				<-timer.C
				h.notify()
			}()
		}
	}

	return nil
}

func (h *Handlers) notify() {
	notifications, err := h.Repo.GetLatestNotifications(context.Background())
	if err != nil {
		postMessage("Error: "+err.Error(), 0)
		return
	}

	groups, err := h.Repo.GetGroups(context.Background())
	if err != nil {
		postMessage("Error: "+err.Error(), 0)
		return
	}

	findGroup := func(id uuid.UUID) *model.Group {
		for _, group := range groups {
			if group.Id == id {
				return group
			}
		}
		return nil
	}

	type formedTask struct {
		Title           string
		Description     string
		GroupName       string
		Tags            []string
		DueDate         time.Time
		NotificationTag string
	}
	tasks := make([]formedTask, 0)

	for _, notification := range notifications {
		task, err := h.Repo.GetTask(context.Background(), notification.TaskID)
		if err != nil {
			postMessage("Error: "+err.Error(), 0)
			return
		}

		tags, err := h.Repo.GetTagNamesByTaskID(context.Background(), task.ID)
		if err != nil {
			postMessage("Error: "+err.Error(), 0)
			return
		}

		if !task.Done {
			tasks = append(tasks, formedTask{
				Title:           task.Title,
				Description:     task.Description,
				GroupName:       findGroup(task.GroupID).Name,
				DueDate:         task.DueDate,
				Tags:            tags,
				NotificationTag: notification.Tag,
			})
		}
	}

	// sort tasks by due date
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].DueDate.Before(tasks[j].DueDate)
	})

	if len(tasks) == 0 {
		return
	}

	// build message
	message := "Tasks to do:\n"
	for _, task := range tasks {
		if len(task.Tags) > 0 {
			task.GroupName += "," + strings.Join(task.Tags, ",")
		}
		message += fmt.Sprintf(
			"%s (%s): %s (Remaining %s)\n",
			task.Title,
			task.GroupName,
			task.DueDate.Format("2006/01/02 15:04"),
			task.NotificationTag,
		)
	}
	postMessage(message, 0)
}

func postMessage(message string, retryCount int) {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	_, _, err := api.PostMessage(os.Getenv("SLACK_CHANNEL_ID"), slack.MsgOptionText(message, false))
	if err != nil {
		if retryCount < 5 {
			time.Sleep(5 * time.Second)
			postMessage(message, retryCount+1)
		} else {
			fmt.Printf("[slack post] %s %s\n", time.Now().Format("2006/01/02 15:04"), err.Error())
		}
	} else if retryCount > 0 {

	}
}

func getNotificationTimesFromDescription(description string, dueDate time.Time) ([]time.Time, []string) {
	r := regexp.MustCompile(`!notice\[(.*)]`)
	m := r.FindStringSubmatch(description)

	res := ""
	if len(m) > 0 {
		res = m[1]
	}

	if dueDate.Minute() == 59 {
		dueDate = dueDate.Add(time.Minute)
	}

	notificationTimes := make([]time.Time, 0)

	resSlice := strings.Split(res, ",")
	for i, v := range resSlice {
		s := strings.TrimSpace(v)
		resSlice[i] = s
		if s == "" {
			continue
		}

		norm := 0.0
		switch s[len(s)-1:] {
		case "d":
			norm = 24 * 60
		case "h":
			norm = 60
		case "m":
			norm = 1
		default:
			continue
		}

		numFloat, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			continue
		}
		min := time.Duration(int(numFloat*norm)) * time.Minute

		notificationTimes = append(notificationTimes, dueDate.Add(-min))
	}

	return notificationTimes, resSlice
}

func (h *Handlers) ResetNotifications() {
	notificationTimes, err := h.Repo.GetValidNotificationTimes(context.Background())
	if err != nil {
		fmt.Printf("[ERROR: reset notifications] %s\n", err.Error())
		return
	}

	for _, t := range notificationTimes {
		fromNow := t.Sub(time.Now())
		if fromNow <= 0 {
			// already passed
			err = h.Repo.SetNotificationTimeNoticed(context.Background(), t)
			continue
		}

		timer := time.NewTimer(fromNow)
		go func() {
			<-timer.C
			h.notify()
		}()
	}
}
