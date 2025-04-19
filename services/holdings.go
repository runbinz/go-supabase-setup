package services

import (
	"context"
	"fmt"

	"github.com/runbinz/dashboard/supabase"
)

// Package services provides functions to interact with the database.
// It includes functions to fetch user holdings and other data.
// This package is responsible for executing database queries and returning the results to the handlers.

// GetUserHoldings retrieves the holdings of a user from the database.
// It executes a SQL query to select asset details for a given user ID.
// The function returns a slice of Asset structs or an error if the query fails.
// This function is crucial for fetching the user's investment data, which is then returned to the client.

func GetUserHoldings(ctx context.Context, userID string) ([]Asset, error) {
	query := `
		SELECT
			h.symbol,
			h.name,
			h.quantity,
			h.value,
			0.0 AS change24h, -- placeholder
			0.0 AS allocation -- placeholder
		FROM holdings h
		JOIN portfolios p ON h.portfolio_id = p.id
		WHERE p.user_id = $1
	`

	rows, err := supabase.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var assets []Asset

	for rows.Next() {
		var asset Asset
		if err := rows.Scan(
			&asset.Symbol,
			&asset.Name,
			&asset.Quantity,
			&asset.Value,
			&asset.Change24h,
			&asset.Allocation,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		assets = append(assets, asset)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return assets, nil
}
