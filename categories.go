package sumaregi

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

func (c *Client) GetCategories(ctx context.Context, opts GetCategoriesOpts) (*GetCategoriesResponse, error) {
	var result GetCategoriesResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	if len(opts.Fields) > 0 {
		fields := strings.Join(opts.Fields, ",")
		v.Set("fields", fields)
	}

	err = c.call(ctx, APIPathCategories, http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// func (c *Client) PostProducts(ctx context.Context, params PostProductsParams) (*PostProductsResponse, error) {
// 	var result PostProductsResponse

// 	err := c.call(ctx, APIPathProducts, http.MethodPost, nil, params, &result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }
