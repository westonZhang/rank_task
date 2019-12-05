package task_service_test

import (
	"rank-task/services/task_service"
	"rank-task/structs/models"
	"rank-task/structs/models/logics"
	"testing"
)

func TestSendCaptureTasksToChan(t *testing.T) {
	var tasks []*models.CaptureTask
	task1 := &models.CaptureTask{KeywordId: 1}
	task2 := &models.CaptureTask{KeywordId: 2}
	tasks = append(tasks, task1, task2)
	channel := make(chan *models.CaptureTask, logics.TASK_发送下载缓冲区大小)

	task_service.SendCaptureTasksToChan(tasks, channel, func(task *models.CaptureTask) {
		task.Status = logics.TASK_STATUS_查询中
	})

	verifyTask1 := <-channel
	verifyTask2 := <-channel

	if verifyTask1.KeywordId != verifyTask1.KeywordId {
		t.Error("keywordId not match")
	}
	if verifyTask2.KeywordId != verifyTask2.KeywordId {
		t.Error("keywordId not match")
	}
	if verifyTask1.Status != logics.TASK_STATUS_查询中 {
		t.Error("status not changed")
	}
	if verifyTask2.Status != logics.TASK_STATUS_查询中 {
		t.Error("status not changed")
	}
}
