package routes

type TaskListRequestData struct {
	PageNum    int    `form:"page_num"`
	PageSize   int    `form:"page_size"`
	NodeId     string `form:"node_id"`
	SpiderId   string `form:"spider_id"`
	ScheduleId string `form:"schedule_id"`
	Database     string `form:"status"`
}