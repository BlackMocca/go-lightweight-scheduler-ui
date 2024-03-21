package api

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"

// 	"github.com/Blackmocca/wedding-ui/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/maxence-charriere/go-app/v9/pkg/app"
// )

// type weddingAPIClient struct {
// 	instance
// }

// var (
// 	WeddingAPI weddingAPIClient
// )

// func init() {
// 	if app.IsClient {
// 		WeddingAPI = weddingAPIClient{
// 			instance: instance{
// 				host:    app.Getenv("API_URL"),
// 				timeout: 30,
// 			},
// 		}
// 	}
// }

// func (w weddingAPIClient) Fetch(ctx context.Context) ([]*models.Celebrate, error) {
// 	client := w.getClient()
// 	uri := "/api/celebrate"

// 	resp, err := client.R().Get(uri)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var respM map[string]interface{}
// 	if err := json.Unmarshal(resp.Body(), &respM); err != nil {
// 		return nil, err
// 	}
// 	if resp.StatusCode() != 200 {
// 		return nil, errors.New(respM["message"].(string))
// 	}

// 	var ptrs = []*models.Celebrate{}
// 	bu, _ := json.Marshal(respM["celebrates"])
// 	json.Unmarshal(bu, &ptrs)
// 	return ptrs, nil
// }
// func (w weddingAPIClient) Create(ctx context.Context, celText string, celFrom string) error {
// 	client := w.getClient()
// 	uri := "/api/celebrate"
// 	body := map[string]interface{}{
// 		"celebrate_text": celText,
// 		"celebrate_from": celFrom,
// 	}

// 	resp, err := client.R().SetHeader("Content-Type", echo.MIMEApplicationJSONCharsetUTF8).SetBody(body).Post(uri)
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode() != 200 {
// 		var respM map[string]interface{}
// 		if err := json.Unmarshal(resp.Body(), &respM); err != nil {
// 			return err
// 		}

// 		return errors.New(respM["message"].(string))
// 	}

// 	return nil
// }
