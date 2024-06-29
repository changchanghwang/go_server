package accountService

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"with.framework/domain/account"
	infrastructure "with.framework/infrastructure/account"
)

type ServiceTestSuite struct {
	suite.Suite
	accountService AccountService
}

func (suite *ServiceTestSuite) SetupTest() {
	account := &account.Account{UserId: "deposit", Balance: 0, Id: "123"}
	accountRepository := infrastructure.NewAccountRepository()
	accountRepository.Save(account)
	accountService := AccountService{}
	accountService.accountRepository = accountRepository

	suite.accountService = accountService
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) TestDeposit() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			suite.accountService.Deposit("deposit", 100)
		}()
	}
	wg.Wait()

	account := suite.accountService.Retrieve("deposit")

	assert.Equal(suite.T(), 500, account.Balance)
}
