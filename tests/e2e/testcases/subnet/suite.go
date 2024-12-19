// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subnet

import (
	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/shubhamdubey02/Cryft-cli/tests/e2e/commands"
	"github.com/shubhamdubey02/Cryft-cli/tests/e2e/utils"
)

const subnetName = "e2eSubnetTest"

var (
	mapping map[string]string
	err     error
)

var _ = ginkgo.Describe("[Subnet]", ginkgo.Ordered, func() {
	_ = ginkgo.BeforeAll(func() {
		mapper := utils.NewVersionMapper()
		mapping, err = utils.GetVersionMapping(mapper)
		gomega.Expect(err).Should(gomega.BeNil())
	})

	ginkgo.It("can create and delete a subnet evm config", func() {
		commands.CreateSubnetEvmConfig(subnetName, utils.SubnetEvmGenesisPath)
		commands.DeleteSubnetConfig(subnetName)
	})

	ginkgo.It("can create and delete a spacesvm config", func() {
		commands.CreateSpacesVMConfig(subnetName, utils.SpacesVMGenesisPath)
		commands.DeleteSubnetConfig(subnetName)
	})

	ginkgo.It("can create and delete a custom vm subnet config", func() {
		// let's use a SubnetEVM version which would be compatible with an existing Avago
		customVMPath, err := utils.DownloadCustomVMBin(mapping[utils.SoloSubnetEVMKey1])
		gomega.Expect(err).Should(gomega.BeNil())

		commands.CreateCustomVMConfig(subnetName, utils.SubnetEvmGenesisPath, customVMPath)
		commands.DeleteSubnetConfig(subnetName)
		exists, err := utils.SubnetCustomVMExists(subnetName)
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(exists).Should(gomega.BeFalse())
	})
})
