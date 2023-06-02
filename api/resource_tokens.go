package api

import (
	"context"
)

func (c *Client) RevokeLimitedAccessToken(ctx context.Context, id string) error {
	query := `
		mutation($input:DeleteLimitedAccessTokenInput!) {
			deleteLimitedAccessToken(input: $input) {
				token
			}
		}
	`
	req := c.NewRequest(query)

	req.Var("input", map[string]interface{}{
		"id": id,
	})

	_, err := c.RunWithContext(ctx, req)
	if err != nil {
		return err
	}

	return nil
}