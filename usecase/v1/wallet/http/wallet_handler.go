package ucv1wallethttp

import (
	"log"
	"time"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity/std"
	"github.com/labstack/echo/v4"
)

func (ds *domainService) Tetsing(c echo.Context) error {
	var apiResp *std.APIResponse

	// Do mutex locking with redis
	mutexR := ds.rSync.NewMutex("testing")
	// lock the mutex of, it will fail if the mutex with the same name already exists
	// and defered unlock mutex of remu
	if err := mutexR.Lock(); err != nil {
		log.Printf("[RemuLock] - Error: %s", err.Error())

		apiResp = std.APIResponseError(std.HTTPStatusCode(std.TOOMANYREQUESTS), err)
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	}

	now := time.Now()
	nowFormat := now.Format("2006-01-02 15:04:05")
	republish := map[string]interface{}{
		"last_update": nowFormat,
		"App":         "update wallet common",
	}

	if ok, err := mutexR.Unlock(); !ok || err != nil {
		log.Printf("[RemuUnlock] - Error: %s", err.Error())
		apiResp = std.APIResponseError(std.StatusBadRequest, err)
		return c.JSON(int(apiResp.StatusCode), apiResp.Body)
	} else {
		log.Printf("[%v][RemuUnlock] - Success!!", ok)
	}

	apiResp = std.APIResponseSuccess(republish)
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
