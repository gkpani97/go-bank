package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T){
	arg := CreateAccountParams{
		Owner: "Tom",
		Balance: 100,
		Currency: "USD",
	}

	account, err := TestQueries.CreateAccount(context.Background(), arg)

	fmt.Printf("%+v", account)

	
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}