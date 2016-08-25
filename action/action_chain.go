package action

type action func()

type Chain struct {
	actions []action
}

func (chain *Chain) Add(fn action) *Chain {
	chain.actions = append(chain.actions, fn)
	return chain
}

func (chain *Chain) Execute() {
	for _, fn := range chain.actions {
		fn()
	}
}
