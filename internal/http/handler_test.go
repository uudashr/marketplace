package http_test

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"testing"

	modelfixture "github.com/uudashr/marketplace/internal/fixture"

	"github.com/uudashr/marketplace/internal/app"

	"github.com/uudashr/marketplace/internal/http"

	"github.com/stretchr/testify/mock"
	"github.com/uudashr/marketplace/internal/http/mocks"
)

func TestHandler_RegisterNewCategories(t *testing.T) {
	appService := new(mocks.AppService)
	defer mock.AssertExpectationsForObjects(t, appService)

	handler, err := http.NewHandler(appService)
	if err != nil {
		t.Fatal("err:", err)
	}

	cat := modelfixture.Category()
	appService.On("RegisterNewCategory", app.RegisterNewCategoryCommand{
		Name: cat.Name(),
	}).Return(cat, nil)

	resp := httpPost(handler, "/categories", map[string]interface{}{
		"name": cat.Name(),
	})
	if got, want := resp.StatusCode, nethttp.StatusCreated; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	if got, want := resp.Header.Get("Location"), fmt.Sprintf("/categories/%s", cat.ID()); got != want {
		t.Errorf("Location got: %q, want: %q", got, want)
	}
}

func TestHandler_Categories(t *testing.T) {
	appService := new(mocks.AppService)
	defer mock.AssertExpectationsForObjects(t, appService)

	handler, err := http.NewHandler(appService)
	if err != nil {
		t.Fatal("err:", err)
	}

	cats := modelfixture.Categories(3)
	appService.On("RetrieveCategories").Return(cats, nil)

	res := httpGet(handler, "/categories")
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	var out []categoryPayload
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatal("err:", err)
	}

	if got, want := len(out), len(cats); got != want {
		t.Fatalf("Length got: %d, want: %d", got, want)
	}

	for i, cat := range cats {
		row := out[i]
		if got, want := row.ID, cat.ID(); got != want {
			t.Errorf("ID got: %q. want: %q, index: %d", got, want, i)
		}

		if got, want := row.Name, cat.Name(); got != want {
			t.Errorf("Name got: %q. want: %q, index: %d", got, want, i)
		}
	}
}

func TestHandler_CategoryByID(t *testing.T) {
	appService := new(mocks.AppService)
	defer mock.AssertExpectationsForObjects(t, appService)

	handler, err := http.NewHandler(appService)
	if err != nil {
		t.Fatal("err:", err)
	}

	cat := modelfixture.Category()
	appService.On("RetrieveCategoryByID", app.RetrieveCategoryByIDCommand{
		ID: cat.ID(),
	}).Return(cat, nil)

	res := httpGet(handler, fmt.Sprintf("/categories/%s", cat.ID()))
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	var out categoryPayload
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatal("err:", err)
	}

	if got, want := out.ID, cat.ID(); got != want {
		t.Errorf("ID got: %q. want: %q", got, want)
	}

	if got, want := out.Name, cat.Name(); got != want {
		t.Errorf("Name got: %q. want: %q", got, want)
	}
}

type categoryPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}