// +build integration

package system

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	v2client "github.com/donggangcj/goharbor-client/v3/apiv2/generate/api/client"
	"github.com/donggangcj/goharbor-client/v3/apiv2/generate/legacyapi/client"
	integrationtest "github.com/donggangcj/goharbor-client/v3/apiv2/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	model "github.com/donggangcj/goharbor-client/v3/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

// TestAPISystemGcScheduleNew tests the creation of a new GC schedule
// and resets it to the default schedule afterwards
func TestAPISystemGcScheduleNew(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	gcSchedule, err := c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	assert.NotNil(t, gcSchedule)

	defer c.ResetSystemGarbageCollection(ctx)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType)
}

// TestAPISystemGcScheduleUpdate tests the update of an existing GC schedule,
// asserting the updated schedule cron matches the given values
func TestAPISystemGcScheduleUpdate(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	cron2 := "* * */1 * *"
	scheduleType2 := "Daily"
	err = c.UpdateSystemGarbageCollection(ctx, &model.AdminJobScheduleObj{
		Cron: cron2,
		Type: scheduleType2,
	})

	gcSchedule, err := c.GetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	assert.Equal(t, gcSchedule.Schedule.Cron, cron2)
	assert.Equal(t, gcSchedule.Schedule.Type, scheduleType2)

	err = c.ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)
}

// TestAPISystemGcScheduleReset tests the reset of an existing GC schedule
func TestAPISystemGcScheduleReset(t *testing.T) {
	cron := "0 * * * *"
	scheduleType := "Hourly"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	_, err := c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)

	_, err = c.NewSystemGarbageCollection(ctx, cron, scheduleType)
	require.NoError(t, err)

	err = c.ResetSystemGarbageCollection(ctx)
	require.NoError(t, err)

	_, err = c.GetSystemGarbageCollection(ctx)
	require.IsType(t, &ErrSystemGcUndefined{}, err)
}

func TestAPIHealth(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	_, err := c.Health(ctx)
	require.NoError(t, err)
}
