package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/Stride-Labs/stride/v4/x/distribution/client/cli"
	"github.com/Stride-Labs/stride/v4/x/distribution/client/rest"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
