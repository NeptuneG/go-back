package controller

import (
	"net/http"
	"strings"
	"time"

	faktory "github.com/contribsys/faktory/client"
	"github.com/gin-gonic/gin"
)

type createScrapeLiveEventsJobRequest struct {
	LiveHouseSlug string     `form:"live_house_slug" binding:"required"`
	YearMonth     *time.Time `form:"year_month" time_format:"200601"`
}

func (controller *Controller) CreateScrapeLiveEventsJob(ctx *gin.Context) {
	var req createScrapeLiveEventsJobRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	client, err := faktory.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var year_month string
	if req.YearMonth != nil {
		year_month = req.YearMonth.Format("200601")
	} else {
		year_month = time.Now().Format("200601")
	}

	job := faktory.NewJob(slugToJobName(req.LiveHouseSlug), year_month)
	if err = client.Push(job); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, map[string]string{"jobId": job.Jid})
}

func slugToJobName(slug string) string {
	words := strings.Split(slug, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return "Scrape" + strings.Join(words, "") + "Job"
}
