package controller

import (
	"net/http"
	"strings"
	"time"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/db/types"
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

	var yearMonth string
	if req.YearMonth != nil {
		yearMonth = req.YearMonth.Format("200601")
	} else {
		yearMonth = time.Now().Format("200601")
	}

	job := faktory.NewJob(slugToJobName(req.LiveHouseSlug), yearMonth)
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

type getLiveEventsRequest struct {
	LiveHouseSlug *string    `form:"live_house_slug"`
	YearMonth     *time.Time `form:"year_month" time_format:"200601"`
}

func (controller *Controller) GetLiveEvents(ctx *gin.Context) {
	var req getLiveEventsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.LiveHouseSlug == nil {
		liveEvents, err := controller.store.GetAllLiveEvents(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, map[string][]db.GetAllLiveEventsRow{"liveEvents": liveEvents})
	} else {
		liveEvents, err := controller.store.GetAllLiveEventsByLiveHouseSlug(ctx, types.NewNullString(*req.LiveHouseSlug))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, map[string][]db.GetAllLiveEventsByLiveHouseSlugRow{"liveEvents": liveEvents})
	}
}
