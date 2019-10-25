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

func TestCheckHealthz(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	resp := httpGet(fix.handler, "/healthz")
	if got, want := resp.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}
}

func TestHandler_RegisterNewCategories(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cat := modelfixture.Category()
	fix.appService.On("RegisterNewCategory", app.RegisterNewCategoryCommand{
		Name: cat.Name(),
	}).Return(cat, nil)

	resp := httpPost(fix.handler, "/categories", map[string]interface{}{
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
	fix := setupFixture(t)
	defer fix.tearDown()

	cats := modelfixture.Categories(3)
	fix.appService.On("RetrieveCategories").Return(cats, nil)

	res := httpGet(fix.handler, "/categories")
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
	fix := setupFixture(t)
	defer fix.tearDown()

	cat := modelfixture.Category()
	fix.appService.On("RetrieveCategoryByID", app.RetrieveCategoryByIDCommand{
		ID: cat.ID(),
	}).Return(cat, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/categories/%s", cat.ID()))
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

func TestHandler_CategoryByID_notFound(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cat := modelfixture.Category()
	fix.appService.On("RetrieveCategoryByID", app.RetrieveCategoryByIDCommand{
		ID: cat.ID(),
	}).Return(nil, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/categories/%s", cat.ID()))
	if got, want := res.StatusCode, nethttp.StatusNotFound; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}
}

func TestHandler_RegisterNewStore(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	fix.appService.On("RegisterNewStore", app.RegisterNewStoreCommand{
		Name: str.Name(),
	}).Return(str, nil)

	resp := httpPost(fix.handler, "/stores", map[string]interface{}{
		"name": str.Name(),
	})
	if got, want := resp.StatusCode, nethttp.StatusCreated; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	if got, want := resp.Header.Get("Location"), fmt.Sprintf("/stores/%s", str.ID()); got != want {
		t.Errorf("Location got: %q, want: %q", got, want)
	}
}

func TestHandler_Stores(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	strs := modelfixture.Stores(3)
	fix.appService.On("RetrieveStores").Return(strs, nil)

	res := httpGet(fix.handler, "/stores")
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	var out []storePayload
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatal("err:", err)
	}

	if got, want := len(out), len(strs); got != want {
		t.Fatalf("Length got: %d, want: %d", got, want)
	}

	for i, str := range strs {
		row := out[i]
		if got, want := row.ID, str.ID(); got != want {
			t.Errorf("ID got: %q. want: %q, index: %d", got, want, i)
		}

		if got, want := row.Name, str.Name(); got != want {
			t.Errorf("Name got: %q. want: %q, index: %d", got, want, i)
		}
	}
}

func TestHandler_StoreByID(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	fix.appService.On("RetrieveStoreByID", app.RetrieveStoreByIDCommand{
		ID: str.ID(),
	}).Return(str, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/stores/%s", str.ID()))
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	var out storePayload
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatal("err:", err)
	}

	if got, want := out.ID, str.ID(); got != want {
		t.Errorf("ID got: %q. want: %q", got, want)
	}

	if got, want := out.Name, str.Name(); got != want {
		t.Errorf("Name got: %q. want: %q", got, want)
	}
}

func TestHandler_StoreByID_notFound(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	fix.appService.On("RetrieveStoreByID", app.RetrieveStoreByIDCommand{
		ID: str.ID(),
	}).Return(nil, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/stores/%s", str.ID()))
	if got, want := res.StatusCode, nethttp.StatusNotFound; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}
}

func TestHandler_OfferNewProduct(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	prd := modelfixture.ProductOfStore(str)

	fix.appService.On("OfferNewProduct", app.OfferNewProductCommand{
		StoreID:     prd.StoreID(),
		CategoryID:  prd.CategoryID(),
		Name:        prd.Name(),
		Price:       prd.Price(),
		Description: prd.Description(),
		Quantity:    prd.Quantity(),
	}).Return(prd, nil)

	resp := httpPost(fix.handler, fmt.Sprintf("/stores/%s/products", str.ID()), map[string]interface{}{
		"categoryId":  prd.CategoryID(),
		"name":        prd.Name(),
		"price":       prd.Price().String(),
		"description": prd.Description(),
		"quantity":    prd.Quantity(),
	})
	if got, want := resp.StatusCode, nethttp.StatusCreated; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	if got, want := resp.Header.Get("Location"), fmt.Sprintf("/products/%s", prd.ID()); got != want {
		t.Errorf("Location got: %q, want: %q", got, want)
	}
}

func TestHandler_ProductByID(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	prd := modelfixture.ProductOfStore(str)
	fix.appService.On("RetrieveProductByID", app.RetrieveProductByIDCommand{
		ID: prd.ID(),
	}).Return(prd, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/products/%s", prd.ID()))
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("StatusCode got: %d, want: %d", got, want)
	}

	var out productPayload
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		t.Fatal("err:", err)
	}

	if got, want := out.ID, prd.ID(); got != want {
		t.Errorf("ID got: %q. want: %q", got, want)
	}

	if got, want := out.StoreID, prd.StoreID(); got != want {
		t.Errorf("StoreID got: %q. want: %q", got, want)
	}

	if got, want := out.CategoryID, prd.CategoryID(); got != want {
		t.Errorf("CategoryID got: %q. want: %q", got, want)
	}

	if got, want := out.Name, prd.Name(); got != want {
		t.Errorf("Name got: %q. want: %q", got, want)
	}

	if got, want := out.Price, prd.Price().String(); got != want {
		t.Errorf("Price got: %q. want: %q", got, want)
	}

	if got, want := out.Description, prd.Description(); got != want {
		t.Errorf("Description got: %q. want: %q", got, want)
	}

	if got, want := out.Quantity, prd.Quantity(); got != want {
		t.Errorf("Quantity got: %d. want: %d", got, want)
	}
}

type categoryPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type storePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type productPayload struct {
	ID          string `json:"id"`
	StoreID     string `json:"storeId"`
	CategoryID  string `json:"categoryId"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type testFixture struct {
	t          *testing.T
	appService *mocks.AppService
	handler    nethttp.Handler
}

func setupFixture(t *testing.T) *testFixture {
	appService := new(mocks.AppService)

	handler, err := http.NewHandler(appService)
	if err != nil {
		t.Fatal(fmt.Errorf("Fail to create handler: %w", err))
	}

	return &testFixture{
		t:          t,
		appService: appService,
		handler:    handler,
	}
}

func (fix *testFixture) tearDown() {
	mock.AssertExpectationsForObjects(fix.t, fix.appService)
}
