package client

import "github.com/boichique/movie-reviews/contracts"

func (c *Client) CreateStar(req *contracts.AuthenticatedRequest[*contracts.CreateStarRequest]) (*contracts.StarDetails, error) {
	var star *contracts.StarDetails

	_, err := c.client.R().
		SetAuthToken(req.AccessToken).
		SetBody(req.Request).
		SetResult(&star).
		Post(c.path("/api/stars"))

	return star, err
}

func (c *Client) GetStars(req *contracts.GetStarsPaginatedRequest) (*contracts.PaginatedResponse[contracts.Star], error) {
	var stars contracts.PaginatedResponse[contracts.Star]

	_, err := c.client.R().
		SetResult(&stars).
		SetQueryParams(req.ToQueryParams()).
		Get(c.path("/api/stars"))

	return &stars, err
}

func (c *Client) GetStarByID(starID int) (*contracts.StarDetails, error) {
	var star contracts.StarDetails

	_, err := c.client.R().
		SetResult(&star).
		Get(c.path("/api/stars/%d", starID))

	return &star, err
}

func (c *Client) UpdateStar(req *contracts.AuthenticatedRequest[*contracts.UpdateStarRequest]) error {
	_, err := c.client.R().
		SetAuthToken(req.AccessToken).
		SetBody(req.Request).
		Put(c.path("/api/stars/%d", req.Request.StarID))

	return err
}

func (c *Client) DeleteStar(req *contracts.AuthenticatedRequest[*contracts.DeleteStarRequest]) error {
	_, err := c.client.R().
		SetAuthToken(req.AccessToken).
		SetBody(req.Request).
		Delete(c.path("/api/stars/%d", req.Request.StarID))

	return err
}
